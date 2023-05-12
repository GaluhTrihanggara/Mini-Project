package routes

import (
	"MiniProject/constants"
	"MiniProject/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	jwtMiddleware := middleware.JWT([]byte(constants.SECRET_JWT))
	// Route /users to handler functions for user authentication and authorization
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	// Route /users/:id to handler functions for updating and deleting a user
	userGroup := e.Group("/users/:id", jwtMiddleware)
	userGroup.GET("", controllers.GetUserController)
	userGroup.PUT("", controllers.UpdateUserController)
	userGroup.DELETE("", controllers.DeleteUserController)
	userGroup.GET("/bills/:id", controllers.GetBillController, jwtMiddleware)
	userGroup.GET("/payments/:id", controllers.GetPaymentController, jwtMiddleware)
	userGroup.GET("/histories/:id", controllers.GetHistoriesByPaymentController)

	// Route /users to handler functions for user authentication and authorization
	e.POST("/admin/login", controllers.LoginAdminController)
	// group routes for admin
	adminGroup := e.Group("/admin", jwtMiddleware)
	adminGroup.GET("", controllers.GetAdminsController)
	adminGroup.GET("/:id", controllers.GetAdminController)
	adminGroup.POST("", controllers.CreateAdminController)
	adminGroup.PUT("/:id", controllers.UpdateAdminController)
	adminGroup.DELETE("/:id", controllers.DeleteAdminController)
	
	adminGroup.GET("/users", controllers.GetUsersController)
	adminGroup.GET("/users/:id", controllers.GetUserController)

	// Route /bills to handler function
	adminGroup.POST("/bills", controllers.CreateBillController)
	adminGroup.GET("/bills", controllers.GetBillsController)
	adminGroup.GET("/bills/:id", controllers.GetBillController)
	adminGroup.PUT("/bills/:id", controllers.UpdateBillController)
	adminGroup.DELETE("/bills/:id", controllers.DeleteBillController)

	// Route /payments to handler function
	adminGroup.POST("/payments", controllers.CreatePaymentController)
	adminGroup.GET("/payments", controllers.GetPaymentsController)
	adminGroup.GET("/payments/:id", controllers.GetPaymentController)
	adminGroup.PUT("/payments/:id", controllers.UpdatePaymentController)
	adminGroup.DELETE("/payments/:id", controllers.DeletePaymentController)

	// Group routes for histories accessible only by admin
	adminGroup.GET("/histories", controllers.GetHistoriesController)
	adminGroup.GET("/history/payment/:id", controllers.GetHistoriesByPaymentController)
	adminGroup.POST("/history", controllers.CreateHistoryController)
	adminGroup.PUT("/history/:id", controllers.UpdateHistoryController)
	adminGroup.DELETE("/history/:id", controllers.DeleteHistoryController)
	return e
}
