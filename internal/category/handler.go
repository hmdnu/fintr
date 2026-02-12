package category

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

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

var validate = validator.New()

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	var categoryDto CategoryDto
	json.NewDecoder(r.Body).Decode(&categoryDto)
	err := validate.Struct(&categoryDto)
	if err != nil {
		response.BadReqError(w, formatter.MapValidationErr(err))
		return err
	}
	err = h.service.Create(&categoryDto)
	if err != nil {
		if errors.Is(err, errortype.ConstraintErrType) {
			response.DuplicateErr(w, err.Error())
			return err
		}
	}
	response.Ok(w, &response.HttpOk{Message: "category created", Status: http.StatusCreated})
	return nil
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) error {
	return nil
}
