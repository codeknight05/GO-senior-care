package routes

import (
	"database/sql"

	"github.com/codeknight05/GO-senior-care/pkg/config"
	"github.com/codeknight05/GO-senior-care/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)
var Db *sql.DB

func Setup(app *fiber.App) {

	Db=config.GetDB()
	app.Get("/health/medications", controllers.GetMedications(Db))
	app.Post("/health/medications",controllers.AddMedication(Db))
	app.Get("/caregiver/status", controllers.GetCaregiverStatus(Db))
	app.Post("/caregiver/status", controllers.UpdateCaregiverStatus(Db))
	app.Get("/admin/users", controllers.GetUsers(Db))
	app.Post("/admin/users", controllers.AddUser(Db))
    app.Post("/emergency/alerts", controllers.SendEmergencyAlert())
}

