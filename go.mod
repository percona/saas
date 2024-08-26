module github.com/percona/saas

go 1.21

replace github.com/percona/platform => ../should-not-exist

require (
	github.com/golang/protobuf v1.5.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/percona/promconfig v0.2.5
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.20.2
	go.starlark.net v0.0.0-20201210151846-e81fc95f7bd5
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.25.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.34.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
)

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240227224415-6ceb2ff114de
)
