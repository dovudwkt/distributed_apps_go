package main

import (
	"app/grades"
	"app/registry"
	"app/service"
	"context"
	"fmt"
	"log"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  serviceAddress,
	}

	ctx, err := service.Start(context.Background(), r, host, port, grades.RegisterHandlers)
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
