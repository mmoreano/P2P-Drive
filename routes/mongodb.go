package routes

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"zendx.io/P2P-Drive/models"
)

// Model for storting database client and is used for closing connection.
type MongoDb struct {
	Client *mongo.Client
}

// -------------------------- Establish DB Connection/Client --------------------------\\

func Connection() *MongoDb {
	url := "mongodb+srv://admin:DOM123@domsdb.agpuaxn.mongodb.net/?retryWrites=true&w=majority"
	//url := "mongodb+srv://" + os.Getenv("USER") + os.Getenv("PASS") + "@domsdb.agpuaxn.mongodb.net/?retryWrites=true&w=majority"
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(url).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return &MongoDb{Client: client}
}

//-------------------------- Register User into DB with client --------------------------\\

func (connection *MongoDb) DBregister(userInfo *models.RegisterRequest) {
	db := connection.Client.Database("P2P")
	coll := db.Collection("Users")
	userInfo.Token = uuid.New().String()
	userInfo.UserPassword = string(encrypt([]byte(userInfo.Username+userInfo.UserPassword), userInfo.Token[:32]))

	docs := bson.M{"_id": userInfo.Email, "Username": userInfo.Username, "Password": userInfo.UserPassword, "Number": userInfo.Number, "Email": userInfo.Email,
		"Fname": userInfo.FirstName, "Lname": userInfo.LastName, "Token": userInfo.Token}
	result, err := coll.InsertOne(context.TODO(), docs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Inserted document with ID %v\n", result.InsertedID)
}
func (connection *MongoDb) DBemailCheck(email string) string {
	var info models.RegisterRequest

	db := connection.Client.Database("P2P") //Set Database
	coll := db.Collection("Users")          //Set Collection
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println("Retreiving information...")
	filter := bson.M{"_id": email} //Set Filter
	//filter := bson.M{"Email": email}

	i, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return "Found"
	}
	for i.Next(context.TODO()) {
		var result models.RegisterRequest
		if err := i.Decode(&result); err != nil {
			panic(err)
		}
		info.Email = result.Email
	}
	fmt.Println("Successfully Retrieved")
	print(info.Email)
	if info.Email == "" {
		return "Not Found"
	} else {
		return "Found"
	}
}

//-------------------------- Register User into DB with client --------------------------\\
//
//func (connection *MongoDb) DBpassCheck(password string) string {
//
//	db := connection.Client.Database("P2P") //Set Database
//	coll := db.Collection("Users")          //Set Collection
//	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	fmt.Println("Retreiving information...")
//	filter := bson.M{"Password": password} //Set Filter
//	//filter := bson.M{"Email": email}
//
//	i, err := coll.Find(context.TODO(), filter)
//	if err != nil {
//		return "Found"
//	}
//	for i.Next(context.TODO()) {
//		var result models.RegisterRequest
//		if err := i.Decode(&result); err != nil {
//			panic(err)
//		}
//		info.Email = result.Email
//	}
//	fmt.Println("Successfully Retrieved")
//	print(info.Email)
//	if info.Email == "" {
//		return "Not Found"
//	} else {
//		return "Found"
//	}
//}

// -------------------------- Get User Token from DB with client --------------------------\\

func (connection *MongoDb) Login(user *models.LoginRequest) string {
	db := connection.Client.Database("P2P")
	coll := db.Collection("Users")
	var result models.RegisterRequest

	fmt.Println("Retreiving information...")
	filter := bson.M{"Username": user.Username}

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}

	pass := string(encrypt([]byte(user.Username+user.UserPassword), result.Token[:32]))

	fmt.Println("1")
	fmt.Println(pass)
	fmt.Println("2")
	//fmt.Println(user.UserPassword)
	fmt.Println(result)
	fmt.Println("Successfully Retrieved")

	if user.UserPassword == pass {
		return result.Token
	} else {
		return "Incorrect Password"
	}
}

//-------------------------- Get User file info from DB with client --------------------------\\

func (connection *MongoDb) GetUserFiles(owner string) []models.AddResponse {
	var files []models.AddResponse
	db := connection.Client.Database("P2P")
	coll := db.Collection("User_Files")

	fmt.Println("Retreiving information...")

	filter := bson.M{"Owner": owner}
	fmt.Println("Retreiving information...")
	i, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return []models.AddResponse{}
	}

	fmt.Println("Retreiving information...")

	for i.Next(context.TODO()) {
		var result models.AddResponse

		if err := i.Decode(&result); err != nil {
			panic(err)
		}
		files = append(files, result)

	}
	fmt.Println("Successfully Retrieved")
	return files
}

//-------------------------- Upload File Data to DB with client --------------------------\\

func (connection *MongoDb) DBupload(file models.AddResponse) {

	db := connection.Client.Database("P2P")
	coll := db.Collection("User_Files")
	docs := bson.M{"Hash": file.Hash, "Name": file.Name,
		"Size": file.Size, "Link": file.Link, "Owner": file.Owner}
	result, err := coll.InsertOne(context.TODO(), docs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("inserted document with ID %v\n", result.InsertedID)
}

//-------------------------- Close Client --------------------------\\

func (client *MongoDb) CloseClientDB() {
	if client == nil {
		return
	}
	err := client.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// TODO optional you can log your closed MongoDB client
	fmt.Println("Connection to MongoDB closed.")
}
