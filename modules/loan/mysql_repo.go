package loan

import (
	"database/sql"
	"errors"
	"loanapp/business"
	loanBusiness "loanapp/business/loan"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type MySQLDBRepository struct {
	db *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) *MySQLDBRepository {
	return &MySQLDBRepository{
		db,
	}
}

func (repo *MySQLDBRepository) InsertLoan(loanSpec loanBusiness.Loan) error {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}
	insertQuery := `INSERT INTO loan (
		amount,
		state,
		createdAt,
		updatedAt,
		version
	)
	VALUES
	(?,?,?,?,?)`

	now := time.Now().Local().Format("2006-01-02 15:04:05")
	_, err = tx.Exec(insertQuery, loanSpec.Amount, loanSpec.State, now, now, loanSpec.Version)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("error duplicate entry data")
			return err
		}

		err = errors.New("resource error")
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}

	return nil
}

func (repo *MySQLDBRepository) FindLoanByID(ID int) (*loanBusiness.Loan, error) {
	var loan loanBusiness.Loan

	selectQuery := `SELECT * FROM loan WHERE id = ?`

	err := repo.db.QueryRowx(selectQuery, ID).StructScan(&loan)
	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, business.ErrNotFound
		}
		err = errors.New("resource error")
		return nil, err
	}

	return &loan, nil
}

func (repo *MySQLDBRepository) FindAllLoan() ([]loanBusiness.Loan, error) {
	var loan loanBusiness.Loan
	var loans []loanBusiness.Loan

	selectQuery := `SELECT * FROM loan`

	row, err := repo.db.Queryx(selectQuery)
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return nil, err
	}

	for row.Next() {
		err = row.StructScan(&loan)
		if err != nil {
			log.Error(err)
			err = errors.New("resource error")
			return nil, err
		}

		loans = append(loans, loan)
	}
	return loans, nil
}

// UpdateLoan
func (repo *MySQLDBRepository) UpdateLoan(id int, amount int, currentVersion int) (bool, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	insertQuery := `UPDATE loan 
		SET
		amount = ?,
		updatedAt = ?,
		version = ?
		WHERE
		id = ?`

	now := time.Now().Local().Format("2006-01-02 15:04:05")
	_, err = tx.Exec(insertQuery, amount, now, currentVersion+1, id)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("error duplicate entry data")
			return false, err
		}

		err = errors.New("resource error")
		return false, err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	return true, nil
}

// ApprovalLoanState
func (repo *MySQLDBRepository) ApprovalLoanState(id int, state string, currentVersion int) (bool, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	insertQuery := `UPDATE loan 
		SET
		state = ?,
		updatedAt = ?,
		version = ?
		WHERE
		id = ?`

	now := time.Now().Local().Format("2006-01-02 15:04:05")
	_, err = tx.Exec(insertQuery, state, now, currentVersion+1, id)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("error duplicate entry data")
			return false, err
		}

		err = errors.New("resource error")
		return false, err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	return true, nil
}
