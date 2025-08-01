// Node: 20
// Path: /api/services/beneficiary.go
// Purpose: Manages designated system beneficiaries and allocations

package services

import (
	"encoding/json"
	"os"
)

// Beneficiary represents a system beneficiary.
// @Description Beneficiary represents a system beneficiary.
type Beneficiary struct {
	Name       string `json:"name"`
	Endpoint   string `json:"endpoint"`
	Allocation int    `json:"allocation"`
}

// BeneficiaryService is a service for managing system beneficiaries.
// @Description BeneficiaryService is a service for managing system beneficiaries.
type BeneficiaryService struct {
	beneficiaries []Beneficiary
}

// NewBeneficiaryService creates a new BeneficiaryService.
// @Description NewBeneficiaryService creates a new BeneficiaryService.
func NewBeneficiaryService(configFile string) (*BeneficiaryService, error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var beneficiaries []Beneficiary
	if err := json.Unmarshal(file, &beneficiaries); err != nil {
		return nil, err
	}

	return &BeneficiaryService{beneficiaries: beneficiaries}, nil
}

// GetBeneficiaries returns the list of beneficiaries.
// @Description GetBeneficiaries returns the list of beneficiaries.
func (s *BeneficiaryService) GetBeneficiaries() []Beneficiary {
	return s.beneficiaries
}