module github.com/percona-platform/saas

go 1.14

replace github.com/percona-platform/platform => ../should-not-exist

require (
	github.com/golang/protobuf v1.4.2
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	go.starlark.net v0.0.0-20200901195727-6e684ef5eeee
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)
