package main

import "fmt"

type DBConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	DBName          string
	SSLMode         string
	MigrationsPath  string
	MigrationsTable string
}

func (c *DBConfig) toString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.User, c.Pass, c.DBName, c.SSLMode)
}

type ExaminatorConfig struct {
	RestAdress string
	DB         DBConfig
}

var appConfig = ExaminatorConfig{
	RestAdress: "0.0.0.0:6789",
	DB: DBConfig{
		Host:            "127.0.0.1",
		Port:            "5432",
		User:            "amir",
		Pass:            "1",
		DBName:          "amir",
		SSLMode:         "disable",
		MigrationsPath:  "database/migrations",
		MigrationsTable: "migrations_ganje",
	},
}
