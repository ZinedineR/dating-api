package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"dating-api/pkg/migration"

	appConfiguration "dating-api/app/appconf"
	"dating-api/internal/base/handler"
	proHandler "dating-api/internal/profile/handler"
	profileRepo "dating-api/internal/profile/repository"
	profileService "dating-api/internal/profile/service"
	"dating-api/pkg/db"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	appConf            *appConfiguration.Config
	baseHandler        *handler.BaseHTTPHandler
	profileHandler     *proHandler.HTTPHandler
	postgresClientRepo *db.PostgreSQLClientRepository
)

func initPostgreSQL() {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbname := os.Getenv("DB_NAME")
	uname := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	var gConfig *gorm.Config
	if os.Getenv("DEV_SHOW_QUERY") == "True" {
		showQuery := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			})

		gConfig = &gorm.Config{Logger: showQuery}
	} else {
		gConfig = &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	}

	postgresClientRepo, _ = db.NewMPostgreSQLRepository(host, uname, pass, dbname, port, gConfig)
	migration.Initmigrate(postgresClientRepo.DB)

}

func initHTTP() {
	appConf = appConfiguration.InitAppConfig()
	initInfrastructure()

	appConf.MysqlTZ = postgresClientRepo.TZ

	baseHandler = handler.NewBaseHTTPHandler(postgresClientRepo.DB, appConf, postgresClientRepo)
	profileRepo := profileRepo.NewRepository(postgresClientRepo.DB, postgresClientRepo)
	profileService := profileService.NewService(profileRepo)
	profileHandler = proHandler.NewHTTPHandler(baseHandler, profileService)
}

func initInfrastructure() {
	initPostgreSQL()
	initLog()

}

func isProd() bool {
	return os.Getenv("APP_ENV") == "production"
}

func initLog() {
	lv := os.Getenv("LOG_LEVEL_DEV")
	level := logrus.InfoLevel
	switch lv {
	case "PanicLevel":
		level = logrus.PanicLevel
	case "FatalLevel":
		level = logrus.FatalLevel
	case "ErrorLevel":
		level = logrus.ErrorLevel
	case "WarnLevel":
		level = logrus.WarnLevel
	case "InfoLevel":
		level = logrus.InfoLevel
	case "DebugLevel":
		level = logrus.DebugLevel
	case "TraceLevel":
		level = logrus.TraceLevel
	default:
	}

	if isProd() {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.WarnLevel)
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

		if lv == "" && os.Getenv("APP_DEBUG") == "True" {
			level = logrus.DebugLevel
		}
		logrus.SetLevel(level)

		if os.Getenv("DEV_FILE_LOG") == "True" {
			logfile, err := os.OpenFile("log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				fmt.Printf("error opening file : %v", err)
			}

			mw := io.MultiWriter(os.Stdout, logfile)
			logrus.SetOutput(mw)
		} else {
			logrus.SetOutput(os.Stdout)
		}
	}
}
