package repository

type TransactionRepository interface {
	Inser (ID string, account string, amout float64, status, string, errorMenssage string) error
}