package PoolOfChecks

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"testing"
)

var AgCheck = AgeCheck{age: 0}

func TestAgeCheck_Init(t *testing.T) {

	mp := map[string]string{
		"Age": "18"}
	AgCheck.Init(mp)
	got := AgCheck.age
	want := 18
	if got == want {
		t.Logf("Age Check Initialization Passed")
	} else {
		t.Errorf("Age Check Initialization failed")
	}
}

func TestAgeCheck_IsEligible(t *testing.T) {
	AgCheck.age = 18
	cases := []models.DecisionData{
		{
			Record: &models.RecordData{
				Income:              0,
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
				Income:              0,
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
				Income:              0,
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
				Income:              0,
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

	want := []bool{false, false, true, true}
	i := 0
	for _, val := range cases {
		AgCheck.IsEligible(&val)
		if want[i] == val.Eligible {
			t.Logf("passed")
		} else {
			t.Errorf("Fail")
		}
		i++
	}
}
