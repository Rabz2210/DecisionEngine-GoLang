package PoolOfChecks

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"testing"
)

var poliExpoCheck = PoliticallyExposedCheck{poliExpo: true}

func TestPoliticallyExposedCheck_Init(t *testing.T) {
	mp := map[string]string{
		"Exposed": "false"}
	poliExpoCheck.Init(mp)
	got := poliExpoCheck.poliExpo
	if got == false {
		t.Logf("Passed")
	} else {
		t.Errorf("Failed")
	}
}

func TestPoliticallyExposedCheck_IsEligible(t *testing.T) {
	poliExpoCheck.poliExpo = false
	f := false
	tr := true
	cases := []models.DecisionData{
		{
			Record: &models.RecordData{
				Income:              0,
				NumberOfCreditCards: nil,
				Age:                 17,
				PoliticallyExposed:  &f,
				JobIndustryCode:     "",
				PhoneNumber:         "",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
		{
			Record: &models.RecordData{
				Income:              0,
				NumberOfCreditCards: nil,
				Age:                 9,
				PoliticallyExposed:  &tr,
				JobIndustryCode:     "",
				PhoneNumber:         "",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
		{
			Record: &models.RecordData{
				Income:              0,
				NumberOfCreditCards: nil,
				Age:                 18,
				PoliticallyExposed:  &f,
				JobIndustryCode:     "",
				PhoneNumber:         "",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
		{
			Record: &models.RecordData{
				Income:              0,
				NumberOfCreditCards: nil,
				Age:                 49,
				PoliticallyExposed:  &tr,
				JobIndustryCode:     "",
				PhoneNumber:         "",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
	}

	want := []bool{true, false, true, false}
	i := 0
	for _, val := range cases {
		poliExpoCheck.IsEligible(&val)
		if want[i] == val.Eligible {
			t.Logf("Passed")
		} else {
			t.Errorf("Failed")
		}
		i++
	}
}
