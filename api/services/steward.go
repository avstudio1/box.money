// Node: 19
// Path: /api/services/steward.go
// Purpose: Reinvestment calculator: listens for settlements, records reinvestment

package services

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

// StewardService is a service for handling settlements and reinvestments.
// @Description StewardService is a service for handling settlements and reinvestments.
type StewardService struct {
	beneficiaryService *BeneficiaryService
	natsConn           *nats.Conn
}

// NewStewardService creates a new StewardService.
// @Description NewStewardService creates a new StewardService.
func NewStewardService(beneficiaryService *BeneficiaryService, natsConn *nats.Conn) *StewardService {
	return &StewardService{
		beneficiaryService: beneficiaryService,
		natsConn:           natsConn,
	}
}

// SettlementEvent represents a settlement event.
// @Description SettlementEvent represents a settlement event.
type SettlementEvent struct {
	GrossProceeds int `json:"gross_proceeds"`
}

// Start starts the steward service.
// @Description Start starts the steward service.
func (s *StewardService) Start() {
	_, err := s.natsConn.Subscribe("system.events.settlement", func(msg *nats.Msg) {
		var event SettlementEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Error unmarshalling settlement event: %v", err)
			return
		}

		reinvestment := event.GrossProceeds / 10

		log.Printf("Reinvesting %d", reinvestment)
	})
	if err != nil {
		log.Fatalf("Error subscribing to settlement events: %v", err)
	}
}