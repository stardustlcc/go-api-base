package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	//默认的时间格式
	DefaultTimeLayout = time.RFC3339

	//默认的日志地址
	DefaultLogPath = "logs/log"
)

type Option func(*option)

type option struct {
	logPath string

	timeLayout string

	fields map[string]string
}

func NewLogger(opts ...Option) (*zap.Logger, error) {

	opt := &option{
		fields:     make(map[string]string),
		logPath:    DefaultLogPath,
		timeLayout: DefaultTimeLayout,
	}
	for _, f := range opts {
		f(opt)
	}

	//日志目录
	now := time.Now()
	infoLogFileName := fmt.Sprintf("%s/info/%04d-%02d-%02d.log", opt.logPath, now.Year(), now.Month(), now.Day())
	errorLogFileName := fmt.Sprintf("%s/error/%04d-%02d-%02d.log", opt.logPath, now.Year(), now.Month(), now.Day())

	//配置zapcore的编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "file",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,     //输出的分隔符
		EncodeLevel:   zapcore.LowercaseLevelEncoder, //序列化字符串的大小写
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format(opt.timeLayout))
		}, //时间的编码格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //时间显示的位数
		EncodeCaller:   zapcore.ShortCallerEncoder,     //输出的运行文件路径长度
	}

	//日志的编码格式：json or Cansole,这里采用了json的编码格式
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// //日志级别-高级别
	highPriority := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.ErrorLevel
	})

	// //日志级别-底级别
	lowPriority := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < zapcore.ErrorLevel
	})

	//创建文件写入器 info debug
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   infoLogFileName, //日志文件路径
		MaxSize:    128,             //每个日志文件保存的大小，单位:M
		MaxAge:     30,              //文件最多保存多少天
		MaxBackups: 100,             //日志文件最多保存多个备份
		Compress:   false,           //是否压缩
	})

	//创建文件写入器 error
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   errorLogFileName, //日志文件路径
		MaxSize:    128,              //每个日志文件保存的大小，单位:M
		MaxAge:     30,               //文件最多保存多少天
		MaxBackups: 100,              //日志文件最多保存多个备份
		Compress:   false,            //是否压缩
	})

	//创建核心
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoFileWriteSyncer, lowPriority),
		zapcore.NewCore(encoder, errorFileWriteSyncer, highPriority),
	)

	stderr := zapcore.Lock(os.Stderr) //lock for concurrent safe
	logger := zap.New(
		core,
		zap.AddCaller(), //记录文件名和行号
		zap.ErrorOutput(stderr),
	)
	for key, value := range opt.fields {
		logger = logger.WithOptions(zap.Fields(zapcore.Field{Key: key, Type: zapcore.StringType, String: value}))
	}
	return logger, nil

}

// 设置时间格式
func WithTimeLayout(timeLayout string) Option {
	return func(opt *option) {
		opt.timeLayout = timeLayout
	}
}

// 日志存储的目录地址
func WithLogPath(logPath string) Option {
	return func(opt *option) {
		opt.logPath = logPath
	}
}

func WithField(key, value string) Option {
	return func(opt *option) {
		opt.fields[key] = value
	}
}
