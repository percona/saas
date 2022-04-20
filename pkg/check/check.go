// Package check implements checks parsing and validation.
package check

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/blake2b"
	"gopkg.in/yaml.v3"

	"github.com/percona-platform/saas/pkg/common"
)

// Verify checks signature of passed data with provided public key and
// returns error in case of any problem.
func Verify(data []byte, publicKey, sig string) error { //nolint: cyclop
	lines := strings.SplitN(sig, "\n", 4)
	if len(lines) < 4 {
		return errors.New("incomplete signature")
	}

	sBin, err := base64.StdEncoding.DecodeString(lines[1])
	if err != nil || len(sBin) != 74 {
		return errors.New("invalid signature")
	}
	gBin, err := base64.StdEncoding.DecodeString(lines[3])
	if err != nil || len(gBin) != 64 {
		return errors.New("invalid global signature")
	}
	kBin, err := base64.StdEncoding.DecodeString(publicKey) //nolint:revive
	if err != nil || len(kBin) != 42 {
		return errors.New("invalid public key")
	}

	sAlg, sKeyID, sSig := sBin[0:2], sBin[2:10], sBin[10:74]
	kAlg, kKeyID, kKey := kBin[0:2], kBin[2:10], kBin[10:42] //nolint:revive

	// Key algorithm should be `Ed`.
	if kAlg[0] != 0x45 || kAlg[1] != 0x64 {
		return errors.New("unsupported key algorithm")
	}
	// Signature algorithm can be `Ed`(legacy) or `ED`(pre-hashed).
	if sAlg[0] != 0x45 || (sAlg[1] != 0x64 && sAlg[1] != 0x44) {
		return errors.New("unsupported signature algorithm")
	}

	// For pre-hashed signature get data hash.
	if sAlg[1] == 0x44 {
		h, _ := blake2b.New512(nil)
		h.Write(data)
		data = h.Sum(nil)
	}

	if !bytes.Equal(kKeyID, sKeyID) {
		return errors.New("incompatible key identifiers")
	}
	if !strings.HasPrefix(lines[2], "trusted comment: ") {
		return errors.New("unexpected format for the trusted comment")
	}
	if !ed25519.Verify(ed25519.PublicKey(kKey), data, sSig) {
		return errors.New("invalid signature")
	}
	if !ed25519.Verify(ed25519.PublicKey(kKey), append(sSig, []byte(lines[2])[17:]...), gBin) {
		return errors.New("invalid global signature")
	}
	return nil
}

// ParseParams represents optional Parse function parameters.
type ParseParams struct {
	DisallowUnknownFields bool // if true, return errors for unexpected YAML fields
	DisallowInvalidChecks bool // if true, return errors for invalid checks instead of skipping them
}

// Parse returns a slice of validated checks parsed from YAML passed via a reader.
// It can handle multi-document YAMLs: parsing result will be a single slice
// that contains checks form every parsed document.
func Parse(reader io.Reader, params *ParseParams) ([]Check, error) {
	if params == nil {
		params = new(ParseParams)
	}

	d := yaml.NewDecoder(reader)
	d.KnownFields(params.DisallowUnknownFields)

	type checks struct {
		Checks []Check `yaml:"checks"`
	}

	var res []Check
	for {
		var c checks
		if err := d.Decode(&c); err != nil {
			if errors.Is(err, io.EOF) {
				return res, nil
			}
			return nil, errors.Wrap(err, "failed to parse checks")
		}

		for _, check := range c.Checks {
			if err := check.Validate(); err != nil {
				if params.DisallowInvalidChecks {
					return nil, err
				}

				continue // skip invalid check
			}

			res = append(res, check)
		}
	}
}

