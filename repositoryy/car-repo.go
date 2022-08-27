package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/rohitchauraisa1997/service-repo-pattern/entity"
	"google.golang.org/api/iterator"
)

type carRepo struct{}

type CarRepository interface {
	FindAllCars() ([]entity.CarDetails, error)
	AddCar(carDetails *entity.CarDetails) (*entity.CarDetails, error)
	FindByCarId(id string) (*entity.CarDetails, error)
}

// To create instance
// constructor function
func NewCarFirestorRepository() CarRepository {
	return &carRepo{}
}

const (
	carProjectId      string = "service-repo"
	carCollectionName string = "cars"
)

func (*carRepo) FindAllCars() ([]entity.CarDetails, error) {
	fmt.Println("FindAllCars triggered")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, carProjectId)
	if err != nil {
		log.Fatal("failed to create firestore client!!")
		return nil, err
	}

	// defer client.Close()
	var cars []entity.CarDetails
	iter := client.Collection(carCollectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal("failed to iterate the list of cars ", err)
			return nil, err
		}
		car := entity.CarDetails{
			ID:    doc.Data()["ID"].(int64),
			Brand: doc.Data()["Brand"].(string),
			Model: doc.Data()["Model"].(string),
			Year:  doc.Data()["Year"].(string),
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (*carRepo) AddCar(carDetails *entity.CarDetails) (*entity.CarDetails, error) {
	fmt.Println("AddCar triggered")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, carProjectId)
	if err != nil {
		log.Fatal("failed to create firestore client!!")
		return nil, err
	}

	fmt.Println(carDetails.Model)
	_, _, err = client.Collection(carCollectionName).Add(ctx, map[string]interface{}{
		"ID":    carDetails.ID,
		"Brand": carDetails.Brand,
		"Model": carDetails.Model,
		"Year":  carDetails.Year,
	})
	if err != nil {
		log.Fatal("failed adding a new car!! ", err)
		return nil, err
	}
	return carDetails, nil
}

func (*carRepo) FindByCarId(id string) (*entity.CarDetails, error) {
	fmt.Println("FindByCarId triggered")
	fmt.Println("Finding carid: ", id)

	ctx := context.Background()
	var car entity.CarDetails
	client, err := firestore.NewClient(ctx, carProjectId)
	if err != nil {
		log.Fatal("failed to create firestore client!!")
		return nil, err
	}
	defer client.Close()
	cars := client.Collection(carCollectionName)
	docref := cars.Doc(id)
	docsnap, err := docref.Get(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, err
		}
		log.Fatal("failed to findByID ", err)
		return nil, err
	}
	docsnap.DataTo(&car)

	return &car, nil
}

