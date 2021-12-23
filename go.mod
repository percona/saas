module github.com/percona-platform/saas

go 1.17

replace github.com/percona-platform/platform => ../should-not-exist

require (
	github.com/golang/protobuf v1.4.3
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/percona/promconfig v0.2.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	go.starlark.net v0.0.0-20201210151846-e81fc95f7bd5
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.15.0 // indirect
	github.com/prometheus/procfs v0.2.0 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a1e // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
