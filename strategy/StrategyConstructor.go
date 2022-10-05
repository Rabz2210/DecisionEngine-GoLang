package strategy

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/PoolOfChecks"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

type Loader interface {
	LoadStrategy(string) []PoolOfChecks.EligibilityCheck
}

type FileStrategyLoader struct {
	fileNamePath string
}

func (l *FileStrategyLoader) Init() {
	l.fileNamePath = "DefinedStrategy"
}

//Function to load strategy from a file based on the strategy passed

func (l *FileStrategyLoader) LoadStrategy(strategy string) []PoolOfChecks.EligibilityCheck {
	log.Println("File Loader Loading Strategy")
	mapOfChecks, err := readStrategyFromFile(strategy)
	if err != nil {
		log.Println("Server cannot proceed further without Reading in  a proper Strategy")
		err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		if err != nil {
			panic(err)
		}
	}
	var ls []PoolOfChecks.EligibilityCheck
	for checkName, params := range mapOfChecks {
		check1, err := ObjectConstructor(checkName, params)
		if err != nil {
			log.Println(err)
			log.Println("Server cannot proceed further without loading all the checks properly")
			err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			if err != nil {
				panic(err)
			}
		}
		ls = append(ls, check1)
		log.Printf("Successfully added check %s, PoolOfChecks size : %d\n", checkName, len(ls))
	}
	return ls
}

//To open and read the file

func readStrategyFromFile(strategy string) (map[string]map[string]string, error) {
	// Open  jsonFile
	jsonFile, err := os.Open(strategy + ".json")
	if err != nil {
		log.Println(err)
		return nil, errors.New("Strategy is not defined ")

	}
	fmt.Printf("Successfully Opened %s.json", strategy)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("Error Reading file for loading Strategy")
		return nil, err
	}
	var x map[string]map[string]string
	err = json.Unmarshal(byteValue, &x)
	if err != nil {
		log.Printf("Error Reading File: %v", err)
		return nil, err
	}
	log.Printf("%v", x)
	return x, nil
}

//Function to construct Objects at runtime based on the check configured in the strategy file.

func ObjectConstructor(name string, params map[string]string) (PoolOfChecks.EligibilityCheck, error) {
	var obj PoolOfChecks.EligibilityCheck
	var iniResult bool
	switch name {
	case "IncomeCheck":
		obj = &PoolOfChecks.IncomeCheck{}
		iniResult = obj.Init(params)
	case "CreditRiskCheck":
		obj = &PoolOfChecks.CreditRiskCheck{}
		iniResult = obj.Init(params)
	case "AgeCheck":
		obj = &PoolOfChecks.AgeCheck{}
		iniResult = obj.Init(params)
	case "PoliticallyExposed":
		obj = &PoolOfChecks.PoliticallyExposedCheck{}
		iniResult = obj.Init(params)
	case "AcceptedAreaCodes":
		obj = &PoolOfChecks.PhoneAreaCodeCheck{}
		iniResult = obj.Init(params)
	default:
		errorString := "Object Constructor : No Configuration to create check : " + name
		return nil, errors.New(errorString)
	}
	if !iniResult {
		errorString := "Object Constructor : Parameter Initialization failed for  check : " + name
		return nil, errors.New(errorString)
	}
	return obj, nil
}
