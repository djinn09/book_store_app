package main

import (
	"crud/controllers"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		var HtmlString = `<html>
			<body>
				<h1>Hello Hell !!!</h1>
			</body>
		</html>
		`

		//Write your 200 header status (or other status codes, but only WriteHeader once)
		c.Writer.WriteHeader(200)
		//Convert your cached html string to byte array
		c.Writer.Write([]byte(HtmlString))
	})
	r.POST("/auth/login", controllers.LoginUser)
	authorized := r.Group("/")
	authorized.Use(controllers.JWTAuthMiddleware())
	{
		authorized.GET("books", controllers.FindBooks) // new
		authorized.POST("/books", controllers.CreateBook)
		authorized.GET("/books/:id", controllers.FindBook)
		authorized.PATCH("/books/:id", controllers.UpdateBook)
		authorized.DELETE("/books/:id", controllers.DeleteBook)
		authorized.GET("/users", controllers.FindUsers)
		authorized.POST("/users", controllers.CreateUser)
	}

	r.Run()
}
