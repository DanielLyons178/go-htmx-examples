package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daniellyons178/htmx/models"
	"github.com/daniellyons178/htmx/repositories"
	"github.com/gin-gonic/gin"
)

type InfiniteScollRoutes struct {
	repo repositories.CustomerRepo
}

func (r InfiniteScollRoutes) ViewCustomers(ctx *gin.Context) {
	pageParam := ctx.Query("page")
	pageSize := 20
	page, err := strconv.Atoi(pageParam)
	template := "infiniteScroll/tableContent"

	if err != nil {
		template = "infiniteScroll/index"
		page = 0
	}
	customers := r.repo.GetCustomers(page, pageSize)
	ctx.HTML(
		http.StatusOK,
		template,
		gin.H{"PageSize": pageSize, "NextPage": page + 1, "Customers": makeCustomersList(customers, pageSize*page)})
}

func makeCustomersList(c []models.Customer, start int) []gin.H {
	res := make([]gin.H, len(c))

	for i, v := range c {
		res[i] = gin.H{"Name": fmt.Sprintf("%v %v", v.Name, v.Surname), "Email": v.Email, "Id": start + i, "Number": i + 1}
	}
	return res
}

func CreateInfiniteScrollRoutes(r repositories.CustomerRepo) InfiniteScollRoutes {
	return InfiniteScollRoutes{repo: r}
}
