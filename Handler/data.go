package data

import (
  "https://mrcuda/game/yala/database"
)

func data() {
	app := fiber.New()
	client := database.Connect("", "mongodb://localhost:4000") // localhost
	defer database.Disconnect(client) // Disconnecting once the main finished execution
	routes.RegisterRoutes(app)

	app.Listen(":4000")

}
