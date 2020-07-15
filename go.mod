module github.com/percona-platform/saas

go 1.14

replace github.com/percona-platform/platform => ../should-not-exist

require (
	github.com/golang/protobuf v1.3.5
	github.com/mwitkow/go-proto-validators v0.3.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	go.starlark.net v0.0.0-20200619143648-50ca820fafb9
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.30.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)
