package check

import (
	"io"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/percona-platform/saas/pkg/common"
)

// Advisor represents advisor structure.
type Advisor struct {
	Version     uint32        `yaml:"version"`
	Name        string        `yaml:"name"`
	Summary     string        `yaml:"summary"`
	Description string        `yaml:"description"`
	Category    string        `yaml:"category"`
	Tiers       []common.Tier `yaml:"tiers,flow"`
	Checks      []Check       `yaml:"checks"`
}

// ParseAdvisors returns a slice of validated advisors parsed from YAML passed via a reader.
// It can handle multi-document YAMLs: parsing result will be a single slice
// that contains advisors form every parsed document.
func ParseAdvisors(reader io.Reader, params *ParseParams) ([]Advisor, error) {
	if params == nil {
		params = new(ParseParams)
	}

	d := yaml.NewDecoder(reader)
	d.KnownFields(params.DisallowUnknownFields)

	type advisors struct {
		Advisors []Advisor `yaml:"advisors"`
	}

	var res []Advisor
	for {
		var c advisors
		if err := d.Decode(&c); err != nil {
			if errors.Is(err, io.EOF) {
				return res, nil
			}
			return nil, errors.Wrap(err, "failed to parse advisors")
		}

		for _, advisor := range c.Advisors {
			if err := advisor.Validate(); err != nil {
				if params.DisallowInvalidChecks {
					return nil, err
				}

				continue // skip invalid advisors
			}

			res = append(res, advisor)
		}
	}
}

// Validate advisor.
func (a *Advisor) Validate() error {
	if a.Version != 1 {
		return errors.Errorf("unexpected version %d", a.Version)
	}

	if !nameRE.MatchString(a.Name) {
		return errors.New("invalid advisor name")
	}

	if a.Summary == "" {
		return errors.New("summary is empty")
	}

	if a.Description == "" {
		return errors.New("description is empty")
	}

	if a.Category == "" {
		return errors.New("category is empty")
	}

	if err := common.ValidateTiers(a.Tiers); err != nil {
		return err
	}

	for _, check := range a.Checks {
		if err := check.Validate(); err != nil {
			return err
		}
	}

	return nil
}
