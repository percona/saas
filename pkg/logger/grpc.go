package logger

import (
	"go.uber.org/zap"
	"google.golang.org/grpc/grpclog"
)

// GRPC is a compatibility wrapper between zap's sugared logger entry and gRPC logger interface.
type GRPC struct {
	*zap.SugaredLogger

	// Set to true for very verbose gRPC logging.
	Verbose bool
}

// V reports whether verbosity level l is at least the requested verbose level.
func (g *GRPC) V(l int) bool {
	return g.Verbose
}

// Infoln prints log message with info level.
func (g *GRPC) Infoln(args ...interface{}) { g.Info(args...) }

// Warning prints log message with warning level.
func (g *GRPC) Warning(args ...interface{}) { g.Warn(args...) }

// Warningln similar to Warning.
func (g *GRPC) Warningln(args ...interface{}) { g.Warn(args...) }

// Warningf prints warning level log message wit given format.
func (g *GRPC) Warningf(format string, args ...interface{}) { g.Warnf(format, args...) }

// Errorln prints log message with error level.
func (g *GRPC) Errorln(args ...interface{}) { g.Error(args...) }

// Fatalln prints log message and exit program.
func (g *GRPC) Fatalln(args ...interface{}) { g.Fatal(args...) }

// Check interfaces.
var _ grpclog.LoggerV2 = (*GRPC)(nil)
