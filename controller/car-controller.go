package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohitchauraisa1997/service-repo-pattern/entity"
	"github.com/rohitchauraisa1997/service-repo-pattern/errors"
	"github.com/rohitchauraisa1997/service-repo-pattern/service"
)

var (
	carService service.CarDetailsService
)

type CarController interface {
	GetCars(response http.ResponseWriter, request *http.Request)
	AddCar(response http.ResponseWriter, request *http.Request)
	GetCarById(response http.ResponseWriter, request *http.Request)
}

type carController struct{}

// constructor function
func GetNewCarController(service service.CarDetailsService) CarController {
	carService = service
	return &carController{}
}

func (*carController) GetCars(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "Application/json")
	cars, err := carService.FindAllCarDetailsService()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error getting the cars!!"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(cars)
}

func (*carController) AddCar(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var carDetails entity.CarDetails
	err := json.NewDecoder(request.Body).Decode(&carDetails)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "error decoding json object"})
		return
	}
	result, err := carService.AddCarService(&carDetails)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errorMessage := fmt.Sprintf("error saving the post %v", err.Error())
		json.NewEncoder(response).Encode(errors.ServiceError{Message: errorMessage})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*carController) GetCarById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	docId := mux.Vars(request)["id"]
	car, _ := carService.FindByIdCarDetailsService(docId)
	if car == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "car not found with correspondng id"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(car)
}
