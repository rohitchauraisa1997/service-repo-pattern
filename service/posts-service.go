package service

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/rohitchauraisa1997/service-repo-pattern/entity"
	repository "github.com/rohitchauraisa1997/service-repo-pattern/repositoryy"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id string) (*entity.Post, error)
}

type service struct{}

var (
	// repo repository.PostRepository = repository.NewPostFirestoreRepository()
	// rather than using the repo in above manner
	// we use this repo as a parameter to the contoller constructor function.
	repo repository.PostRepository
)

// To create instance of service
// constructor function
func NewPostService(rpstry repository.PostRepository) PostService {
	repo = rpstry
	return &service{}
}

func (serv *service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post's title is empty")
		return err
	}
	return nil
}

func (serv *service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (serv *service) FindAll() ([]entity.Post, error) {
	fmt.Println("Service FindAll triggered")
	return repo.FindAll()
}

func (serv *service) FindByID(id string) (*entity.Post, error) {
	return repo.FindByID(id)
}
