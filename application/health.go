package application

import (
	"banking-system-backend/domain"
)

func HealthGet() (response domain.HealthGetResponse, err error) {

	response.Status = "I am up!"
	return
}
