package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Instance *MongoDatabase

func init() {
	Instance = NewMongoDataBase()
	Instance.connectDB()
}

type MongoDatabase struct {
	client          *mongo.Client
	projectDB       *mongo.Database
	usersCollection *mongo.Collection
	ctx             context.Context
}

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FirstName  string             `bson:"first_name"`
	LastName   string             `bson:"last_name"`
	Email      string             `bson:"email"`
	Age        int                `bson:"age"`
	Balance    int                `bson:"balance"`
	Password   string             `bson:"password"`
	Permission string             `bson:"permission"`
}

func NewMongoDataBase() *MongoDatabase {
	return &MongoDatabase{}
}

func (databaseInstance *MongoDatabase) connectDB() {
	clientOptions := options.Client()
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	databaseInstance.client = client
	databaseInstance.projectDB = client.Database("bankingsystem")
	databaseInstance.usersCollection = databaseInstance.projectDB.Collection("users")
	databaseInstance.ctx = context.TODO()
}

func (databaseInstance *MongoDatabase) CreateNewUser(user User) error {
	if databaseInstance.ExistingUserCheck(user.Email) {
		return errors.New("this email already exists")
	}

	_, err := databaseInstance.usersCollection.InsertOne(
		databaseInstance.ctx, user)

	return err
}

func (databaseInstance *MongoDatabase) GetUserByEmail(email string) User {
	filter := bson.D{{Key: "email", Value: email}}

	var document User

	err := databaseInstance.usersCollection.FindOne(
		databaseInstance.ctx, filter).Decode(&document)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}
		}
		return document
	}
	return document
}

/*
func (databaseInstance *MongoDatabase) GetAllUsers() []User {
	return []User{}
}
*/

func (databaseInstance *MongoDatabase) UpdateUserData(user User, query bson.D) error {
	filter := bson.D{
		{Key: "email", Value: user.Email},
	}
	opts := options.Update().SetUpsert(true)

	_, err := databaseInstance.usersCollection.UpdateOne(
		databaseInstance.ctx, filter, query, opts)
	if err != nil {
		return err
	}
	return err
}

func (databaseInstance *MongoDatabase) ExistingUserCheck(email string) bool {
	data := databaseInstance.GetUserByEmail(email)
	if data == (User{}) {
		return false
	}
	return true
}
