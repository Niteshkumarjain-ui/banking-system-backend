package application

import (
	"banking-system-backend/domain"
	"banking-system-backend/util"
	"context"
)

func HealthGet(ctx context.Context) (response domain.HealthGetResponse, err error) {

	_, span := util.Tracer.Start(ctx, "HealthGet")
	defer span.End()
	response.Status = "I am up!"
	return
}