// Supported check types.
const (
	MySQLShow                = Type("MYSQL_SHOW")
	MySQLSelect              = Type("MYSQL_SELECT")
	PostgreSQLShow           = Type("POSTGRESQL_SHOW")
	PostgreSQLSelect         = Type("POSTGRESQL_SELECT")
	MongoDBGetParameter      = Type("MONGODB_GETPARAMETER")
	MongoDBBuildInfo         = Type("MONGODB_BUILDINFO")
	MongoDBGetCmdLineOpts    = Type("MONGODB_GETCMDLINEOPTS")
	MongoDBReplSetGetStatus  = Type("MONGODB_REPLSETGETSTATUS")
	MongoDBGetDiagnosticData = Type("MONGODB_GETDIAGNOSTICDATA")
)

// Type represents check type.
type Type string

// Validate validates check type.
func (t Type) Validate() error { // nolint:cyclop
	switch t {
	case MySQLShow:
		fallthrough
	case MySQLSelect:
		fallthrough
	case PostgreSQLShow:
		fallthrough
	case PostgreSQLSelect:
		fallthrough
	case MongoDBGetParameter:
		fallthrough
	case MongoDBBuildInfo:
		fallthrough
	case MongoDBGetCmdLineOpts:
		fallthrough
	case MongoDBReplSetGetStatus:
		fallthrough
	case MongoDBGetDiagnosticData:
		return nil
	case "":
		return errors.New("check type is empty")
	default:
		return errors.Errorf("unknown check type: %s", t)
	}
}

// Supported DB families.
const (
	MySQL      = Family("MYSQL")
	PostgreSQL = Family("POSTGRESQL")
	MongoDB    = Family("MONGODB")
)

// Family represents monitored service family.
type Family string

// Validate validates check family.
func (f Family) Validate() error {
	switch f {
	case MySQL:
		fallthrough
	case PostgreSQL:
		fallthrough
	case MongoDB:
		return nil
	case "":
		return errors.New("check family is empty")
	default:
		return errors.Errorf("unknown check family: %s", f)
	}
}

// Supported check intervals.
const (
	Standard = Interval("standard")
	Frequent = Interval("frequent")
	Rare     = Interval("rare")
)

// Interval represents check execution interval.
type Interval string

// Validate validates check interval.
func (i Interval) Validate() error {
	switch i {
	case Standard:
		fallthrough
	case Frequent:
		fallthrough
	case Rare:
		fallthrough
	case "":
		return nil
	default:
		return errors.Errorf("unknown check interval: %s", i)
	}
}

// Query represents DB query of specified type.
type Query struct {
	Query string
	Type  Type
}

// Check represents security check structure. Fields marked with v1 should not be used for version 2, and vice versa.
type Check struct {
	Version     uint32        `yaml:"version"`
	Name        string        `yaml:"name"`
	Summary     string        `yaml:"summary"`
	Description string        `yaml:"description"`
	Type        Type          `yaml:"type,omitempty"`     // for v1
	Category    string        `yaml:"category,omitempty"` // optional for v1, required for v2
	Family      Family        `yaml:"family,omitempty"`   // for v2
	Tiers       []common.Tier `yaml:"tiers,flow,omitempty"`
	Interval    Interval      `yaml:"interval,omitempty"`
	Query       string        `yaml:"query,omitempty"`   // for v1
	Queries     []Query       `yaml:"queries,omitempty"` // for v2
	Script      string        `yaml:"script"`
}

// The same as Prometheus label format.
var nameRE = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// Validate validates check for minimal correctness.
func (c *Check) Validate() error { //nolint: cyclop
	var err error

	if !nameRE.MatchString(c.Name) {
		return errors.New("invalid check name")
	}

	if err = common.ValidateTiers(c.Tiers); err != nil {
		return err
	}

	if err = c.Interval.Validate(); err != nil {
		return err
	}

	if err = c.validateScript(); err != nil {
		return err
	}

	if c.Script == "" {
		return errors.New("check script is empty")
	}

	if c.Summary == "" {
		return errors.New("summary is empty")
	}

	if c.Description == "" {
		return errors.New("description is empty")
	}

	switch c.Version {
	case 1:
		return c.validateV1()
	case 2:
		return c.validateV2()
	default:
		return errors.Errorf("unexpected version %d", c.Version)
	}
}

