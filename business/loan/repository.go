package loan

type Repository interface {
	InsertLoan(Loan) error
	FindLoanByID(ID int) (*Loan, error)
	FindAllLoan() ([]Loan, error)
	UpdateLoan(id int, amount int, currentVersion int) (bool, error)
	ApprovalLoanState(id int, state string, currentVersion int) (bool, error)
}
