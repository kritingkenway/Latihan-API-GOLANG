package routes

import (
	"coba/handler"
	"coba/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App)  {
	app.Get("/", handler.Hello)

	// app.Post("/user", handler.HelloName)

	//bikin API user
	app.Get("/user",middleware.JWTProtectedRoute(), handler.GetUsers)

	app.Get("/user/:id",middleware.JWTProtectedRoute(), handler.GetUserById)

	app.Post("/user", handler.CreateUser)
	
	app.Put("/user/:id",middleware.JWTProtectedRoute(), handler.UpdateUser)
	
	app.Post("/login", handler.Login)
	
	app.Post("/logout", middleware.JWTProtectedRoute(), handler.Logout)



	// Route Product
	app.Get("/product", handler.GetProducts)
	app.Get("/product/:id", handler.GetProductById)
	app.Post("/product", handler.CreateProduct)
	app.Put("/product/:id", handler.UpdateProduct)
	app.Delete("/product/:id", handler.DeleteProduct)


	//Route CART
	app.Get("/cart", middleware.JWTProtectedRoute(),middleware.Cart(), handler.GetCartItems)
	app.Get("/cart/add/:id", middleware.JWTProtectedRoute(),middleware.Cart(), handler.AddItemCart)
	app.Get("/cart/subtract/:id", middleware.JWTProtectedRoute(),middleware.Cart(), handler.SubtractItemCart)

	//melakukan order
	app.Get("/myorder", middleware.JWTProtectedRoute(),middleware.Cart(), handler.GetMyOrders)
	app.Post("/makeorder", middleware.JWTProtectedRoute(),middleware.Cart(), handler.MakeOrders)

	
}