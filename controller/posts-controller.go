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
	// postService = service.NewPostService()
	// rather than using the postService in above manner
	// we use this service as a parameter to the contoller constructor function.
	postService service.PostService
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPosts(response http.ResponseWriter, request *http.Request)
	GetPostByDocumentID(response http.ResponseWriter, request *http.Request)
}

type controller struct{}

// constructor function
func GetNewController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error getting the posts!!"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) AddPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "error decoding json object"})
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errorMessage := fmt.Sprintf("error while validating request %v", err.Error())
		json.NewEncoder(response).Encode(errors.ServiceError{Message: errorMessage})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errorMessage := fmt.Sprintf("error saving the post %v", err.Error())
		json.NewEncoder(response).Encode(errors.ServiceError{Message: errorMessage})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*controller) GetPostByDocumentID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	docId := mux.Vars(request)["id"]
	post, _ := postService.FindByID(docId)
	if post == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "post not found with correspondng id"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
