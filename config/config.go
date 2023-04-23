package config

import "mahasiswa/helper"

type Configuration struct {
	DB_USERNAME string 
	DB_PASSWORD string 
	DB_PORT     string 
	DB_HOST     string 
	DB_NAME     string	
}

func GetConfig() Configuration {
	
	configuration := Configuration{
		DB_USERNAME: helper.GoDotEnv("DB_USERNAME"),
		DB_PASSWORD: helper.GoDotEnv("DB_PASSWORD"),
		DB_PORT:     helper.GoDotEnv("DB_PORT"),
		DB_HOST:     helper.GoDotEnv("DB_HOST"),
		DB_NAME:     helper.GoDotEnv("DB_NAME"),		
	}	

	return configuration
}
