package PoolOfChecks

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"github.com/honestbank/tech-assignment-backend-engineer/risk"
	"log"
	"strconv"
)

//Check to determine Eligibility Based on the No Of cards held by the customer along with his/her calculated CreditRisk

type CreditRiskCheck struct {
	noOfCards int
}

func (c *CreditRiskCheck) Init(mp map[string]string) bool {
	cards, err := strconv.Atoi(mp["NoOfCards"])
	if err != nil {
		log.Println("Error Initializing CardCheck Parameters")
		return false
	}
	c.noOfCards = cards
	log.Printf("Parameter for CreditRiskCheck Initialized")
	return true
}

func (c *CreditRiskCheck) IsEligible(d *models.DecisionData) {
	if *d.Record.NumberOfCreditCards > c.noOfCards {
		d.RejectionReason = append(d.RejectionReason, "CreditRiskCheck-NoOfCards")
		d.Eligible = false
		return
	}
	CreditRisk := "LOW" == risk.CalculateCreditRisk(d.Record.Age, *d.Record.NumberOfCreditCards)
	if !CreditRisk {
		d.RejectionReason = append(d.RejectionReason, "CreditRiskCheck-CalculatedRisk")
		d.Eligible = false
	}
}
