package service_payment

import (
	"net/http"
)

type BankAPI struct {
	url string
}

func NewBankAPI(url string) *BankAPI {
	return &BankAPI{url: url}
}

type PaymentService struct {
	bankAPI *BankAPI
}

func NewPaymentService(bankAPI *BankAPI) *PaymentService {
	return &PaymentService{bankAPI: bankAPI}
}

func (ps *PaymentService) HandlePayment(w http.ResponseWriter, r *http.Request) {
	// Обработка платежей
}
