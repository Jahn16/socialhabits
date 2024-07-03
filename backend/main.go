package main

import (
	"log"

	"github.com/Jahn16/habitshare/database"
	"github.com/Jahn16/habitshare/handlers"

	"os"

	"github.com/Jahn16/habitshare/middlewares/auth0"
	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	db, _ := database.Setup()
	godotenv.Load()
	authMiddleware := auth0.New(auth0.Config{Issuer: os.Getenv("AUTH_AUTH0_DOMAIN"), Audience: []string{os.Getenv("AUTH_AUTH0_AUDIENCE")}})

	app.Get("/habits", handlers.HabitList(db))
	app.Get("/habits/:id", handlers.HabitGet(db))
	app.Post("/habits", authMiddleware, handlers.HabitCreate(db))
	app.Delete("/habits/:id", authMiddleware, handlers.DeleteHabit(db))
	app.Post("/habits/:id/record", authMiddleware, handlers.RecordHabit(db))
	app.Delete("/records/:id", authMiddleware, handlers.DeleteRecord(db))

	app.Get("/users", handlers.UserList(db))
	app.Post("/users", authMiddleware, handlers.UserCreate(db))

	app.Get("/users/me", authMiddleware, handlers.GetAuthenticatedUser(db))
	app.Get("/users/:id", handlers.UserGet(db))

	log.Fatal(app.Listen(":8000"))
}
