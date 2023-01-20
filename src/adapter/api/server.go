package api

import (
	"clean_architecture/src/application/process_transaction"
	"clean_architecture/src/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	repository domain.TransactionRepository
}

func NewWebServer(repository domain.TransactionRepository) *WebServer {
	return &WebServer{repository: repository}
}

func (w WebServer) Serve() {
	e := echo.New()

	e.POST("/transaction", w.process)

	e.Logger.Fatal(e.Start(":8080"))
}

func (w WebServer) process(c echo.Context) error {
	transactionDto := &process_transaction.TransactionDtoInput{}

	c.Bind(transactionDto)
	usecase := process_transaction.NewProcessTransaction(w.repository)

	output, err := usecase.Execute(*transactionDto)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, output)
}
