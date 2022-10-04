package PoolOfChecks

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"reflect"
	"testing"
)

var PhCheck = PhoneAreaCodeCheck{areaCode: nil}

func TestPhoneAreaCodeCheck_Init(t *testing.T) {
	mp := map[string]string{
		"AreaCodes": "0,2,5,8"}
	PhCheck.Init(mp)
	got := PhCheck.areaCode
	var want = map[string]bool{
		"0": true,
		"2": true,
		"5": true,
		"8": true}
	if reflect.DeepEqual(got, want) {
		t.Logf("Passed")
	} else {
		t.Errorf("Failed")
	}
}

func TestPhoneAreaCodeCheck_IsEligible(t *testing.T) {
	PhCheck.areaCode = map[string]bool{
		"0": true,
		"2": true,
		"5": true,
		"8": true}
	cases := []models.DecisionData{
		{
			Record: &models.RecordData{
				Income:              0,
				NumberOfCreditCards: nil,
				Age:                 17,
				PoliticallyExposed:  nil,
				JobIndustryCode:     "",
				PhoneNumber:         "987", //fail
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
				PhoneNumber:         "8797", //Pass
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
				PhoneNumber:         "4657",
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
				PhoneNumber:         "5678",
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
				PhoneNumber:         "394775",
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
				PhoneNumber:         "0000",
			},
			RejectionReason: nil,
			Eligible:        true,
		},
	}

	want := []bool{false, true, false, true, false, true}
	i := 0
	for _, val := range cases {
		PhCheck.IsEligible(&val)
		if want[i] == val.Eligible {
			t.Logf("Passed")
		} else {
			t.Errorf("Failed")
		}
		i++
	}

}
