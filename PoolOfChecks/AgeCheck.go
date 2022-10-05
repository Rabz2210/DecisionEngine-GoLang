package PoolOfChecks

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"log"
	"strconv"
)

//Age Check, Eligibility to check if the customer is Old enough to be eligibile

type AgeCheck struct {
	age int
}

func (A *AgeCheck) Init(mp map[string]string) bool {
	age, err := strconv.Atoi(mp["Age"])
	if err != nil {
		log.Println("Error Initializing AgeCheck")
		return false
	}
	A.age = age
	log.Printf("Parameter for AgeCheck Initialized")
	return true
}

func (A *AgeCheck) IsEligible(d *models.DecisionData) {
	if d.Record.Age < A.age {
		d.RejectionReason = append(d.RejectionReason, "AgeCheck")
		d.Eligible = false
	}
}
