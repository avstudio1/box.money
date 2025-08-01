// Node: 3
// Path: /api/grpc/auction_server.go
// Purpose: gRPC stream: bid submit + push updates

package grpc

import (
	"io"

	"github.com/box-money/goldstream/api/gen/auction"
	"github.com/box-money/goldstream/api/services"
)

// AuctionServer is the gRPC server for the auction service.
// @Description AuctionServer is the gRPC server for the auction service.
type AuctionServer struct {
	auction.UnimplementedAuctionServiceServer
	auctionService *services.AuctionService
}

// NewAuctionServer creates a new AuctionServer.
// @Description NewAuctionServer creates a new AuctionServer.
func NewAuctionServer(auctionService *services.AuctionService) *AuctionServer {
	return &AuctionServer{auctionService: auctionService}
}

// BidStream is the bidirectional stream for bids and updates.
// @Description BidStream is the bidirectional stream for bids and updates.
func (s *AuctionServer) BidStream(stream auction.AuctionService_BidStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		s.auctionService.ConcludeAuction(int(req.Amount))
	}
}