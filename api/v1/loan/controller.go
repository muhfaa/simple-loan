package loan

import (
	"loanapp/api/common"
	"loanapp/api/v1/loan/request"
	"loanapp/api/v1/loan/response"
	"loanapp/business"
	"loanapp/business/loan"
	"loanapp/util/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	loanService loan.Service
}

func NewController(loanService loan.Service) *Controller {
	return &Controller{
		loanService,
	}
}

// Insert Loan
func (controller Controller) InsertLoan(c echo.Context) error {
	insertLoan := new(request.InsertLoanRequest)

	if err := c.Bind(insertLoan); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(insertLoan); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	loanSpec := insertLoan.ToUpsertLoanSpec()
	err := controller.loanService.InsertLoan(*loanSpec)
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(business.ErrDuplicate))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}

// Get Loan
func (controller Controller) GetAllLoan(c echo.Context) error {

	loans, err := controller.loanService.FindAll()
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	successResponse := common.NewSuccessResponse(response.NewLoanDataResponse(loans))

	return c.JSON(http.StatusOK, successResponse)

}

// Update Loan
func (controller Controller) UpdateLoan(c echo.Context) error {
	updateLoan := new(request.UpdateLoanRequest)

	if err := c.Bind(updateLoan); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(updateLoan); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	loanSpec := updateLoan.ToUpdateLoanSpec()
	isUpdated, err := controller.loanService.UpdateLoan(*loanSpec)
	if err != nil && !isUpdated {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}

func (controller Controller) UpdateStateLoan(c echo.Context) error {
	updateLoanState := new(request.UpdateLoanStateRequest)

	if err := c.Bind(updateLoanState); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(updateLoanState); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	loanSpec := updateLoanState.ToUpdateStateLoanSpec()
	isUpdated, err := controller.loanService.ApprovalLoanState(*loanSpec)
	if err != nil && !isUpdated {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponse(response.NewLoanUpdateStateResponse(*updateLoanState)))
}
