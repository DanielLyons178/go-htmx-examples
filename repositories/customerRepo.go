package repositories

import (
	"fmt"

	"github.com/daniellyons178/htmx/models"
)

type CustomerRepo interface {
	GetCustomer(id int) models.Customer
	SetCustomerDetails(id int, details models.Customer) models.Customer
	GetCustomers(page int, pageSize int) []models.Customer
}

type FakeRepo struct {
	customers []models.Customer
}

func (r FakeRepo) GetCustomer(id int) models.Customer {
	return r.customers[id]
}

func (r FakeRepo) SetCustomerDetails(id int, details models.Customer) models.Customer {
	r.customers[id] = details
	return details
}

func (r FakeRepo) GetCustomers(page int, pageSize int) []models.Customer {
	start := page * pageSize
	res := make([]models.Customer, pageSize)
	for i := 0; i < pageSize; i++ {
		res[i] = models.Customer{Name: fmt.Sprintf("FirstName %v", start+i), Surname: fmt.Sprintf("LastName %v", start+i), Email: fmt.Sprintf("a%v@abc.com", start+i)}
	}
	return res
}

func CreateFakeRepo() CustomerRepo {
	return &FakeRepo{
		customers: []models.Customer{
			{Name: "SOMEONE", Surname: "SRNAME", Email: "a@abc.com"},
		},
	}
}
