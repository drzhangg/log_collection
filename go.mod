module log_collection

go 1.14

require (
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/nsqio/go-nsq v1.0.8
	github.com/streadway/amqp v1.0.0
	go.uber.org/zap v1.15.0 // indirect
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485 // indirect
	google.golang.org/grpc v1.31.0 // indirect
)

//解决etcd依赖冲突
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

//解决etcd依赖冲突
replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
