module github.com/percona-platform/saas

go 1.14

replace github.com/percona-platform/platform => ../should-not-exist

require (
	github.com/golang/protobuf v1.5.1
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/percona/promconfig v0.2.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	go.starlark.net v0.0.0-20201210151846-e81fc95f7bd5
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)
