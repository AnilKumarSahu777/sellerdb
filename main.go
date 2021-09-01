package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	db "sellerdb/database"

	"github.com/gorilla/mux"
)

// REQUEST is struct of body
type REQUEST []struct {
	NAME         string `json:"name"`
	IMAGEURL     string `json:"imageURL"`
	DESCRIPTION  string `json:"description"`
	PRICE        string `json:"price"`
	TOTALREVIEWS string `json:"totalreviews"`
	URL          string `json:"url"`
}

// PRODUCT is struct of dbmodel
type PRODUCT struct {
	NAME         string    `json:"name"`
	IMAGEURL     string    `json:"imageURL"`
	DESCRIPTION  string    `json:"description"`
	PRICE        string    `json:"price"`
	TOTALREVIEWS string    `json:"totalreviews"`
	CREATEDATE   time.Time `json:"createdate"`
}

// DBMODEL is struct of database
type DBMODEL struct {
	URL     string  `json:"url"`
	PRODUCT PRODUCT `json:"product"`
}

func homepage(w http.ResponseWriter, r *http.Request) {
	var body REQUEST
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		return
	}
	var dbmodel DBMODEL
	var dbmodels []DBMODEL
	for _, Body := range body {
		dbmodel.PRODUCT.CREATEDATE = time.Now().UTC()
		dbmodel.PRODUCT.DESCRIPTION = Body.DESCRIPTION
		dbmodel.PRODUCT.IMAGEURL = Body.IMAGEURL
		dbmodel.PRODUCT.NAME = Body.NAME
		dbmodel.PRODUCT.PRICE = Body.PRICE
		dbmodel.PRODUCT.TOTALREVIEWS = Body.TOTALREVIEWS
		dbmodel.URL = Body.URL
		dbmodels = append(dbmodels, dbmodel)
	}
	err = CreateMany(dbmodels)
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := json.NewEncoder(w)
	enc.Encode(body)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter()
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homepage).Methods("POST")
	log.Fatal(http.ListenAndServe(":10091", myRouter))
}
func main() {
	handleRequests()
}

//CreateMany - Insert multiple documents at once in the collection.
func CreateMany(list []DBMODEL) error {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	//Get MongoDB connection using connectionhelper.
	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(db.DB).Collection(db.ISSUES)
	//Perform InsertMany operation & validate against the error.
	_, err = collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}
