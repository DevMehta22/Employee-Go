package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DevMehta22/mongoapi/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dotEnv(key string) string {
	
	err := godotenv.Load(".env")

	if err !=nil{
		panic(err)
	}

	return os.Getenv(key)
}

var collection *mongo.Collection

func init()  {
	connectionString := dotEnv("MONGO_URI")
	dbName := dotEnv("DB_NAME")
	collectionName := dotEnv("COLLECTION_NAME")

	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
		}
	
	fmt.Println("Connected to DB!")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection Instance is ready")

}

//Helper - file

func addEmp(emp model.Employee)  {
	inserted,err := collection.InsertOne(context.Background(),emp)
	if err != nil {
		panic(err)
		}
	fmt.Println("Inserted",inserted.InsertedID)
}

func updateEmp(empID string,emp model.Employee)  {
	id,_ := primitive.ObjectIDFromHex(empID)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":emp}

	result,err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
		}
	fmt.Println("Updated",result.ModifiedCount)	
}

func deleteEmp(empID string){
	id,_ := primitive.ObjectIDFromHex(empID)
	filter := bson.M{"_id":id}
	result,err := collection.DeleteOne(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
		}
	fmt.Println("Deleted",result.DeletedCount)
}

func deleteAll() int64 {
	result,err := collection.DeleteMany(context.Background(),bson.D{})
	if err != nil {
		log.Fatal(err)
		}
	fmt.Println("Deleted all records",result.DeletedCount)
	return result.DeletedCount
}

func findAll()  []primitive.M{
	cursor,err := collection.Find(context.Background(),bson.D{{}})
	if err != nil {
		log.Fatal(err)
		}
	var results []primitive.M

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()){
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
			}
		results = append(results,result)
	}

	return results
}

func findByID(empID string) bson.M{
	id,_ := primitive.ObjectIDFromHex(empID)
	filter := bson.M{"_id":id}
	var result bson.M
	err := collection.FindOne(context.Background(),filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		}
		fmt.Println(result)
	return result
} 


//controllers

func GetAllEmps(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("content-Type","application/json")
	emps := findAll()
	json.NewEncoder(w).Encode(emps)
}

func GetEmpByID(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	// w.Header().Set("Allow-Control-Allow-Methods","GET")

	params := mux.Vars(r)
	id := params["id"]
	emp := findByID(id)
	json.NewEncoder(w).Encode(emp)
}

func AddEmp(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var emp model.Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	addEmp(emp)
	json.NewEncoder(w).Encode(emp)

}

func UpdateEmp(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)
	id := params["id"]
	var emp model.Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	updateEmp(id,emp)
	json.NewEncoder(w).Encode(emp)
}

func DeleteOne(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	id := params["id"]
	deleteEmp(id)
	json.NewEncoder(w).Encode("Deleted Successfully:"+id)
}

func DeleteAll(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	
	count:= deleteAll()
	json.NewEncoder(w).Encode(count)
}
