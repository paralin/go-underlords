module github.com/paralin/go-underlords

go 1.13

// note: protobuf is intentionally held at 1.3.x
replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.5

require (
	github.com/faceit/go-steam v0.0.0-20190206021251-2be7df6980e1
	github.com/fatih/camelcase v1.0.0
	github.com/golang/protobuf v1.4.2
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
	github.com/sirupsen/logrus v1.6.0
	github.com/urfave/cli v1.22.4
)
