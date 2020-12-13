package main

import (
	"context"
	"database/database/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	db := client.Database("swapi")
	err = InputData(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
}

// InputData Safely puts the swapi data into the DB
func InputData(ctx context.Context, db *mongo.Database) error {
	var films structs.Films
	filmsC := ReadIn("films", "./database/data/films.json", &films, db)
	for _, f := range films {
		ir, err := filmsC.InsertOne(ctx, f)
		if err != nil {
			fmt.Printf("%s - %s\n", f.Fields.Title, "Insert Error")
			return err
		}
		fmt.Printf("%s - %s\n", f.Fields.Title, ir.InsertedID)
	}
	var people structs.People
	peopleC := ReadIn("people", "./database/data/people.json", &people, db)
	for _, f := range people {
		ir, err := peopleC.InsertOne(ctx, f)
		if err != nil {
			fmt.Printf("%s - %s\n", f.Fields.Name, "Insert Error")
			return err
		}
		fmt.Printf("%s - %s\n", f.Fields.Name, ir.InsertedID)
	}
	var planets structs.Planets
	planetsC := ReadIn("planets", "./database/data/planets.json", &planets, db)
	for _, f := range planets {
		ir, err := planetsC.InsertOne(ctx, f)
		if err != nil {
			fmt.Printf("%s - %s\n", f.Fields.Name, "Insert Error")
			return err
		}
		fmt.Printf("%s - %s\n", f.Fields.Name, ir.InsertedID)
	}
	var species structs.Species
	speciesC := ReadIn("species", "./database/data/species.json", &species, db)
	for _, f := range species {
		ir, err := speciesC.InsertOne(ctx, f)
		if err != nil {
			fmt.Printf("%s - %s\n", f.Fields.Name, "Insert Error")
			return err
		}
		fmt.Printf("%s - %s\n", f.Fields.Name, ir.InsertedID)
	}
	var transport structs.Transport
	transportC := ReadIn("transport", "./database/data/transport.json", &transport, db)
	for _, f := range transport {
		ir, err := transportC.InsertOne(ctx, f)
		if err != nil {
			fmt.Printf("%s - %s\n", f.Fields.Name, "Insert Error")
			return err
		}
		fmt.Printf("%s - %s\n", f.Fields.Name, ir.InsertedID)
	}
	var starships structs.Starships
	starshipsC := ReadIn("startships", "./database/data/starships.json", &starships, db)
	for _, f := range starships {
		ir, err := starshipsC.InsertOne(ctx, f)
		if err != nil {
			fmt.Printf("%s - %s\n", f.Fields.Pilots, "Insert Error")
			return err
		}
		fmt.Printf("%s - %s\n", f.Fields.Pilots, ir.InsertedID)
	}
	var vehicles structs.Vehicles
	vehiclesC := ReadIn("vehicles", "./database/data/vehicles.json", &vehicles, db)
	for _, f := range vehicles {
		ir, err := vehiclesC.InsertOne(ctx, f)
		if err != nil {
			fmt.Printf("%s - %s\n", f.Fields.Pilots, "Insert Error")
			return err
		}
		fmt.Printf("%s - %s\n", f.Fields.Pilots, ir.InsertedID)
	}
	return nil
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
