package transaction

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	errortype "github.com/hmdnu/fintr/pkg/errorType"
	"github.com/hmdnu/fintr/pkg/formatter"
	"github.com/hmdnu/fintr/pkg/response"
)

type Handler struct {
	service *Service
}

var validate = validator.New()

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	var transactionDto TransactionDto
	json.NewDecoder(r.Body).Decode(&transactionDto)
	err := validate.Struct(transactionDto)
	if err != nil {
		response.BadReqError(w, formatter.MapValidationErr(err))
		return err
	}
	err = h.service.Create(&transactionDto)
	if err != nil {
		if errors.Is(err, errortype.ConstraintErrType) {
			response.Fail(w, &response.HttpFail{Message: err.Error(), Status: http.StatusConflict})
			return err
		}
		return err
	}
	response.Ok(w, &response.HttpOk{Message: "ok", Status: http.StatusCreated})
	return nil
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) error {
	transactions, err := h.service.List()
	if err != nil {
		response.IntServError(w)
		return err
	}
	response.Ok(w, &response.HttpOk{Data: transactions, Message: "list of transactions retrieved", Status: http.StatusOK})
	return nil
}
