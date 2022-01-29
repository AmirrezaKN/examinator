package main

import (
	"database/sql"

	"github.com/AmirrezaKN/examinator/adapters/repository"
	"github.com/AmirrezaKN/examinator/api/rest"
	"github.com/AmirrezaKN/examinator/pkg/log"
	"github.com/AmirrezaKN/examinator/service/core"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	logger := log.NewZapLogger()

	db, err := sql.Open("postgres", appConfig.DB.toString())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	dbRepo := repository.NewPGRepository(db)
	examinatorService := core.NewExaminatorService(dbRepo)

	httpServer := echo.New()
	examinatorController := rest.NewRestController(logger, httpServer, examinatorService)
	logger.Errorf("%+v", examinatorController.Serve(appConfig.RestAdress))
}
