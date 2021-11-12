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
	RequestValidator "github.com/TechnoHandOver/backend/internal/tools/validator"
	UserDelivery "github.com/TechnoHandOver/backend/internal/user/delivery"
	UserRepository "github.com/TechnoHandOver/backend/internal/user/repository"
	UserUsecase "github.com/TechnoHandOver/backend/internal/user/usecase"
	Validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const driverName = "postgres"

func main() {
	flag.Parse()

	var configFileName string
	flag.StringVar(&configFileName, "configFileName", "config.json", "path to server config file")

	config_, err := config.LoadConfigFile(configFileName)
	if err != nil {
		log.Print(err)
		log.Fatal(err)
	}

	var logFileName string
	flag.StringVar(&logFileName, "logFileName", "", "path to server log file")

	var logFile *os.File
	if logFileName != "" {
		logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			_ = logFile.Close()
		}()

		log.SetOutput(logFile)
		defer func() {
			log.Println("--- SERVER STOPPED HERE ---")
			log.Println()
		}()
	}

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

	//validator := Validator.New()
	//validator.RegisterCustomTypeFunc(timestamps.ValidateDateTime, timestamps.DateTime{})

	echo_ := echo.New()
	//echo_.Logger.SetLevel(LabstackLog.ERROR)
	if logFile != nil {
		echo_.Logger.SetOutput(logFile)
	}
	echo_.Use(middlewaresManager.RecoverMiddleware.Recover())
	echo_.Validator = RequestValidator.NewRequestValidator(Validator.New())

	adsDelivery.Configure(echo_, middlewaresManager)
	sessionDelivery.Configure(echo_, middlewaresManager)
	userDelivery.Configure(echo_, middlewaresManager)

	if err := echo_.Start(config_.GetServerConfigString()); err != nil {
		log.Fatal(err)
	}
}
