package PoolOfChecks

import "github.com/honestbank/tech-assignment-backend-engineer/models"

//Interface to be implemented by every check to be configured into the decision engine

type EligibilityCheck interface {
	Init(map[string]string) bool
	IsEligible(d *models.DecisionData)
}
