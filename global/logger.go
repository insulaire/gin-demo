package global

import (
	"gin-demo/pkg/logger"
	"log"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
)

var (
	Logger      *logger.Logger
	LogSavePath string
	LogFileName string
	LogFileExt  string
)

func init() {
	var err error
	LogSavePath, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	LogFileName = "log"
	LogFileExt = ".txt"
	if err = setupLogger(); err != nil {
		log.Panic(err)
	}
}

func setupLogger() error {
	fileName := LogSavePath + "/" + LogFileName + LogFileExt
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
