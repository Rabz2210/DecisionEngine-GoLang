package PoolOfChecks

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"log"
	"strconv"
)

//Check to Determine Eligibility Of a customer based On the his/her Political Exposure

type PoliticallyExposedCheck struct {
	poliExpo bool
}

func (poli *PoliticallyExposedCheck) Init(mp map[string]string) bool {

	temp, err := strconv.ParseBool(mp["Exposed"])
	if err != nil {
		log.Println("Error Initializing the Paramter for PoliticallyExposedCheck")
		return false
	}
	poli.poliExpo = temp
	return true
}

func (poli *PoliticallyExposedCheck) IsEligible(d *models.DecisionData) {
	var temp *bool
	temp = d.Record.PoliticallyExposed
	if *temp != poli.poliExpo {
		d.Eligible = false
		d.RejectionReason = append(d.RejectionReason, "PoliticallyExposedCheck")
	}
}
