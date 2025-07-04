package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/controller"
	"github.com/sb-67-go-quiz-salman-input-data-buku/middleware"
)

func listRouteUser(router *gin.Engine) {
	userRoute := router.Group("/api/users")
	{
		userRoute.POST("/login", controller.LoginUser)
		userRoute.POST("/register", controller.RegisterUser)
		userRoute.PUT("/update/:id", middleware.CheckJwt(), controller.UpdateUser)
	}
}

func listRouteCategory(router *gin.Engine) {
	categoryRoute := router.Group("/api/categories")
	categoryRoute.Use(middleware.CheckJwt())
	{
		categoryRoute.GET("/", controller.GetAllCategory)
		categoryRoute.GET("/:id", controller.GetCategoryById)
		categoryRoute.POST("/", controller.CreateCategory)
		categoryRoute.GET("/:id/books", controller.GetBookInCategory)
		categoryRoute.DELETE("/:id", controller.DeleteCategory)
	}
}

func listRouteBook(router *gin.Engine) {
	bookRoute := router.Group("/api/books")
	bookRoute.Use(middleware.CheckJwt())
	{
		bookRoute.GET("/", controller.GetAllBook)
		bookRoute.GET("/:id", controller.GetBookById)
		bookRoute.POST("/", controller.CreateBook)
		bookRoute.DELETE("/:id", controller.DeleteBook)
	}
}

func ListAllRouter(router *gin.Engine) {
	listRouteUser(router)
	listRouteCategory(router)
	listRouteBook(router)
}
