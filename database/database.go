package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"JobVacancyManagement/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// var connectionString string = "mongodb+srv://userFromDSi:Bze7UojNtnQ76fZs@cluster0.5010nc6.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
var connectionString string = "mongodb://localhost:27017"

type Database struct {
	client *mongo.Client
}

func Connect() *Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		client: client,
	}
}

func (database Database) CreateJobDescription(input model.JobDescriptionCreationInput) *model.JobDescription {
	jobCollection := database.client.Database("talentManagement").Collection("jobCollection")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	insertedDocument, error := jobCollection.InsertOne(ctx,
		bson.M{"positionName": input.PositionName, "skills": input.Skills, "yearsOfExperience": input.YearsOfExperience, "minSalary": input.MinSalary, "active": input.Active})

	fmt.Println(">>> >>> Inserted Data:", insertedDocument)

	if error != nil {
		log.Fatal(error)
	}

	insertedDocumentID := insertedDocument.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobDescription{ID: insertedDocumentID, PositionName: input.PositionName, YearsOfExperience: input.YearsOfExperience, Skills: input.Skills, Active: input.Active}
	return &returnJobListing
}

func (database Database) UpdateJobDescription(id string, input model.JobDescriptionUpdateInput) *model.JobDescription {
	jobCollection := database.client.Database("talentManagement").Collection("jobCollection")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	jobDescriptionUpdateInfo := bson.M{}

	if input.PositionName != nil {
		jobDescriptionUpdateInfo["positionName"] = input.PositionName
	}
	if input.Skills != nil {
		jobDescriptionUpdateInfo["skills"] = input.Skills
	}
	if input.YearsOfExperience != nil {
		jobDescriptionUpdateInfo["yearsOfExperience"] = input.YearsOfExperience
	}
	if input.MinSalary != nil {
		jobDescriptionUpdateInfo["minSalary"] = input.MinSalary
	}
	if input.Active != nil {
		jobDescriptionUpdateInfo["active"] = input.Active
	}

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	updateCommand := bson.M{"$set": jobDescriptionUpdateInfo}

	results := jobCollection.FindOneAndUpdate(ctx, filter, updateCommand, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobDescription model.JobDescription

	if err := results.Decode(&jobDescription); err != nil {
		log.Fatal(err)
	}

	return &jobDescription
}

func (database Database) DeleteJobDescription(id string) *model.JobDescriptionDeleteResponse {
	jobCollection := database.client.Database("talentManagement").Collection("jobCollection")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	_, err := jobCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.JobDescriptionDeleteResponse{DeletedJobDescriptionID: id}
}

func (database Database) GetJobDescriptions() []*model.JobDescription {
	jobCollection := database.client.Database("talentManagement").Collection("jobCollection")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var jobDescriptions []*model.JobDescription
	cursor, error := jobCollection.Find(ctx, bson.D{})
	if error != nil {
		log.Fatal(error)
	}

	// Iterate over each job description and make its ID visible
	for i := range jobDescriptions {
		// Convert the ObjectID to its hexadecimal representation
		hexId, _ := primitive.ObjectIDFromHex(jobDescriptions[i].ID)
		fmt.Println(">=-=-=-> Hex ID: ", hexId)
	}

	if error = cursor.All(context.TODO(), &jobDescriptions); error != nil {
		panic(error)
	}

	return jobDescriptions
}

func (database Database) GetJobDescription(id string) *model.JobDescription {
	jobCollection := database.client.Database("talentManagement").Collection("jobCollection")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobDescription model.JobDescription
	err := jobCollection.FindOne(ctx, filter).Decode(&jobDescription)
	if err != nil {
		log.Fatal(err)
	}
	return &jobDescription
}
