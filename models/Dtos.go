package models

type RecordData struct {
	Income              int    `json:"income" validate:"required"`
	NumberOfCreditCards *int   `json:"number_of_credit_cards" validate:"required"`
	Age                 int    `json:"age"  validate:"required,max=130"`
	PoliticallyExposed  *bool  `json:"politically_exposed" validate:"required"`
	JobIndustryCode     string `json:"job_industry_code"`
	PhoneNumber         string `json:"phone_number" validate:"required,min=5,max=15"`
}

type JsonResponse struct {
	Status string `json:"status"`
}

type ApprovedListModel struct {
	PNumbers []string `json:"phone_numbers" validate:"required,min=1"`
}

func (l *ApprovedListModel) IsEmpty() bool {
	if len(l.PNumbers) == 0 {
		return true
	}
	return false
}

type DecisionData struct {
	Record          *RecordData
	RejectionReason []string
	Eligible        bool
	Overriden       bool
}
