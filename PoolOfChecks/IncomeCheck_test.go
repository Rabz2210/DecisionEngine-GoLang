package PoolOfChecks

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"testing"
)

var IncCheck = IncomeCheck{income: 0}

func TestIncomeCheck_Init(t *testing.T) {
	mp := map[string]string{
		"Income": "100000"}
	IncCheck.Init(mp)
	got := IncCheck.income
	if got == 100000 {
		t.Logf("Passed")
	} else {
		t.Errorf("Failed")
	}
}

func TestIncomeCheck_IsEligible(t *testing.T) {
	IncCheck.income = 100000
	cases := []models.DecisionData{
		{
			Record: &models.RecordData{
				Income:              764,
				NumberOfCreditCards: nil,
				Age:                 17,
				PoliticallyExposed:  nil,
				JobIndustryCode:     "",
				PhoneNumber:         "",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
		{
			Record: &models.RecordData{
				Income:              100001,
				NumberOfCreditCards: nil,
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
				Income:              100000,
				NumberOfCreditCards: nil,
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
				Income:              99999,
				NumberOfCreditCards: nil,
				Age:                 49,
				PoliticallyExposed:  nil,
				JobIndustryCode:     "",
				PhoneNumber:         "",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
	}

	want := []bool{false, true, false, false}
	i := 0
	for _, val := range cases {
		IncCheck.IsEligible(&val)
		if want[i] == val.Eligible {
			t.Logf("Passed")
		} else {
			t.Errorf("Failed")
		}
		i++
	}
}
