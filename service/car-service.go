package service

import (
	"fmt"
	"math/rand"

	"github.com/rohitchauraisa1997/service-repo-pattern/entity"
	repository "github.com/rohitchauraisa1997/service-repo-pattern/repositoryy"
)

type carService struct{}

type CarDetailsService interface {
	FindAllCarDetailsService() ([]entity.CarDetails, error)
	AddCarService(carDetails *entity.CarDetails) (*entity.CarDetails, error)
	FindByIdCarDetailsService(id string) (*entity.CarDetails, error)
}

var (
	// we use this repo as a parameter to the contoller constructor function.
	carRepo repository.CarRepository
)

// To create instance of service
// constructor function
func NewCarDetailsService(rpstry repository.CarRepository) CarDetailsService {
	carRepo = rpstry
	return &carService{}
}

func (serv *carService) FindAllCarDetailsService() ([]entity.CarDetails, error) {
	fmt.Println("FindAllCarDetailsService triggered")
	return carRepo.FindAllCars()
}

func (serv *carService) AddCarService(carDetails *entity.CarDetails) (*entity.CarDetails, error) {
	carDetails.ID = rand.Int63()
	return carRepo.AddCar(carDetails)
}

func (serv *carService) FindByIdCarDetailsService(id string) (*entity.CarDetails, error) {
	return carRepo.FindByCarId(id)
}
