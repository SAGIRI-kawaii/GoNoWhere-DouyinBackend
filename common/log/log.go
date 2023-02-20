package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger

func init() {
	writeSync := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSync, zapcore.DebugLevel)

	logger := zap.New(core)
	SugarLogger = logger.Sugar()
}

// 设置日志写入或者输出路径
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./log/log.json")
	return zapcore.AddSync(file)
}

// 设置日志输出以及编码格式 json 或其他
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
