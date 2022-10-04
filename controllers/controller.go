package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/honestbank/tech-assignment-backend-engineer/Cache"
	"github.com/honestbank/tech-assignment-backend-engineer/PoolOfChecks"
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"github.com/honestbank/tech-assignment-backend-engineer/strategy"
	"log"
	"net/http"
)

var ListOfChecks []PoolOfChecks.EligibilityCheck
var Lc Cache.AbsCache

func InitializeCheck(s string) {
	log.Println("Initiazlixing Checks")
	f := strategy.FileStrategyLoader{}
	ListOfChecks = f.LoadStrategy(s)
	Lc = Cache.NewLocalCache()
	fmt.Println(len(ListOfChecks))
}

//Controller too Handler Proces data calls to decison engine

func ProcessData(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var data models.RecordData
		result, stat := validProcessRequest(req, &data)
		if !stat {
			resp.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(resp).Encode(result)
			if err != nil {
				log.Println("Error Encoding Invalid validation message")
			}
			return
		}
		decisionData := models.DecisionData{
			Record:          &data,
			RejectionReason: make([]string, len(ListOfChecks)),
			Eligible:        true,
		}
		if !Lc.Get(decisionData.Record.PhoneNumber) {
			for _, element := range ListOfChecks {
				element.IsEligible(&decisionData)
			}
		} else {
			decisionData.Eligible = true
			decisionData.Overriden = true
		}

		var FinalStatus models.JsonResponse
		if decisionData.Eligible {
			FinalStatus.Status = "approved"
		} else {
			FinalStatus.Status = "declined"
		}
		log.Printf("Rejection Reason: %v,Overriden %v\n", decisionData.RejectionReason, decisionData.Overriden)
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(FinalStatus)

	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		log.Println(resp, "not found")
	}
}

//  Validating Process Data
func validProcessRequest(req *http.Request, data *models.RecordData) ([]string, bool) {
	if req.Body == nil {
		log.Println("Empty Body Passed")
		return []string{"Empty Body Passed"}, false
	}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Printf("Error : %v\n", err)
		return []string{"ill Formatted Body"}, false
	}
	v := validator.New()
	err = v.Struct(data)
	var validationOutput []string
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			validationOutput = append(validationOutput, er.Field()+":"+er.Tag())
			log.Printf("%v", validationOutput)
		}
		log.Println("Validation for structures failed")
		return validationOutput, false
	}
	return []string{"Ok"}, true
}

/*
	Controller to facillitate updating PreApprovedList of PhoneNumbers.
	/PATCH : To Update
	/DELETE : To delete a value from the list
*/

func ApprovedList(resp http.ResponseWriter, req *http.Request) {
	var ls models.ApprovedListModel
	st, flag := ValidateReqApprovdList(req, &ls)
	if !flag {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(st)
		return
	}
	switch req.Method {
	case http.MethodGet:
		res := Lc.Get(ls.PNumbers[0])
		json.NewEncoder(resp).Encode(res)
		return
	case http.MethodPatch:
		for i := 0; i < len(ls.PNumbers); i++ {
			Lc.Add(ls.PNumbers[i])
		}
		resp.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(resp).Encode("List Updated")
		return
	case http.MethodDelete:
		for i := 0; i < len(ls.PNumbers); i++ {
			Lc.Delete(ls.PNumbers[i])
		}
		resp.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(resp).Encode("List Updated")
		return
	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		log.Println(resp, "not found")
	}
}

//Validating Payload for updating pre-approved list of PhoneNumbers

func ValidateReqApprovdList(req *http.Request, ls *models.ApprovedListModel) (string, bool) {
	if req.Body == nil {
		return "Empty Body Passed", false
	}

	err := json.NewDecoder(req.Body).Decode(&ls)
	if err != nil {
		return "Ill Formed Body", false
	}
	if ls.IsEmpty() {
		return "Empty List Passed", false
	}
	log.Println("Request Vlidated")
	return "Ok", true
}
