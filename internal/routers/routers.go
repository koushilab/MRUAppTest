import (
	"github.com/koushilab/MRUAppTest/internal/handlers"
	"github.com/gin-gonic/gin"
)

//StartServer starts the server with the provided list of routes from  SetupRoutes() method
func StartServer() {

	runServer := SetupRoutes()
	log.Info("Calling Setup Routes")
	runServer.Run("localhost:8082")

}

func SetupRoutes() *gin.Engine {
	ci := handlers.MRUInfo{}
	router := gin.Default()

	v1 := router.Group("/api/v1")

	{

		v1.GET("/getMRUs", ci.GetMRUs)
	}
	return router
}