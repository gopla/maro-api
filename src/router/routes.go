package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gopla/maro/src/controller"
	"github.com/gopla/maro/src/middleware"
	"github.com/gopla/maro/src/service"
)

func SetupRouter() *gin.Engine {

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var authController controller.AuthController = controller.NewAuthController(loginService, jwtService)


	r := gin.Default()

	questionRoute := r.Group("/question")
	{
		questionRoute.GET("/",controller.IndexQuestion)
		questionRoute.POST("/",controller.StoreQuestion)
		questionRoute.GET("/:id",controller.ShowQuestion)
		questionRoute.GET("/user/:username",controller.ShowQuestionByUsername)
		questionRoute.PUT("/:id",controller.UpdateQuestion)
		questionRoute.DELETE("/:id",controller.DeleteQuestion)
	}

	userRoute := r.Group("/user")
	{
		userRoute.GET("/",controller.IndexUser)
		userRoute.POST("/",controller.StoreUser)
		userRoute.GET("/:id",controller.ShowUser)
		userRoute.PUT("/:id",controller.UpdateUser)
		userRoute.DELETE("/:id",controller.DeleteUser)
	}

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}

	userQuestionRoute := r.Group("/qa", middleware.AuthorizeJWT(jwtService))
	{
		userQuestionRoute.GET("/", controller.ShowQuestionPerUser)
		userQuestionRoute.GET("/question/:id",controller.ShowQuestion)
		userQuestionRoute.PUT("/answer/:id", controller.AnswerQuestion)
	}

	profileRoute := r.Group("/profile", middleware.AuthorizeJWT(jwtService))
	{
		profileRoute.GET("/",controller.GetProfile)
		profileRoute.PUT("/updateStatus",controller.UpdateStatus)
	}

	return r
}