func (c *Check) validateV1() error {
	var err error
	if err = c.Type.Validate(); err != nil {
		return err
	}

	if err = validateQuery(c.Type, c.Query); err != nil {
		return err
	}

	if c.Family != "" {
		return errors.New("field 'family' is part of check format version 2 and can't be used in version 1")
	}

	if len(c.Queries) != 0 {
		return errors.New("field 'queries' is part of check format version 2 and can't be used in version 1")
	}

	return nil
}

func (c *Check) validateV2() error {
	var err error
	if err = c.Family.Validate(); err != nil {
		return err
	}

	if err = c.validateQueries(); err != nil {
		return err
	}

	if c.Type != "" {
		return errors.New("field 'type' is part of check format version 1 and can't be used in version 2")
	}

	if c.Query != "" {
		return errors.New("field 'query' is part of check format version 1 and can't be used in version 2")
	}

	// category is optional for v1 so that existing checks can continue working, but we make it required for v2.
	if c.Category == "" {
		return errors.New("category is empty")
	}

	return nil
}

func (c *Check) validateScript() error {
	if strings.ContainsRune(c.Script, '\t') {
		return errors.New("script should use spaces for indentation, not tabs")
	}

	return nil
}

func validateQuery(typ Type, query string) error { // nolint:cyclop
	switch typ {
	case PostgreSQLShow:
		fallthrough
	case MongoDBGetParameter:
		fallthrough
	case MongoDBBuildInfo:
		fallthrough
	case MongoDBGetCmdLineOpts:
		fallthrough
	case MongoDBReplSetGetStatus:
		fallthrough
	case MongoDBGetDiagnosticData:
		if query != "" {
			return errors.Errorf("query should be empty for %s type", typ)
		}
	case PostgreSQLSelect:
		fallthrough
	case MySQLShow:
		fallthrough
	case MySQLSelect:
		if query == "" {
			return errors.New("query is empty")
		}
	}

	return nil
}

func (c *Check) validateQueries() error {
	if len(c.Queries) == 0 {
		return errors.New("check should have at least one query")
	}

	switch c.Family {
	case MySQL:
		return validateMySQLCompatibleQueries(c.Queries)
	case PostgreSQL:
		return validatePostgreSQLCompatibleQueries(c.Queries)
	case MongoDB:
		return validateMongoDBCompatibleQueries(c.Queries)
	default:
		return errors.Errorf("unknown check family: %s", c.Family)
	}
}

func validateMySQLCompatibleQueries(queries []Query) error {
	var err error
	for _, q := range queries {
		if err = q.Type.Validate(); err != nil {
			return err
		}

		switch q.Type { //nolint:exhaustive
		case MySQLShow:
		case MySQLSelect:
		default:
			return errors.Errorf("unsupported query type '%s' for mySQL family", q.Type)
		}

		if err = validateQuery(q.Type, q.Query); err != nil {
			return err
		}
	}

	return nil
}

func validatePostgreSQLCompatibleQueries(queries []Query) error {
	var err error
	for _, q := range queries {
		if err = q.Type.Validate(); err != nil {
			return err
		}

		switch q.Type { //nolint:exhaustive
		case PostgreSQLShow:
		case PostgreSQLSelect:
		default:
			return errors.Errorf("unsupported query type '%s' for postgreSQL family", q.Type)
		}

		if err = validateQuery(q.Type, q.Query); err != nil {
			return err
		}
	}

	return nil
}

func validateMongoDBCompatibleQueries(queries []Query) error {
	var err error
	for _, q := range queries {
		if err = q.Type.Validate(); err != nil {
			return err
		}

		switch q.Type { //nolint:exhaustive
		case MongoDBGetParameter:
		case MongoDBBuildInfo:
		case MongoDBGetCmdLineOpts:
		default:
			return errors.Errorf("unsupported query type '%s' for mongoDB family", q.Type)
		}

		if err = validateQuery(q.Type, q.Query); err != nil {
			return err
		}
	}

	return nil
}
