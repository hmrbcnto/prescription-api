package doctor_repo

import (
	"context"
	"time"

	"github.com/hmrbcnto/prescription-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DoctorRepo interface {
	CreateDoctor(*entities.Doctor) (*entities.Doctor, error)
	GetAllDoctors() ([]entities.Doctor, error)
	GetDoctorById(string) (*entities.Doctor, error)
}

type doctorRepo struct {
	db *mongo.Collection
}

func NewRepo(db *mongo.Client) DoctorRepo {
	return &doctorRepo{
		db: db.Database("prescription-api").Collection("doctors"),
	}
}

func (dr *doctorRepo) CreateDoctor(doctor *entities.Doctor) (*entities.Doctor, error) {

	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	// Creating doctor
	insertionResult, err := dr.db.InsertOne(ctx, doctor)

	if err != nil {
		return nil, err
	}

	// Getting inserted data
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

	// Query
	createdRecord := dr.db.FindOne(ctx, filter)

	// Decode to doctor entity
	createdDoctor := &entities.Doctor{}
	createdRecord.Decode(createdDoctor)

	// returning
	return createdDoctor, nil
}

func (dr *doctorRepo) GetAllDoctors() ([]entities.Doctor, error) {

	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	/// Getting all doctors
	// Creating query
	query := bson.D{{}}

	cursor, err := dr.db.Find(ctx, query)

	if err != nil {
		return nil, err
	}

	var doctors []entities.Doctor

	// Iterate and decode
	err = cursor.All(ctx, &doctors)

	if err != nil {
		return nil, err
	}

	return doctors, nil
}

func (dr *doctorRepo) GetDoctorById(doctorId string) (*entities.Doctor, error) {

	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()
	// Creating query
	query := bson.D{{Key: "_id", Value: doctorId}}

	// Querying
	result := dr.db.FindOne(ctx, query)

	// Converting to user entity
	foundUser := new(entities.Doctor)
	result.Decode(foundUser)

	// Returning
	return foundUser, nil
}
