package main

import (
	"fmt"
)

func main() {
	app := gin.Default()

	godotenv.Load()

	//Starting database connection
	database.ConnectDB()

	fmt.Println("Server started at port 4000")
	
	routes.SetupRoutes(app)

	// Start and run the server
	app.Run(":4000")

}
