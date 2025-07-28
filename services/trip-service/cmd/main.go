package main

import (
	"context"
	"fmt"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main()  {
	ctx := context.Background()
	fmt.Println("starting the service..")
	//inmemRe
	inmemRepository := repository.NewInmemRepository()
	svc := service.NewService(inmemRepository)
	fare := &domain.RideFareModel{
		UserId: "42",
	}
	t,err := svc.CreateTrip(ctx,fare)
	
	if err != nil {
		log.Println(err)
	}
	log.Println(t)
	//temporary keep the program running for now


	for{
		time.Sleep(time.Second)
	}


}