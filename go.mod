module github.com/percona-platform/saas

go 1.14

replace github.com/percona-platform/platform => ../should-not-exist

require (
	github.com/golang/protobuf v1.3.5
	github.com/mwitkow/go-proto-validators v0.3.0
	github.com/pkg/errors v0.9.1
	go.starlark.net v0.0.0-20200330013621-be5394c419b6
	google.golang.org/grpc v1.29.1
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)
