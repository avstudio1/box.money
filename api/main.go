// Node: 1
// Path: /api/main.go
// Purpose: Entrypoint: Init all services, including steward/beneficiary

package main

import (
	"log"
	"net"

	"github.com/box-money/goldstream/api/gen/auction"
	grpc_server "github.com/box-money/goldstream/api/grpc"
	"github.com/box-money/goldstream/api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	// NATS
	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer natsConn.Close()

	// Services
	beneficiaryService, err := services.NewBeneficiaryService("configs/beneficiaries.json")
	if err != nil {
		log.Fatalf("Failed to create beneficiary service: %v", err)
	}

	stewardService := services.NewStewardService(beneficiaryService, natsConn)
	stewardService.Start()

	auctionService := services.NewAuctionService(natsConn)

	// gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auction.RegisterAuctionServiceServer(s, grpc_server.NewAuctionServer(auctionService))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
