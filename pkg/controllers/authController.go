package controllers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Health Monitoring APIs
func GetMedications(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, name, time FROM medications")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		var medications []fiber.Map
		for rows.Next() {
			var id int
			var name, time string
			if err := rows.Scan(&id, &name, &time); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			medications = append(medications, fiber.Map{"id": id, "name": name, "time": time})
		}
		return c.Status(fiber.StatusOK).JSON(medications)
	}
}

func AddMedication(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Name string `json:"name"`
			Time string `json:"time"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		_, err := db.Exec("INSERT INTO medications (name, time) VALUES (?, ?)", req.Name, req.Time)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Medication added"})
	}
}

// Caregiver Support APIs
func GetCaregiverStatus(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, name, status FROM caregiver_status")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		var statuses []fiber.Map
		for rows.Next() {
			var id int
			var name, status string
			if err := rows.Scan(&id, &name, &status); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			statuses = append(statuses, fiber.Map{"id": id, "name": name, "status": status})
		}
		return c.Status(fiber.StatusOK).JSON(statuses)
	}
}

func UpdateCaregiverStatus(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		_, err := db.Exec("INSERT INTO caregiver_status (name, status) VALUES (?, ?)", req.Name, req.Status)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Status updated"})
	}
}

// System Administration APIs
func GetUsers(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, name, role FROM users")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		var users []fiber.Map
		for rows.Next() {
			var id int
			var name, role string
			if err := rows.Scan(&id, &name, &role); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			users = append(users, fiber.Map{"id": id, "name": name, "role": role})
		}
		return c.Status(fiber.StatusOK).JSON(users)
	}
}

func AddUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Name string `json:"name"`
			Role string `json:"role"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		_, err := db.Exec("INSERT INTO users (name, role) VALUES (?, ?)", req.Name, req.Role)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "User added"})
	}
}

// Emergency Response APIs
func SendEmergencyAlert() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Message string `json:"message"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		log.Printf("Emergency alert sent: %s", req.Message)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Alert sent"})
	}
}
