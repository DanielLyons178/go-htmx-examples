package main

import (
	"github.com/daniellyons178/htmx/repositories"
	"github.com/daniellyons178/htmx/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := repositories.CreateFakeRepo()

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	clickToEdit := r.Group("clicktoedit")
	{
		routes := routes.CreateClickToEditRoutes(repo)
		clickToEdit.GET("/contact/:id", routes.ViewCustomer)
		clickToEdit.GET("/contact/:id/edit", routes.CustomerEditForm)
		clickToEdit.PUT("/contact/:id", routes.EditCustomer)
	}
	infiniteScroll := r.Group("infinite-scroll")
	{
		routes := routes.CreateInfiniteScrollRoutes(repo)
		infiniteScroll.GET("/contacts", routes.ViewCustomers)
	}
	r.Run(":8080")
}
