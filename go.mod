module github.com/percona-platform/saas

go 1.17

replace github.com/percona-platform/platform => ../should-not-exist

require (
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/percona/promconfig v0.2.5
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.16.0
	go.starlark.net v0.0.0-20201210151846-e81fc95f7bd5
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.9.0
	google.golang.org/genproto v0.0.0-20230526203410-71b5a4ffd15e
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)
