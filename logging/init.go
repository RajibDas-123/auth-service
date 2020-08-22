package logging

import (
	"log"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	AppLogger     *zap.SugaredLogger
	RequestLogger *zap.SugaredLogger
	DBLogger      *zap.SugaredLogger
	CacheLogger   *zap.SugaredLogger
	RestLogger    *zap.SugaredLogger
)

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
	})
}
func Initialize() {
	var build = os.Getenv("BUILD")
	var path string
	if build == "Dev" {
		path, _ = os.Getwd()
		path = filepath.Join(path, "log")
	} else if build == "Prod" {
		path = filepath.Join("var", "log", "dp_cron")
	} else {
		log.Fatal("Unexpected BUILD value in .env file")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Fatal("Failed to create the folder for logs", err)
		}
	}

	appLoggerCore := zapcore.NewCore(getEncoder(), getLogWriter(filepath.Join(path, "app.log")), zapcore.InfoLevel)
	requestLoggerCore := zapcore.NewCore(getEncoder(), getLogWriter(filepath.Join(path, "request.log")), zapcore.InfoLevel)
	dbLoggerCore := zapcore.NewCore(getEncoder(), getLogWriter(filepath.Join(path, "db.log")), zapcore.InfoLevel)
	cacheLoggerCore := zapcore.NewCore(getEncoder(), getLogWriter(filepath.Join(path, "cache.log")), zapcore.InfoLevel)
	restAPIRequestLoggerCore := zapcore.NewCore(getEncoder(), getLogWriter(filepath.Join(path, "rest_req.log")), zapcore.InfoLevel)

	AppLogger = zap.New(appLoggerCore, zap.AddCaller()).Sugar()
	RequestLogger = zap.New(requestLoggerCore, zap.AddCaller()).Sugar()
	DBLogger = zap.New(dbLoggerCore, zap.AddCaller()).Sugar()
	CacheLogger = zap.New(cacheLoggerCore, zap.AddCaller()).Sugar()
	RestLogger = zap.New(restAPIRequestLoggerCore, zap.AddCaller()).Sugar()

}
