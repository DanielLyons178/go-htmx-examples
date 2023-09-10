package routes

import (
	"net/http"
	"strconv"

	"github.com/daniellyons178/htmx/models"
	"github.com/daniellyons178/htmx/repositories"
	"github.com/gin-gonic/gin"
)

type ClickToEditRoutes struct {
	repo repositories.CustomerRepo
}

func (r ClickToEditRoutes) ViewCustomer(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Id must be integer", nil)
		return
	}
	customer := r.repo.GetCustomer(id)

	ctx.HTML(
		http.StatusOK,
		"clicktoedit/customer",
		createViewModel(id, customer))
}

func createViewModel(id int, customer models.Customer) gin.H {
	x := gin.H{
		"id":       id,
		"customer": gin.H{"FirstName": customer.Name, "Surname": customer.Surname, "Email": customer.Email}}
	return x
}

func (r ClickToEditRoutes) CustomerEditForm(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Id must be integer", nil)
		return
	}
	customer := r.repo.GetCustomer(id)

	ctx.HTML(
		http.StatusOK,
		"clicktoedit/edit",
		createViewModel(id, customer))
}

type CustomerRequest struct {
	Fistname string `form:"firstName"`
	Surname  string `form:"lastName"`
	Email    string `form:"email"`
}

func (r ClickToEditRoutes) EditCustomer(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Id must be integer", nil)
		return
	}

	req := &CustomerRequest{}
	ctx.Bind(req)

	customer := r.repo.SetCustomerDetails(id, models.Customer{
		Name:    req.Fistname,
		Surname: req.Surname,
		Email:   req.Email,
	})

	ctx.HTML(
		http.StatusOK,
		"clicktoedit/customer",
		createViewModel(id, customer))

}

func CreateClickToEditRoutes(repo repositories.CustomerRepo) *ClickToEditRoutes {
	return &ClickToEditRoutes{
		repo: repo,
	}
}
