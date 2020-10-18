package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zafiranursabila/intro-microservice/auth/handler"
	"log"
	"net/http"
)

func main() {
	cfg.err := fetConfig(); if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(cfg)
	}
	router := mux.NewRouter()

	router.Handle("/admin-auth", http.HandlerFunc(handler.ValidateAuth))

	fmt.Printf("Auth service listen on :8001")
	log.Panic(http.ListenAndServe(":8001", router))
}


func getConfig() (config.Config.error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetCinfigName("config.yml")

	if err:= viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initDB(cfg config.Database) (*gorm.DB,error){
	dsn := fat.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Config)
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		return nil, err
	}

	db. AutoMigrate(&database.Auth{})
	if err != nil {
		return nil, err
	}
	return db, nil
}