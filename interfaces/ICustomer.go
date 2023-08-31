package interfaces

import "github.com/SinekaV/grpc-dal/models"

type ICustomer interface{
	CreateCustomer(Info *models.CustomerRequest)(*models.CustomerResponse,error)
}