package main

import (
	"database/sql"
	"flag"
	"github.com/TechnoHandOver/backend/config"
	AdsDelivery "github.com/TechnoHandOver/backend/internal/ad/delivery"
	AdsRepository "github.com/TechnoHandOver/backend/internal/ad/repository"
	AdsUsecase "github.com/TechnoHandOver/backend/internal/ad/usecase"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	SessionDelivery "github.com/TechnoHandOver/backend/internal/session/delivery"
	SessionRepository "github.com/TechnoHandOver/backend/internal/session/repository"
	SessionUsecase "github.com/TechnoHandOver/backend/internal/session/usecase"
	"github.com/TechnoHandOver/backend/internal/tools/properties"
	HandoverValidator "github.com/TechnoHandOver/backend/internal/tools/validator"
	UserDelivery "github.com/TechnoHandOver/backend/internal/user/delivery"
	UserRepository "github.com/TechnoHandOver/backend/internal/user/repository"
	UserUsecase "github.com/TechnoHandOver/backend/internal/user/usecase"
	"github.com/labstack/echo/v4"
	LabstackLog "github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const driverName = "postgres"

func main() {
	var configFileName, logFileName string
	flag.StringVar(&configFileName, "configFileName", "config.json", "path to server config file")
	flag.StringVar(&logFileName, "logFileName", "logs.log", "path to server log file")
	flag.Parse()

	config_, err := config.LoadConfigFile(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	properties.Properties = config_.Properties

	var logFile *os.File
	if logFile, err = os.OpenFile(logFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = logFile.Close()
	}()

	log.SetOutput(logFile)

	db, err := sql.Open(driverName, config_.GetDatabaseConfigString())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println(config_.GetDatabaseConfigString())

	adsRepository := AdsRepository.NewAdRepositoryImpl(db)
	sessionRepository := SessionRepository.NewSessionRepositoryImpl()
	userRepository := UserRepository.NewUserRepositoryImpl(db)

	adsUsecase := AdsUsecase.NewAdUsecaseImpl(adsRepository)
	userUsecase := UserUsecase.NewUserUsecaseImpl(userRepository)
	sessionUsecase := SessionUsecase.NewSessionUsecaseImpl(sessionRepository)

	adsDelivery := AdsDelivery.NewAdDelivery(adsUsecase)
	sessionDelivery := SessionDelivery.NewSessionDelivery(sessionUsecase, userUsecase)
	userDelivery := UserDelivery.NewUserDelivery(userUsecase)

	recoverMiddleware := middlewares.NewRecoverMiddleware()
	authMiddleware := middlewares.NewAuthMiddleware(sessionUsecase, userUsecase)
	middlewaresManager := middlewares.NewManager(recoverMiddleware, authMiddleware)

	echo_ := echo.New()
	echo_.Logger.SetLevel(LabstackLog.ERROR)
	if logFile != nil {
		echo_.Logger.SetOutput(logFile)
	}
	echo_.Use(middlewaresManager.RecoverMiddleware.Recover())
	echo_.Validator = HandoverValidator.NewRequestValidator()

	adsDelivery.Configure(echo_, middlewaresManager)
	sessionDelivery.Configure(echo_, middlewaresManager)
	userDelivery.Configure(echo_, middlewaresManager)

	if err := echo_.Start(config_.GetServerConfigString()); err != nil {
		log.Fatal(err)
	}
}
