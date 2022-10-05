package PoolOfChecks

//Phone Area Code check , Approve or Reject Customers based on the area code of their phone numbers

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"log"
	"strings"
)

type PhoneAreaCodeCheck struct {
	areaCode map[string]bool
}

func (p *PhoneAreaCodeCheck) Init(mp map[string]string) bool {
	ls := strings.Split(mp["AreaCodes"], ",")
	p.areaCode = make(map[string]bool, len(ls))
	for _, value := range ls {
		p.areaCode[value] = true
	}
	log.Println("Parameter for PhoneAreaCode Initialized")
	return true
}

func (p *PhoneAreaCodeCheck) IsEligible(d *models.DecisionData) {
	phNo := strings.TrimSpace(d.Record.PhoneNumber)
	if firstDigit := phNo[0:1]; !p.areaCode[firstDigit] {
		d.RejectionReason = append(d.RejectionReason, "PhoneAreaCodeCheck")
		d.Eligible = false
	}
}
