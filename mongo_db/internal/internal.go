package internal

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Point struct {
	X float64 `json:"x" bson:"x"`
	Y float64 `json:"y" bson:"y"`
}

type Table struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	Type      string  `json:"type" bson:"type"`
	Points    []Point `json:"points" bson:"points"`
	BelongsTo string  `json:"belongsTo" bson:"belongsTo"`
	OfficeID  string  `json:"officeId" bson:"officeId"`
}

type Chair struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Type      string `json:"type" bson:"type"`
	Position  Point  `json:"position" bson:"position"`
	BelongsTo string `json:"belongsTo" bson:"belongsTo"`
	OfficeID  string `json:"officeId" bson:"officeId"`
}

type Equipment struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	Type      string  `json:"type" bson:"type"`
	Position  Point   `json:"position" bson:"position"`
	Radius    float64 `json:"radius,omitempty" bson:"radius,omitempty"`
	BelongsTo string  `json:"belongsTo" bson:"belongsTo"`
	OfficeID  string  `json:"officeId" bson:"officeId"`
}

type Furniture struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	Type      string  `json:"type" bson:"type"`
	Points    []Point `json:"points" bson:"points"`
	BelongsTo string  `json:"belongsTo" bson:"belongsTo"`
	OfficeID  string  `json:"officeId" bson:"officeId"`
}

type Kitchen struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	Type      string  `json:"type" bson:"type"`
	Position  Point   `json:"position" bson:"position"`
	Points    []Point `json:"points,omitempty" bson:"points,omitempty"`
	BelongsTo string  `json:"belongsTo" bson:"belongsTo"`
	OfficeID  string  `json:"officeId" bson:"officeId"`
}

// CreateTable создает новый стол
func CreateTable(collection *mongo.Collection, table *Table) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, table)
	return err
}

// CreateChair создает новое кресло
func CreateChair(collection *mongo.Collection, chair *Chair) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, chair)
	return err
}

// CreateEquipment создает новое оборудование
func CreateEquipment(collection *mongo.Collection, equipment *Equipment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, equipment)
	return err
}

// CreateFurniture создает новую мебель
func CreateFurniture(collection *mongo.Collection, furniture *Furniture) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, furniture)
	return err
}

// CreateKitchen создает новую кухню
func CreateKitchen(collection *mongo.Collection, kitchen *Kitchen) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, kitchen)
	return err
}

// GetObjectsByBelongsTo получает все объекты по полю BelongsTo
func GetObjectsByBelongsTo(client *mongo.Client, belongsTo string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collections := []string{"tables", "chairs", "equipment", "furniture", "kitchen"}
	var results []bson.M

	for _, collectionName := range collections {
		collection := client.Database("mongo").Collection(collectionName)
		filter := bson.M{"belongsTo": belongsTo}
		cur, err := collection.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
		defer cur.Close(ctx)

		var docs []bson.M
		if err := cur.All(ctx, &docs); err != nil {
			return nil, err
		}
		results = append(results, docs...)
	}

	return results, nil
}

// GetObjectsByOfficeID получает все объекты по полю OfficeID
func GetObjectsByOfficeID(client *mongo.Client, officeID string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collections := []string{"tables", "chairs", "equipment", "furniture", "kitchen"}
	var results []bson.M

	for _, collectionName := range collections {
		collection := client.Database("mongo").Collection(collectionName)
		filter := bson.M{"officeId": officeID}
		cur, err := collection.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
		defer cur.Close(ctx)

		var docs []bson.M
		if err := cur.All(ctx, &docs); err != nil {
			return nil, err
		}
		results = append(results, docs...)
	}

	return results, nil
}

// DeleteObjectByID удаляет объект по ObjectID
func DeleteObjectByID(client *mongo.Client, collectionName, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := client.Database("mongo").Collection(collectionName)
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(ctx, filter)
	return err
}
type SVGImage struct {
    Data   []byte `json:"data" bson:"data"`
    Name   string `json:"name" bson:"name"`
    OfficeID string `json:"officeId" bson:"officeId"`
}

// CreateSVGImage создает новое SVG-изображение
func CreateSVGImage(collection *mongo.Collection, svgImage *SVGImage) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, svgImage)
    return err
}

// GetSVGImageByID получает SVG-изображение по ID
func GetSVGImageByID(collection *mongo.Collection, id string) (*SVGImage, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    filter := bson.M{"_id": objectID}
    var svgImage SVGImage
    err = collection.FindOne(ctx, filter).Decode(&svgImage)
    if err != nil {
        return nil, err
    }

    return &svgImage, nil
}
