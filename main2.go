package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type game struct {
	gorm.Model

	Name string

	License string

	Cars []Car
}

type Car struct {
	gorm.Model

	Year int

	Make string

	ModelName string

	DriverID int
}

var db *gorm.DB

var err error

var (
	drivers = []Driver{

		{Name: "ACCOUNT", License: "MIT"}, // mit????????
    
	}

	cars = []Car{

		{Year: 2000, Make: "", ModelName: "", DriverID: 1},

		{Year: 2001, Make: "", ModelName: "", DriverID: 1},
	}
)

func GetCars(w http.ResponseWriter, r *http.Request) {

	var cars []Car

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}

func GetCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var car Car

	db.First(&car, params["id"])

	json.NewEncoder(w).Encode(&car)

}

func GetDriver(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var driver Driver

	var cars []Car

	db.First(&driver, params["id"])

	db.Model(&driver).Related(&cars)

	driver.Cars = cars

	json.NewEncoder(w).Encode(&driver)

}

func DeleteMatchmaking(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var car Car

	db.First(&car, params["id"])

	db.Delete(&car)

	var cars []Car

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}

func main() {

	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=localhost port=4000 user=postgres dbname=shashank sslmode=disable password=mobcoder@123")

	if err != nil {

		panic("failed to connect database")
	} else {

		fmt.Println("Successfully connected to the database")
	}

	defer db.Close()

	db.AutoMigrate(&Driver{})

	db.AutoMigrate(&Car{})

	for index := range cars {

		db.Create(&cars[index])

	}

	for index := range drivers {

		db.Create(&drivers[index])

	}

	router.HandleFunc("/playlist", Getplaylist).Methods("GET")

	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")

	router.HandleFunc("/party/{id}", GetParty).Methods("GET")

	router.HandleFunc("/matchmaking/{id}", DeleteMatchMaking).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":4000", handler))

}
