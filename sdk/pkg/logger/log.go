package logger

import (
	"io"
	"os"

	"github.com/go-admin-team/go-admin-core/debug/writer"
	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/plugins/logger/zap"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	log "github.com/go-admin-team/go-admin-core/logger"
)

// SetupLogger 日志
func SetupLogger(logType, path, levelStr, outputType string) logger.Logger {
	if !pkg.PathExist(path) {
		err := pkg.PathCreate(path)
		if err != nil {
			log.Fatalf("create dir error: %s", err.Error())
		}
	}
	var err error
	var output io.Writer
	switch outputType {
	case "file":
		output, err = writer.NewFileWriter(
			writer.WithPathOption(path),
		)
		if err != nil {
			log.Fatal("logger setup error: %s", err.Error())
		}
	default:
		output = os.Stdout
	}
	var level logger.Level
	level, err = logger.GetLevel(levelStr)
	if err != nil {
		log.Fatalf("get logger level error, %s", err.Error())
	}

	switch logType {
	case "zap":
		log.DefaultLogger, err = zap.NewLogger(logger.WithLevel(level), logger.WithOutput(output), zap.WithCallerSkip(2))
		if err != nil {
			log.Fatalf("new zap logger error, %s", err.Error())
		}
	//case "logrus":
	//	setLogger = logrus.NewLogger(logger.WithLevel(level), logger.WithOutput(output), logrus.ReportCaller())
	default:
		log.DefaultLogger = logger.NewLogger(logger.WithLevel(level), logger.WithOutput(output))
	}
	return log.DefaultLogger
}
