package PoolOfChecks

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"log"
	"strconv"
)

//Income Check, Eligibility to check if the customers income is sufficient enough

type IncomeCheck struct {
	income int
}

func (I *IncomeCheck) Init(mp map[string]string) bool {
	inc, err := strconv.Atoi(mp["Income"])
	if err != nil {
		log.Println("Error Initializing IncomeCheck")
		log.Println(err)
		return false
	}
	I.income = inc
	log.Println("Parameters for IncomeCheck Initialized")
	return true
}

func (I *IncomeCheck) IsEligible(d *models.DecisionData) {
	declaredIncome := d.Record.Income
	if declaredIncome <= I.income {
		d.RejectionReason = append(d.RejectionReason, "IncomeCheck")
		d.Eligible = false
	}
}
