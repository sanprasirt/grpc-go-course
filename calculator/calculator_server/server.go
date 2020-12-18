package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sanprasirt/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

// func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
// 	fmt.Printf("Greet fn was invokeed with: %v\n", req)
// 	firstName := req.GetGreeting().GetFirstName()
// 	result := "Hello " + firstName
// 	res := &greetpb.GreetResponse{
// 		Result: result,
// 	}
// 	return res, nil
// }
func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Calculator sum was invoked with: %v\n", req)
	fisrtNumber := req.FisrtNumber
	secondNumber := req.SecondNumber
	sum := fisrtNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumber RPC: %v\n", req)

	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil

}
func main() {
	fmt.Println("Hello from Calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// greetpb.RegisterGreetServiceServer(s, &server{})
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
