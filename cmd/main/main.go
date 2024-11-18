package main

import (
	/* "log"
	"net/http" */


	"github.com/gofiber/fiber/v2"
	/* "github.com/rs/cors" */
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/codeknight05/GO-senior-care/pkg/config"
	"github.com/codeknight05/GO-senior-care/pkg/routes"
	



	_ "github.com/mattn/go-sqlite3"
	"encoding/json"
    "net/http"
)

func main() {

    config.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000", // Use "http" if your frontend is on HTTP
	}))
	routes.Setup(app)

	
	app.Listen(":9000")

	// Define the health check handler
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"status": "API is working!"})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Define the data receiving handler
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var requestData map[string]interface{}
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&requestData); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{"received": requestData})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}

