package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

type DB struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

type Product struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"na"`
	brand    string        `json:"brand" bson:"ba"`
	Variants []Variant     `json:"variants" bson:"va"`
}

type Variant struct {
	Description string  `json:"description" bson:"ds"`
	Sku         string  `json:"sku" bson:"sk"`
	Price       float32 `json:"price" bson:"pr"`
	Stock       float32 `json:"stock" bson:"st"`
}

func (db *DB) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	producID := bson.ObjectIdHex(params["id"])
	var product Product
	err := db.Collection.Find(bson.M{"_id": producID}).One(&product)
	if err != nil {
		panic(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func (db *DB) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var products []Product
	err := db.Collection.Find(nil).All(&products)
	if err != nil {
		panic(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func (db *DB) CreateProduct(w http.ResponseWriter, r *http.Request) {
	jsonBody := json.NewDecoder(r.Body)
	var product Product
	err := jsonBody.Decode(&product)
	if err == nil {
		db.Collection.Insert(product)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	} else {
		panic(err)
	}
}

func (db *DB) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	jsonBody := json.NewDecoder(r.Body)
	params := mux.Vars(r)
	producID := bson.ObjectIdHex(params["id"])
	var product Product
	err := jsonBody.Decode(&product)
	if err == nil {
		err = db.Collection.Update(bson.M{"_id": producID}, bson.M{"$set": &product})
		if err != nil {
			panic(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
		}
	} else {
		panic(err)
	}
}

func (db *DB) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	producID := bson.ObjectIdHex(params["id"])
	err := db.Collection.Remove(bson.M{"_id": producID})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	} else {
		panic(err)
	}

}

func main() {
	session, err := mgo.Dial(GetConnectionString())
	c := session.DB("ecommerce").C("products")
	db := &DB{Session: session, Collection: c}
	if err != nil {
		panic(err)
	}
	defer session.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", db.GetAllProduct).Methods("GET")
	router.HandleFunc("/products/{id}", db.GetProductByID).Methods("GET")
	router.HandleFunc("/products", db.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", db.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", db.DeleteProduct).Methods("DELETE")

	server := http.ListenAndServe(Port(), router)
	log.Fatal(server)
}

func GetConnectionString() string {
	connectionString := os.Getenv("DB")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}
	return connectionString
}

func Port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"

	}
	return ":" + port
}
