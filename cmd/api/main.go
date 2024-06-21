package main

import (
	"log"
	"net/http"
	"os"

	model "go-api/internal/domain/model/user"
	"go-api/pkg/router"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Message struct {
	Message string `json:"message"`
}

var db *gorm.DB

func initDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" dbname=" + os.Getenv("DB_NAME") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE") +
		" TimeZone=America/Sao_Paulo"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	models := []interface{}{
		&model.User{},
	}

	// Executar migrações automáticas para cada modelo
	for _, model := range models {
		err = db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("failed to migrate model %T: %v", model, err)
		}
	}
}

func main() {
	initDB()
	router := router.NewRouter(db)

	log.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", router)
}
