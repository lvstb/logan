package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"github.com/glnds/logan/internal/logan"
	"go.uber.org/zap"
	// "github.com/aws/aws-lambda-go/lambda"
)

var (
	secretCache, _               = secretcache.New()
	logger                       *zap.Logger
	version, build, commit, date string
)

func main() {
	// 	conf := masl.GetConfig()
	// 	if conf.Debug {
	// 		logger = masl.GetLogger("debug")
	// 	} else {
	// 		logger = masl.GetLogger("info")
	// 	}
	logger = logan.GetLogger("info")
	// logger.Info("------------------ w00t w00t logan for you!?  ------------------")

	// flags := parseFlags(conf)
	logger.Info("Started...")

	DoLogan()
	return
}

// DoMasl Allow other tools to integrate with Masl to assume an AWS role
func DoLogan() {
	// lambda.Start(HandleRequest)
	logger.Info("Get Secret...")
	result, error := secretCache.GetSecretString("test-secret")
	fmt.Println(result)
	fmt.Println(error)

	logger.Info("Secret retrieved...")
	// Use the secret, return success
	os.Exit(0)
}
