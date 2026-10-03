package statistics_transport_http

import (
	"github.com/qury53/finapp/internal/core/domain"
)

type StatisticsResponseDTO struct {
	Balance              string `json:"balance"`
	TotalIncome          string `json:"total_income"`
	TotalSpending        string `json:"total_spending"`
	MoneyLeftPrecentage  string `json:"money_left_precentage"`
	MoneySpentPrecentage string `json:"money_spent_precentage"`
}

func dtoFromDomain(domain domain.Statistics) StatisticsResponseDTO {
	return StatisticsResponseDTO{
		Balance:              domain.Balance.String(),
		TotalIncome:          domain.TotalIncome.String(),
		TotalSpending:        domain.TotalSpending.String(),
		MoneyLeftPrecentage:  domain.MoneyLeftPrecentage,
		MoneySpentPrecentage: domain.MoneySpentPrecentage,
	}
}
