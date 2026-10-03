package domain

import "github.com/shopspring/decimal"

type Statistics struct {
	Balance              decimal.Decimal
	TotalIncome          decimal.Decimal
	TotalSpending        decimal.Decimal
	MoneyLeftPrecentage  string
	MoneySpentPrecentage string
}
