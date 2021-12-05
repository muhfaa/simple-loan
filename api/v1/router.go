package router

import (
	"loanapp/api/v1/loan"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	loanController loan.Controller,
) {

	// loan
	loanV1 := e.Group("v1/loan")
	loanV1.POST("/add", loanController.InsertLoan)
	loanV1.GET("", loanController.GetAllLoan)
	loanV1.PUT("/update", loanController.UpdateLoan)
	loanV1.PUT("/approval", loanController.UpdateStateLoan)

}
