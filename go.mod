module github.com/wwengg/arsenal

go 1.15

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/rpcxio/rpcx-etcd v0.0.0-20210606082724-1a5593db7a0d
	github.com/smallnest/rpcx v1.6.9
	github.com/spf13/viper v1.7.0
	go.uber.org/zap v1.19.0
	google.golang.org/grpc/examples v0.0.0-20210818220435-8ab16ef276a3 // indirect
	github.com/wwengg/proto v0.0.0
)


replace (
	github.com/wwengg/proto => ../proto
)