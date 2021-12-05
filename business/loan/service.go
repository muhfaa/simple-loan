package loan

import (
	"loanapp/business"
	"time"
)

type InsertLoanSpec struct {
	ID     int
	Amount int
}

type Service interface {
	InsertLoan(insertLoanSpec InsertLoanSpec) error
	FindLoanByID(id int) (*Loan, error)
	FindAll() ([]Loan, error)
	UpdateLoan(UpdateLoanSpec UpdateLoan) (bool, error)
	ApprovalLoanState(updateLoanStateSpec UpdateLoanState) (bool, error)
}

type service struct {
	loanRepository Repository
}

func NewService(loanRepository Repository) Service {

	return &service{
		loanRepository,
	}
}

func (s *service) InsertLoan(insertLoanSpec InsertLoanSpec) error {
	// existingLoan, err := s.FindLoanByID(insertLoanSpec.ID)
	// if err != nil && err != business.ErrNotFound {
	// 	return err
	// } else if existingLoan != nil {
	// 	return business.ErrDuplicate
	// }

	loan := NewLoan(
		insertLoanSpec.ID,
		insertLoanSpec.Amount,
		StatePending,
		time.Now(),
	)

	err := s.loanRepository.InsertLoan(loan)
	if err != nil && err != business.ErrNotFound {
		return err
	}

	return nil
}

func (s *service) FindLoanByID(id int) (*Loan, error) {
	loan, err := s.loanRepository.FindLoanByID(id)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (s *service) FindAll() ([]Loan, error) {
	loans, err := s.loanRepository.FindAllLoan()
	if err != nil {
		return loans, err
	}

	return loans, nil
}

func (s *service) UpdateLoan(updateLoanSpec UpdateLoan) (bool, error) {
	existingLoan, err := s.FindLoanByID(updateLoanSpec.ID)
	if err != nil {
		return false, err
	} else if existingLoan == nil {
		return false, business.ErrNotFound
	} else if existingLoan.Version != updateLoanSpec.Version {
		return false, business.ErrHasBeenModified
	}

	result, err := s.loanRepository.UpdateLoan(updateLoanSpec.ID, updateLoanSpec.Amount, updateLoanSpec.Version)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *service) ApprovalLoanState(updateLoanStateSpec UpdateLoanState) (bool, error) {
	existingLoan, err := s.FindLoanByID(updateLoanStateSpec.ID)

	if err != nil {
		return false, err
	} else if existingLoan == nil {
		return false, business.ErrNotFound
	} else if existingLoan.Version != updateLoanStateSpec.Version {
		return false, business.ErrHasBeenModified
	} else if existingLoan.State != StatePending && updateLoanStateSpec.State == string(StatePending) {
		return false, nil
	}

	result, err := s.loanRepository.ApprovalLoanState(updateLoanStateSpec.ID, updateLoanStateSpec.State, updateLoanStateSpec.Version)
	if err != nil {
		return false, err
	}

	return result, nil
}
