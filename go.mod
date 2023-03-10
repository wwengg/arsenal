module github.com/wwengg/arsenal

go 1.15

require (
	github.com/chanxuehong/wechat v0.0.0-20230222024006-36f0325263cd // indirect
	github.com/fsnotify/fsnotify v1.5.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/gorilla/websocket v1.4.2
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/rpcxio/rpcx-etcd v0.0.0-20210606082724-1a5593db7a0d
	github.com/smallnest/rpcx v1.7.4
	github.com/spf13/viper v1.7.0
	github.com/wwengg/proto v0.0.3
	go.uber.org/zap v1.19.0
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/grpc v1.36.0 // indirect
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.22.5
)

//replace (
//github.com/wwengg/proto => ../proto
//)
