package response

import (
	"loanapp/business/loan"
	"time"
)

type LoanResponseData struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Version   int       `json:"version"`
}

func NewLoanDataResponse(loans []loan.Loan) []LoanResponseData {
	var (
		loanResponseData LoanResponseData
		loansData        []LoanResponseData
	)

	if loans == nil {
		return loansData
	}

	for _, loan := range loans {

		loanResponseData.ID = loan.ID
		loanResponseData.Amount = loan.Amount
		loanResponseData.State = string(loan.State)
		loanResponseData.CreatedAt = loan.CreatedAt
		loanResponseData.UpdatedAt = loan.UpdatedAt
		loanResponseData.Version = loan.Version

		loansData = append(loansData, loanResponseData)
	}

	return loansData
}
