package log

import (
	"bossfi-indexer/src/common"
	"bossfi-indexer/src/core/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var Logger *zap.Logger

func InitLog() *zap.Logger {
	// 通用编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 彩色高亮级别（控制台生效）
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   shortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 日志级别控制
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel) // 可调为 Info/Warn/Error

	var core zapcore.Core

	if common.IsDev() {
		// 开发环境：仅控制台，彩色文本
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleOutput := zapcore.AddSync(os.Stdout)

		core = zapcore.NewCore(consoleEncoder, consoleOutput, atomicLevel)

	} else {
		// 非开发环境：控制台 JSON + 文件 JSON
		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "./logs/" + config.Conf.App.Name + ".log",
			MaxSize:    50, // MB
			MaxBackups: 20,
			MaxAge:     7, // days
			Compress:   true,
		})

		jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

		core = zapcore.NewTee(
			zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), atomicLevel),
			zapcore.NewCore(jsonEncoder, fileWriter, atomicLevel),
		)
	}

	// 附加配置项
	caller := zap.AddCaller()
	development := zap.Development()
	field := zap.Fields(zap.String("service", config.Conf.App.Name))

	Logger = zap.New(core, caller, development, field)
	return Logger
}

// 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 精简line路径
func shortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	root := common.GetCurrentAbPath() // 获取项目根目录
	path := caller.File
	if rel, err := filepath.Rel(root, path); err == nil {
		enc.AppendString(rel + ":" + strconv.Itoa(caller.Line))
	} else {
		enc.AppendString(path + ":" + strconv.Itoa(caller.Line))
	}
}
