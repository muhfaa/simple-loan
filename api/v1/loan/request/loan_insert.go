package request

import "loanapp/business/loan"

type InsertLoanRequest struct {
	ID     int `json:"id"`
	Amount int `json:"amount" validate:"required"`
}

func (req *InsertLoanRequest) ToUpsertLoanSpec() *loan.InsertLoanSpec {
	var loanSpec loan.InsertLoanSpec

	loanSpec.ID = req.ID
	loanSpec.Amount = req.Amount

	return &loanSpec

}
