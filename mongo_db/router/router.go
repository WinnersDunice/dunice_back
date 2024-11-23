package router

import (
	"encoding/json"
	"log"
	"net/http"

	fn "github.com/WinnersDunice/dunice_back/mongo_db/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Rout(client *mongo.Client) error {
	tableCollection := client.Database("mongo").Collection("tables")
	chairCollection := client.Database("mongo").Collection("chairs")
	equipmentCollection := client.Database("mongo").Collection("equipment")
	furnitureCollection := client.Database("mongo").Collection("furniture")
	kitchenCollection := client.Database("mongo").Collection("kitchen")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling root request")
		w.Write([]byte("never should have pwn here, FSB went 4 u <3"))
	})

	api := chi.NewRouter()
	r.Mount("/mongo", api)

	// Tables routes
	v1 := chi.NewRouter()
	api.Mount("/tables", v1)

	// Create a new table
	v1.Post("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling POST /tables request")
		table := new(fn.Table)
		if err := json.NewDecoder(r.Body).Decode(table); err != nil {
			log.Printf("Error decoding table: %v", err)
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		table.ID = primitive.NewObjectID().Hex()
		if err := fn.CreateTable(tableCollection, table); err != nil {
			log.Printf("Error creating table: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Table created: %v", table)
		json.NewEncoder(w).Encode(table)
	})

	// Chairs routes
	v2 := chi.NewRouter()
	api.Mount("/chairs", v2)

	// Create a new chair
	v2.Post("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling POST /chairs request")
		chair := new(fn.Chair)
		if err := json.NewDecoder(r.Body).Decode(chair); err != nil {
			log.Printf("Error decoding chair: %v", err)
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		chair.ID = primitive.NewObjectID().Hex()
		if err := fn.CreateChair(chairCollection, chair); err != nil {
			log.Printf("Error creating chair: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Chair created: %v", chair)
		json.NewEncoder(w).Encode(chair)
	})

	// Equipment routes
	v3 := chi.NewRouter()
	api.Mount("/equipment", v3)

	// Create a new equipment
	v3.Post("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling POST /equipment request")
		equipment := new(fn.Equipment)
		if err := json.NewDecoder(r.Body).Decode(equipment); err != nil {
			log.Printf("Error decoding equipment: %v", err)
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		equipment.ID = primitive.NewObjectID().Hex()
		if err := fn.CreateEquipment(equipmentCollection, equipment); err != nil {
			log.Printf("Error creating equipment: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Equipment created: %v", equipment)
		json.NewEncoder(w).Encode(equipment)
	})

	// Furniture routes
	v4 := chi.NewRouter()
	api.Mount("/furniture", v4)

	// Create a new furniture
	v4.Post("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling POST /furniture request")
		furniture := new(fn.Furniture)
		if err := json.NewDecoder(r.Body).Decode(furniture); err != nil {
			log.Printf("Error decoding furniture: %v", err)
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		furniture.ID = primitive.NewObjectID().Hex()
		if err := fn.CreateFurniture(furnitureCollection, furniture); err != nil {
			log.Printf("Error creating furniture: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Furniture created: %v", furniture)
		json.NewEncoder(w).Encode(furniture)
	})

	// Kitchen routes
	v5 := chi.NewRouter()
	api.Mount("/kitchen", v5)

	// Create a new kitchen
	v5.Post("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling POST /kitchen request")
		kitchen := new(fn.Kitchen)
		if err := json.NewDecoder(r.Body).Decode(kitchen); err != nil {
			log.Printf("Error decoding kitchen: %v", err)
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		kitchen.ID = primitive.NewObjectID().Hex()
		if err := fn.CreateKitchen(kitchenCollection, kitchen); err != nil {
			log.Printf("Error creating kitchen: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Kitchen created: %v", kitchen)
		json.NewEncoder(w).Encode(kitchen)
	})

	// Get all objects by BelongsTo
	api.Get("/belongsTo/{belongsTo}", func(w http.ResponseWriter, r *http.Request) {
		belongsTo := chi.URLParam(r, "belongsTo")
		log.Printf("Handling GET /belongsTo/%s request", belongsTo)
		objects, err := fn.GetObjectsByBelongsTo(client, belongsTo)
		if err != nil {
			log.Printf("Error getting objects by BelongsTo: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(objects)
	})

	// Get all objects by OfficeID
	api.Get("/officeId/{officeId}", func(w http.ResponseWriter, r *http.Request) {
		officeId := chi.URLParam(r, "officeId")
		log.Printf("Handling GET /officeId/%s request", officeId)
		objects, err := fn.GetObjectsByOfficeID(client, officeId)
		if err != nil {
			log.Printf("Error getting objects by OfficeID: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(objects)
	})

	// Delete object by ObjectID
	api.Delete("/{collection}/{id}", func(w http.ResponseWriter, r *http.Request) {
		collectionName := chi.URLParam(r, "collection")
		id := chi.URLParam(r, "id")
		log.Printf("Handling DELETE /%s/%s request", collectionName, id)
		err := fn.DeleteObjectByID(client, collectionName, id)
		if err != nil {
			log.Printf("Error deleting object by ID: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	return http.ListenAndServe(":8004", r)
}
