package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sanprasirt/grpc-go-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	// c := greetpb.NewGreetServiceClient(conn)
	c := calculatorpb.NewCalculatorServiceClient(conn)
	// fmt.Printf("Created client: %f", c)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary Rpc")
	req := &calculatorpb.SumRequest{
		FisrtNumber:  40,
		SecondNumber: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}
