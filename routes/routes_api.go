package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/parinyapt/StreamySnap_AuthService/controller/handler"
	APIroutes "github.com/parinyapt/StreamySnap_AuthService/routes/api"
	// ApiRoutes "github.com/parinyapt/PT-Friendship_Backend/routes/api"
)

func configApiRoutes(router *gin.Engine) {
	// No Route 404 Notfound
	router.NoRoute(ctrlHandler.NoRouteHandler)

	// All Route
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//Public API
			APIroutes.InitClientAPI(v1)
			APIroutes.InitClientServiceAPI(v1)
			APIroutes.InitAuthPageAPI(v1)
			APIroutes.InitAccountAPI(v1)
			// ApiRoutes.SetupTestEndpoint(v1)
			// ApiRoutes.SetupHealthEndpoint(v1)
			// ApiRoutes.SetupProfileEndpoint(v1)

			//Private API with JWT Auth
			
		}
	}
}