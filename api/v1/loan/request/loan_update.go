package request

import (
	"loanapp/business/loan"
)

type UpdateLoanRequest struct {
	ID      int `json:"id" validate:"required"`
	Amount  int `json:"amount" validate:"required"`
	Version int `json:"version" validate:"required"`
}

func (req *UpdateLoanRequest) ToUpdateLoanSpec() *loan.UpdateLoan {
	var updateLoanRequest loan.UpdateLoan

	updateLoanRequest.ID = req.ID
	updateLoanRequest.Amount = req.Amount
	updateLoanRequest.Version = req.Version

	return &updateLoanRequest
}
