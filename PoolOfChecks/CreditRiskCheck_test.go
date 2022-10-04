package PoolOfChecks

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"testing"
)

var CredCheck = CreditRiskCheck{noOfCards: 0}

func TestCreditRiskCheck_Init(t *testing.T) {
	mp := map[string]string{
		"NoOfCards": "4"}
	CredCheck.Init(mp)
	got := CredCheck.noOfCards
	want := 4
	if got == want {
		t.Logf("Passed")
	} else {
		t.Errorf("Failed")
	}
}

func TestCreditRiskCheck_IsEligible(t *testing.T) {
	CredCheck.noOfCards = 3
	case1 := 1
	case2 := 3
	case3 := 4
	cases := []models.DecisionData{
		{
			Record: &models.RecordData{
				Income:              0,
				NumberOfCreditCards: &case1,
				Age:                 9,
				PoliticallyExposed:  nil,
				JobIndustryCode:     "",
				PhoneNumber:         "",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
		{
			Record: &models.RecordData{
				Income:              0,
				NumberOfCreditCards: &case2,
				Age:                 18,
				PoliticallyExposed:  nil,
				JobIndustryCode:     "",
				PhoneNumber:         "",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
		{
			Record: &models.RecordData{
				Income:              0,
				NumberOfCreditCards: &case3,
				Age:                 49,
				PoliticallyExposed:  nil,
				JobIndustryCode:     "",
				PhoneNumber:         "",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
	}

	want := []bool{false, true, false}
	i := 0
	for _, val := range cases {
		CredCheck.IsEligible(&val)
		if want[i] == val.Eligible {
			t.Logf("Passed")
		} else {
			t.Logf("Failed")
		}
		i++
	}
}
