package bootstrap

import (
	"github.com/Wpc-0601/copilot-genuine/common"
	"github.com/Wpc-0601/copilot-genuine/config/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var level zapcore.Level
var options []zap.Option

func InitLog() *zap.Logger {
	if ok, _ := common.PathExist(global.App.Configuration.Log.RootDir); !ok {
		_ = os.Mkdir(global.App.Configuration.Log.RootDir, os.ModePerm)
	}
	setLogLevel()
	if global.App.Configuration.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}
	return zap.New(ExtendZap(), options...)
}

func ExtendZap() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.Configuration.App.Env + "." + l.String())
	}

	// 设置编码器
	if global.App.Configuration.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.App.Configuration.Log.RootDir + "/" + global.App.Configuration.Log.FileName,
		MaxSize:    global.App.Configuration.Log.MaxSize,
		MaxBackups: global.App.Configuration.Log.MaxBackups,
		MaxAge:     global.App.Configuration.Log.MaxAge,
		Compress:   global.App.Configuration.Log.Compress,
	}

	return zapcore.AddSync(file)
}

func setLogLevel() {
	switch global.App.Configuration.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}
