package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/HMadhav/CRM/configs"
	"github.com/HMadhav/CRM/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var customerCollection *mongo.Collection = configs.GetCollection(configs.DB, "customers")
var validate = validator.New()

// Get All Customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := customerCollection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := models.CustomerResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	var customers []models.Customer
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &customers); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := models.CustomerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := models.CustomerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": customers}}
	json.NewEncoder(w).Encode(response)
}

// Get Single Customers
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	params := mux.Vars(r)
	id := params["id"]
	objId, _ := primitive.ObjectIDFromHex(id)

	var customer models.Customer
	err := customerCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := models.CustomerResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := models.CustomerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": customer}}
	json.NewEncoder(w).Encode(response)
}

// Add Customers
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var customer models.Customer

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.CustomerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	if validationErr := validate.Struct(&customer); validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.CustomerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	newCustomer := models.Customer{
		Id:        primitive.NewObjectID(),
		Name:      customer.Name,
		Role:      customer.Role,
		Email:     customer.Email,
		Phone:     customer.Phone,
		Contacted: customer.Contacted,
	}

	result, err := customerCollection.InsertOne(ctx, newCustomer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := models.CustomerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := models.CustomerResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(w).Encode(response)

}

// Edit Customers
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	params := mux.Vars(r)
	id := params["id"]
	objId, _ := primitive.ObjectIDFromHex(id)

	var customer map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.CustomerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	var updatedCustomer models.Customer
	err := customerCollection.FindOneAndUpdate(ctx, bson.M{"id": objId}, bson.M{"$set": customer}).Decode(&updatedCustomer)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := models.CustomerResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	var updatedDocument models.Customer
	err = customerCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedDocument)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := models.CustomerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := models.CustomerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedDocument}}
	json.NewEncoder(w).Encode(response)

}

// Delete Customers
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	params := mux.Vars(r)
	id := params["id"]
	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := customerCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := models.CustomerResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	if result.DeletedCount < 1 {
		w.WriteHeader(http.StatusNotFound)
		response := models.CustomerResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := models.CustomerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Customer successfully deleted!"}}
	json.NewEncoder(w).Encode(response)

}
