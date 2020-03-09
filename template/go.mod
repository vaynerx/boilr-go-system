module {{ if Owner }}{{ Owner }}{{ end }}.{{ Project }}

go 1.14

require (
)

replace (
  {{ if Owner }}{{ Owner }}{{ end }}.{{ Project }}/cmd/app => ./cmd/app
  {{ if Owner }}{{ Owner }}{{ end }}.{{ Project }}/internal/app => ./internal/app
  {{ if Owner }}{{ Owner }}{{ end }}.{{ Project }}/api/proto/{{ Project }}/api => ./api/proto/{{ Project }}/api
  github.com/ugorji/go => github.com/ugorji/go/codec v1.1.7
  github.com/docker/docker => github.com/docker/docker v1.4.2-0.20190319215453-e7b5f7dbe98c
  github.com/containerd/containerd => github.com/containerd/containerd v1.2.1-0.20190507210959-7c1e88399ec0
	github.com/docker/docker => github.com/docker/docker v1.4.2-0.20190319215453-e7b5f7dbe98c
	golang.org/x/crypto v0.0.0-20190129210102-0709b304e793 => golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190507160741-ecd444e8653b
)