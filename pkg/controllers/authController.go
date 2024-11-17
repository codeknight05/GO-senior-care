package controllers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Health Monitoring APIs
func GetUsers(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT user_id, name, email, role, date_created FROM Users")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		var users []fiber.Map
		for rows.Next() {
			var id int
			var name, email, role, dateCreated string
			if err := rows.Scan(&id, &name, &email, &role, &dateCreated); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			users = append(users, fiber.Map{
				"user_id": id, "name": name, "email": email, "role": role, "date_created": dateCreated,
			})
		}
		return c.Status(fiber.StatusOK).JSON(users)
	}
}

// Add User to the database
func AddUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Hash the password before storing it
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Insert the new user into the Users table
		_, err = db.Exec("INSERT INTO Users (name, email, password_hash, role) VALUES (?, ?, ?, ?)",
			req.Name, req.Email, passwordHash, req.Role)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "User added"})
	}
}

// Get Medications for a User
func GetMedications(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Params("user_id")

		// Query Medications table based on user_id
		rows, err := db.Query("SELECT medication_id, medication_name, dosage, frequency, start_date, end_date FROM Medications WHERE user_id = ?", userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		var medications []fiber.Map
		for rows.Next() {
			var id int
			var name, dosage, frequency, startDate, endDate string
			if err := rows.Scan(&id, &name, &dosage, &frequency, &startDate, &endDate); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			medications = append(medications, fiber.Map{
				"medication_id": id, "medication_name": name, "dosage": dosage, "frequency": frequency,
				"start_date": startDate, "end_date": endDate,
			})
		}
		return c.Status(fiber.StatusOK).JSON(medications)
	}
}

// Add Medication for a User
func AddMedication(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			UserID          int    `json:"user_id"`
			MedicationName  string `json:"medication_name"`
			Dosage          string `json:"dosage"`
			Frequency       string `json:"frequency"`
			StartDate       string `json:"start_date"`
			EndDate         string `json:"end_date"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Insert new medication into Medications table
		_, err := db.Exec("INSERT INTO Medications (user_id, medication_name, dosage, frequency, start_date, end_date) VALUES (?, ?, ?, ?, ?, ?)",
			req.UserID, req.MedicationName, req.Dosage, req.Frequency, req.StartDate, req.EndDate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Medication added"})
	}
}

// Get Sleep Patterns for a User
func GetSleepPatterns(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Params("user_id")

		// Query SleepPatterns table based on user_id
		rows, err := db.Query("SELECT sleep_pattern_id, sleep_start, sleep_end, duration FROM SleepPatterns WHERE user_id = ?", userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		var patterns []fiber.Map
		for rows.Next() {
			var id int
			var sleepStart, sleepEnd string
			var duration int
			if err := rows.Scan(&id, &sleepStart, &sleepEnd, &duration); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			patterns = append(patterns, fiber.Map{
				"sleep_pattern_id": id, "sleep_start": sleepStart, "sleep_end": sleepEnd, "duration": duration,
			})
		}
		return c.Status(fiber.StatusOK).JSON(patterns)
	}
}

// Add Sleep Pattern for a User
func AddSleepPattern(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			UserID    int    `json:"user_id"`
			SleepStart string `json:"sleep_start"`
			SleepEnd   string `json:"sleep_end"`
			Duration   int    `json:"duration"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Insert new sleep pattern into SleepPatterns table
		_, err := db.Exec("INSERT INTO SleepPatterns (user_id, sleep_start, sleep_end, duration) VALUES (?, ?, ?, ?)",
			req.UserID, req.SleepStart, req.SleepEnd, req.Duration)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Sleep pattern added"})
	}
}

// Get Status Updates for a User
func GetStatusUpdates(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Params("user_id")

		// Query StatusUpdates table based on user_id
		rows, err := db.Query("SELECT status_update_id, caregiver_id, update_time, status_message FROM StatusUpdates WHERE user_id = ?", userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		var statusUpdates []fiber.Map
		for rows.Next() {
			var id, caregiverID int
			var updateTime, statusMessage string
			if err := rows.Scan(&id, &caregiverID, &updateTime, &statusMessage); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			statusUpdates = append(statusUpdates, fiber.Map{
				"status_update_id": id, "caregiver_id": caregiverID, "update_time": updateTime, "status_message": statusMessage,
			})
		}
		return c.Status(fiber.StatusOK).JSON(statusUpdates)
	}
}

// Add Status Update for a User
func AddStatusUpdate(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			UserID        int    `json:"user_id"`
			CaregiverID   int    `json:"caregiver_id"`
			StatusMessage string `json:"status_message"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Insert new status update into StatusUpdates table
		_, err := db.Exec("INSERT INTO StatusUpdates (user_id, caregiver_id, status_message) VALUES (?, ?, ?)",
			req.UserID, req.CaregiverID, req.StatusMessage