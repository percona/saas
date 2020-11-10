package logger

import "go.uber.org/zap"

// FlagsParsed is used to catch the common service initialization problem.
// Do not set the value directly in the service code.
//
// TODO Can we do that better, without an exported global variable, dependency cycle,
// and too much complexity of mutex-protected getters/setters?
// https://jira.percona.com/browse/SAAS-275
var FlagsParsed bool //nolint:gochecknoglobals

// SetupGlobalOpts contains logger options.
type SetupGlobalOpts struct {
	LogDebug   bool // enable debug level logging
	LogDevMode bool // enable development mode logging: text instead of JSON, DPanic panics instead of logging errors
}

// SetupGlobal setups global zap logger.
func SetupGlobal(opts *SetupGlobalOpts) {
	// catch the common service initialization problem
	if !FlagsParsed {
		panic("logger.SetupGlobal should be called after app.Setup and kingpin.Parse")
	}

	if opts == nil {
		opts = new(SetupGlobalOpts)
	}

	cfg := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	if opts.LogDebug {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	if opts.LogDevMode {
		cfg.Development = true
		cfg.Encoding = "console"
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(l)
}
