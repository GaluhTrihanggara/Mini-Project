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
	// Route / to handler function
	e.GET("/users", controllers.GetUsersController, jwtMiddleware)
	e.GET("/users/:id", controllers.GetUserController, jwtMiddleware)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController, jwtMiddleware)
	e.PUT("/users/:id", controllers.UpdateUserController, jwtMiddleware)

	// Route for admin login using JWT

	e.GET("/admin/:id", controllers.GetAdminController, jwtMiddleware)
	e.GET("/admin", controllers.GetAdminsController, jwtMiddleware)
	e.POST("/admin/login", controllers.LoginAdminController)
	e.POST("/admin", controllers.CreateAdminController)
	e.PUT("/admin/:id", controllers.UpdateAdminController, jwtMiddleware)
	e.DELETE("/admin/:id", controllers.DeleteAdminController, jwtMiddleware)

	// Route /payments to handler function
	e.POST("/payments", controllers.CreatePaymentController)
	e.GET("/payments", controllers.GetPaymentsController)
	e.GET("/payment/:id", controllers.GetPaymentController)
	e.PUT("/payments/:id", controllers.UpdatePaymentController)
	e.DELETE("/payments/:id", controllers.DeletePaymentController)

	// History endpoints
	e.POST("/history", controllers.CreateHistoryController)
	e.GET("/history", controllers.GetHistorysController)
	e.GET("/history/:id", controllers.GetHistoryController)
	e.GET("/history/user/:id", controllers.GetHistoryByUserController)
	e.GET("/history/payment/:id", controllers.GetHistoryByPaymentController)
	e.PUT("/history/:id", controllers.UpdateHistoryController)
	e.DELETE("/history/:id", controllers.DeleteHistoryController)
	return e
}
