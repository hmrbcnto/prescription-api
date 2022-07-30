package drug_repo

import (
	"context"
	"time"

	"github.com/hmrbcnto/prescription-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type DrugRepo interface {
	GetDrugs() ([]entities.Drug, error)
	GetDrugById(string) (*entities.Drug, error)
	CreateDrug(*entities.Drug) (*entities.Drug, error)
	DeleteDrugById(string) (error)
	UpdateDrugById(string, *entities.Drug) (*entities.Drug, error)
}

type drugRepo struct {
	db *mongo.Collection
}

func NewRepo(db *mongo.Client) DrugRepo {
	return &drugRepo{
		db: db.Database("prescription-api").Collection("drugs"),
	}
}

func (dr *drugRepo) GetDrugs() ([]entities.Drug, error) {
	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	// Creating query
	query := bson.D{{}}
	
	// Querying
	cursor, err := dr.db.Find(ctx, query)

	if err != nil {
		return nil, err;
	}

	// Decode to object
	var drugs []entities.Drug

	// Iterate and decode
	err = cursor.All(ctx, drugs)

	if err != nil {
		return nil, err
	}

	return drugs, nil
}

func (dr *drugRepo) GetDrugById(id string) (*entities.Drug, error) {
	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	// Creating query
	query := bson.D{{ Key: "_id", Value: id}}

	// Querying
	result := dr.db.FindOne(ctx, query)
	err := result.Err()

	if err != nil {
		return nil, err
	}

	// Converting into entity
	foundDrug := new(entities.Drug)
	result.Decode(foundDrug);

	return foundDrug, nil
}

func (dr *drugRepo) CreateDrug(drug *entities.Drug) (*entities.Drug, error) {
	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	// Creating drug
	insertionResult, err := dr.db.InsertOne(ctx, drug)

	if err != nil {
		return nil, err
	}

	// Getting inserted data
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

	// Query
	createdRecord := dr.db.FindOne(ctx, filter)
	err = createdRecord.Err()

	if err != nil {
		return nil, err
	}
	
	// Decoding
	createdDrug := new(entities.Drug)
	createdRecord.Decode(createdDrug)

	return createdDrug, nil
}

func (dr *drugRepo) DeleteDrugById(id string) (error) {
	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	// Creating query
	query := bson.D{{Key: "_id", Value: id}}

	// Deleting
	deleteResult := dr.db.FindOneAndDelete(ctx, query)
	err := deleteResult.Err()
	
	if err != nil {
		return err
	}

	return nil
}

func (dr *drugRepo) UpdateDrugById(id string, drug *entities.Drug) (*entities.Drug, error) {
	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	
	defer cancel();

	// Creating filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Update filter
	result, err := dr.db.UpdateOne(ctx, filter, drug)

	if err != nil {
		return nil, err
	}

	// Get updated document
	query := bson.D{{Key: "_id", Value: result.UpsertedID}}

	updatedDrug := dr.db.FindOne(ctx, query)
	err = updatedDrug.Err()

	if err != nil {
		return nil, err
	}

	// Decoding result
	foundDrug := new(entities.Drug)
	updatedDrug.Decode(foundDrug)

	return foundDrug, err;
}