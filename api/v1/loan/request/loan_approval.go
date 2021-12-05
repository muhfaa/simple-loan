package request

import (
	"loanapp/business/loan"
)

type UpdateLoanStateRequest struct {
	ID      int    `json:"id"`
	State   string `json:"state"`
	Version int    `json:"version"`
}

func (req *UpdateLoanStateRequest) ToUpdateStateLoanSpec() *loan.UpdateLoanState {
	var updateLoan loan.UpdateLoanState

	updateLoan.ID = req.ID
	updateLoan.State = req.State
	updateLoan.Version = req.Version

	return &updateLoan
}
