package process_transaction

type TransactionDtoInput struct {
	Id        string  `json:"id"`
	AccountId string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type TransactionDtoOutput struct {
	Id           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
