package routes

import (
	"github.com/HiteshKumarMeghwar/L-M-S-2/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Login and registration routes ................
	app.Post("/api/register", controllers.Register)
	// app.Post("/api/login", controllers.Login)

	// Middleware for check user is Authenticated or not ...........
	// app.Use(middleware.IsAuthenticate)

	// // Post Controlxler routes ............
	// app.Post("/api/createpost", controllers.CreatePost)
	// app.Post("/api/allpost", controllers.AllPost)
	// app.Get("/api/allpost/:id", controllers.DetailPost)
	// app.Put("/api/updatepost/:id", controllers.UpdatePost)
	// app.Post("/api/uniquepost", controllers.UniquePost)
	// app.Delete("/api/deletepost/:id", controllers.DeletePost)
	// // Send Email by user .............
	// app.Post("/api/sendMail/:id", controllers.SendMail)

	// // Image upload route for practice ..................
	// app.Post("/api/upload-image", controllers.Upload)
	// app.Static("/api/upload", "./uploads")

	// // User and logout routes ..............
	// app.Get("/api/user", controllers.User)
	// app.Post("/api/allUsers", controllers.AllUser)
	// app.Post("/api/allUsers/:id", controllers.SingleUser)
	// app.Put("/api/updateProfile/:id", controllers.UpdateProfile)
	// app.Put("/api/updateUser/:id", controllers.UpdateUser)
	// app.Delete("/api/deleteUser/:id", controllers.DeleteUser)
	// app.Post("/api/logout", controllers.Logout)

	// // Commodity Module routes ..............
	// app.Get("/api/getAllCommodities", controllers.GetCommodities)
	// app.Post("/api/addCommodity", controllers.AddCommodity)
	// app.Post("/api/getAllCommodities/:id", controllers.GetCommodityById)
	// app.Put("/api/updateCommodity/:id", controllers.UpdateCommodityById)
	// app.Delete("/api/deleteCommodity/:id", controllers.DeleteCommodityById)
	// // Weather API route ....
	// app.Post("/api/weatherData/:location", controllers.WeatherAPI)

	// // Test Cases routes ..............
	// // app.Get("/api/testCase", controllers.TestAPI)
}
