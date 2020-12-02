package main

import (
	"context"
	"database/database/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("swapi")
	var films structs.Films
	filmsC := ReadIn("films", "./database/data/films.json", &films, db)
	for _, f := range films {
		ir, err := filmsC.InsertOne(ctx, f)
		if err != nil {
			panic(err)
		}
		fmt.Println(ir.InsertedID)
	}
	var people structs.People
	peopleC := ReadIn("people", "./database/data/people.json", &people, db)
	for _, f := range people {
		ir, err := peopleC.InsertOne(ctx, f)
		if err != nil {
			panic(err)
		}
		fmt.Println(ir.InsertedID)
	}
	var planets structs.Planets
	planetsC := ReadIn("planets", "./database/data/planets.json", &planets, db)
	for _, f := range planets {
		ir, err := planetsC.InsertOne(ctx, f)
		if err != nil {
			panic(err)
		}
		fmt.Println(ir.InsertedID)
	}
	var species structs.Species
	speciesC := ReadIn("species", "./database/data/species.json", &species, db)
	for _, f := range species {
		ir, err := speciesC.InsertOne(ctx, f)
		if err != nil {
			panic(err)
		}
		fmt.Println(ir.InsertedID)
	}
	var starships structs.Starships
	starshipsC := ReadIn("startships", "./database/data/Starships.json", &starships, db)
	for _, f := range starships {
		ir, err := starshipsC.InsertOne(ctx, f)
		if err != nil {
			panic(err)
		}
		fmt.Println(ir.InsertedID)
	}
	var vehicles structs.Vehicles
	vehiclesC := ReadIn("vehicles", "./database/data/vehicles.json", &vehicles, db)
	for _, f := range vehicles {
		ir, err := vehiclesC.InsertOne(ctx, f)
		if err != nil {
			panic(err)
		}
		fmt.Println(ir.InsertedID)
	}
}

// ReadIn takes a jsonfile and a matching empty interface and fills it with data.
func ReadIn(coll string, filename string, inf interface{}, db *mongo.Database) *mongo.Collection {
	datafile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer datafile.Close()
	b, err := ioutil.ReadAll(datafile)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(b, inf)
	c := db.Collection(coll)
	return c
}
