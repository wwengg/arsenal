package internal

import (
	"fmt"

	"gorm.io/gorm/logger"

	"github.com/wwengg/arsenal/config/conf"
	logger2 "github.com/wwengg/arsenal/logger"
)

type writer struct {
	logger.Writer
	conf.Mysql
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer,m conf.Mysql) *writer {
	return &writer{Writer: w,Mysql: m}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	logZap = w.LogZap
	if logZap {
		logger2.ZapLog.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
