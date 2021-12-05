package loan

import "time"

type State string

const (
	StateApproved State = "approved"
	StatePending  State = "pending"
	StateCancel   State = "cancel"
)

type Loan struct {
	ID        int       `json:"_" db:"id"`
	Amount    int       `json:"amount" db:"amount"`
	State     State     `json:"state" db:"state"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedAt"`
	Version   int       `json:"version" db:"version"`
}

func NewLoan(
	id int,
	amount int,
	state State,
	createdAt time.Time,
) Loan {

	return Loan{
		ID:        id,
		Amount:    amount,
		State:     state,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Version:   1,
	}
}

type UpdateLoan struct {
	ID      int
	Amount  int
	Version int
}

func NewLoanUpdate(
	id int,
	amount int,
	version int,
) UpdateLoan {

	return UpdateLoan{
		ID:      id,
		Amount:  amount,
		Version: version,
	}
}

// Update State in Loan
type UpdateLoanState struct {
	ID      int
	State   string
	Version int
}

func NewStateLoanUpdate(
	id int,
	state string,
	version int,
) UpdateLoanState {

	return UpdateLoanState{
		ID:      id,
		State:   state,
		Version: version,
	}
}
