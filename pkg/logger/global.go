package logger

import (
	"go.uber.org/zap"
)

// SetupGlobal setups global zap logger.
func SetupGlobal() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(l)
}
