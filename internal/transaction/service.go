package transaction

import (
	"strings"

	"github.com/hmdnu/fintr/pkg/errorType"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(transactionDto *TransactionDto) error {
	query := "INSERT INTO transactions (account_id, category_id, amount, note, date) VALUES (:amount_id, :category_id, :amount, :note, :date)"
	tx := s.db.MustBegin()
	_, err := tx.NamedExec(query, &TransactionDto{AccountId: transactionDto.AccountId, CategoryId: transactionDto.CategoryId, Amount: transactionDto.Amount, Note: transactionDto.Note, Date: transactionDto.Date})
	if err != nil {
		if strings.Contains(err.Error(), "2067") {
			return errortype.ConstraintErr(err)
		}
	}
	tx.Commit()
	return nil
}

func (s *Service) List() ([]Transaction, error) {
	query := `
SELECT t.id, a.id, a.name as account, c.id ,c.name as category, t.amount, t.note, t.date FROM transactions as t
JOIN accounts as a ON t.account_id = a.id
JOIN categories as c ON t.category_id = c.id;
`
	transactions := []Transaction{}
	err := s.db.Select(&transactions, query)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// func (s *Service) Get(id int) (Transaction, error) {

// }
