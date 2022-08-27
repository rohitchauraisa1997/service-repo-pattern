package main

import (
	"fmt"
	"net/http"

	"github.com/rohitchauraisa1997/service-repo-pattern/controller"
	router "github.com/rohitchauraisa1997/service-repo-pattern/http"
	repository "github.com/rohitchauraisa1997/service-repo-pattern/repositoryy"
	"github.com/rohitchauraisa1997/service-repo-pattern/service"
)

var (
	repo           repository.PostRepository = repository.NewPostFirestoreRepository()
	postService    service.PostService       = service.NewPostService(repo)
	postController controller.PostController = controller.GetNewController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
	// httpRouter                               = router.NewChiRouter()
)

var (
	carsRepo       repository.CarRepository  = repository.NewCarFirestorRepository()
	carsService    service.CarDetailsService = service.NewCarDetailsService(carsRepo)
	carsController controller.CarController  = controller.GetNewCarController(carsService)
	carHttpRouter  router.Router             = router.NewMuxRouter()
	// httpRouter                               = router.NewChiRouter()
)

func main() {
	const port string = ":9000"

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByDocumentID)

	carHttpRouter.GET("/cars", carsController.GetCars)
	carHttpRouter.GET("/cars/{id}", carsController.GetCarById)
	carHttpRouter.POST("/cars", carsController.AddCar)
	httpRouter.SERVE(port)

}
