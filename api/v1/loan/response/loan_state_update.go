package response

import "loanapp/api/v1/loan/request"

type LoanStateResponseData struct {
	ID    int    `json:"id"`
	State string `json:"state"`
}

func NewLoanUpdateStateResponse(updateLoan request.UpdateLoanStateRequest) LoanStateResponseData {
	var loanState LoanStateResponseData

	loanState.ID = updateLoan.ID
	loanState.State = updateLoan.State

	return loanState
}
