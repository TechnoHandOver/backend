package main

import (
	"database/sql"
	"flag"
	"github.com/TechnoHandOver/backend/config"
	AdsDelivery "github.com/TechnoHandOver/backend/internal/ads/delivery"
	AdsRepository "github.com/TechnoHandOver/backend/internal/ads/repository"
	AdsUsecase "github.com/TechnoHandOver/backend/internal/ads/usecase"
	UserDelivery "github.com/TechnoHandOver/backend/internal/user/delivery"
	UserRepository "github.com/TechnoHandOver/backend/internal/user/repository"
	UserUsecase "github.com/TechnoHandOver/backend/internal/user/usecase"
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

	if logFileName != "" {
		logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		log.Print(err)
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Print(err)
			log.Fatal(err)
		}
	}()

	if err := db.Ping(); err != nil {
		log.Print(err)
		log.Fatal(err)
	}

	log.Println(config_.GetDatabaseConfigString())

	adsRepository := AdsRepository.NewAdsRepository(db)
	userRepository := UserRepository.NewUserRepository(db)

	adsUsecase := AdsUsecase.NewAdsUsecase(adsRepository)
	userUsecase := UserUsecase.NewUserUsecase(userRepository)

	adsDelivery := AdsDelivery.NewAdsDelivery(adsUsecase)
	userDelivery := UserDelivery.NewUserDelivery(userUsecase)

	echo_ := echo.New()

	adsDelivery.Configure(echo_)
	userDelivery.Configure(echo_)

	if err := echo_.Start(config_.GetServerConfigString()); err != nil {
		log.Print(err)
		log.Fatal(err)
	}
}
