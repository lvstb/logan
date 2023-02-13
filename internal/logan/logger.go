package logan

import (
	"sync"

	"go.uber.org/zap"
)

// cfg       zap.Config
var (
	zapLogger *zap.Logger
	once      sync.Once
)

func GetLogger(level string) *zap.Logger {
	// once.Do(func() {
	// 	// usr, err := user.Current()
	// 	// if err != nil {
	// 	// 	fmt.Printf("\n%s", err.Error())
	// 	// 	os.Exit(1)
	// 	// }
	// 	// var zapLevel zap.AtomicLevel
	// 	// if level == "debug" {
	// 	// 	zapLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	// 	// } else {
	// 	// 	zapLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	// 	// }
	// 	// cfg = zap.Config{
	// 	// 	Encoding:         "console",
	// 	// 	ErrorOutputPaths: []string{"stderr"},
	// 	// 	EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
	// 	// 	Level:            zapLevel,
	// 	// 	OutputPaths: []string{
	// 	// 		usr.HomeDir + string(os.PathSeparator) + ".logan" + string(os.PathSeparator) + "logan.log",
	// 	// 	},
	// 	// }
	// 	// zapLogger, err = cfg.Build()
	// 	// if err != nil {
	// 	// 	zapLogger, _ = zap.NewProduction()
	// 	// }
	// 	zapLogger, _ := zap.NewProduction()
	// 	defer func() {
	// 		_ = zapLogger.Sync()
	// 	}()
	// })
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	return zapLogger
}
