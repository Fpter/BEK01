package main

import (
	calculator "awesomeProject/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer cc.Close()
	client := calculator.NewCalculatorServiceClient(cc)

	log.Printf("service client %f", client)

}
