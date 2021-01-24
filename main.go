package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int     `json:"id,omitempty" db:"ID"`
	Name        string  `json:"name,omitempty" db:"Name"`
	CountryCode string  `json:"countryCode,omitempty" db:"CountryCode"`
	District    string  `json:"district,omitempty" db:"District"`
	Population  int     `json:"population,omitempty" db:"Population"`
}

type Country struct {
	Code       string `json:"code,omitempty"  db:"Code"`
	Name       string `json:"name,omitempty"  db:"Name"`
	Continent  string `json:"continent,omitempty"  db:"Continent"`
	Population int    `json:"population,omitempty"  db:"Population"`
}

func main() {
	db, err := sqlx.Connect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	fmt.Println("Connected!")

	city := City{}
	_ = db.Get(&city, "SELECT * FROM city WHERE Name = '" + os.Args[1] + "' LIMIT 1")

	country := Country{}
	_ = db.Get(&country, "SELECT Code, Population FROM country WHERE Code = '" + city.CountryCode + "' LIMIT 1")

	fmt.Printf("%vの人口割合は%vです\n", os.Args[1], float64(city.Population) / float64(country.Population))
}

