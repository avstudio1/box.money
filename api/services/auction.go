// Node: 2
// Path: /api/services/auction.go
// Purpose: Core auction logic: validation, timer, emit settlement

package services

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

// AuctionService is a service for managing auctions.
// @Description AuctionService is a service for managing auctions.
type AuctionService struct {
	natsConn *nats.Conn
}

// NewAuctionService creates a new AuctionService.
// @Description NewAuctionService creates a new AuctionService.
func NewAuctionService(natsConn *nats.Conn) *AuctionService {
	return &AuctionService{natsConn: natsConn}
}

// ConcludeAuction concludes an auction and emits a settlement event.
// @Description ConcludeAuction concludes an auction and emits a settlement event.
func (s *AuctionService) ConcludeAuction(grossProceeds int) {
	settlementEvent := SettlementEvent{GrossProceeds: grossProceeds}

	payload, err := json.Marshal(settlementEvent)
	if err != nil {
		log.Printf("Error marshalling settlement event: %v", err)
		return
	}

	if err := s.natsConn.Publish("system.events.settlement", payload); err != nil {
		log.Printf("Error publishing settlement event: %v", err)
	}
}