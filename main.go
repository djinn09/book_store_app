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
	r.GET("/books", controllers.FindBooks) // new
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.POST("/auth/login", controllers.LoginUser)

	r.Run()
}
