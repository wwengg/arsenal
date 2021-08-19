// @Title  
// @Description  
// @Author  Wangwengang  2021/8/18 下午10:01
// @Update  Wangwengang  2021/8/18 下午10:01
package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wwengg/arsenal/config"
	"github.com/wwengg/arsenal/utils"
)


var (
	level zapcore.Level
	ZapLog *zap.Logger
)

func Reset(){
	ZapLog = Zap()
}


func Zap()(logger *zap.Logger){
	if ok, _ := utils.PathExists(config.ConfigHub.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.ConfigHub.Zap.Director)
		_ = os.Mkdir(config.ConfigHub.Zap.Director, os.ModePerm)
	}

	switch config.ConfigHub.Zap.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if config.ConfigHub.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}


// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (conf zapcore.EncoderConfig) {
	conf = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  config.ConfigHub.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case config.ConfigHub.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		conf.EncodeLevel = zapcore.LowercaseLevelEncoder
	case config.ConfigHub.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		conf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case config.ConfigHub.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		conf.EncodeLevel = zapcore.CapitalLevelEncoder
	case config.ConfigHub.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		conf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		conf.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return conf
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if config.ConfigHub.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := utils.GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(config.ConfigHub.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}



