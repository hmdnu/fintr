package transaction

type TransactionDto struct {
	AccountId  int    `json:"accountId" validate:"required"`
	CategoryId int    `json:"categoryId" validate:"required"`
	Amount     int    `json:"amount" validate:"required"`
	Note       string `json:"note"`
	Date       string `json:"date" validate:"required,datetime"`
}

type Transaction struct {
	Id         int    `json:"id" db:"id"`
	AccountId  int    `json:"accountId" db:"account_id"`
	Account    string `json:"account" db:"account"`
	CategoryId int    `json:"categoryId" db:"category_id"`
	Category   string `json:"category" db:"category"`
	Amount     int    `json:"amount" db:"amount"`
	Note       string `json:"note" db:"note"`
	Date       string `json:"date" db:"date"`
}
