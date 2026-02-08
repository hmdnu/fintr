package user

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var userDto CreateUserDto
	json.NewDecoder(r.Body).Decode(&userDto)
	err := validate.Struct(userDto)
	if err != nil {
		response.BadReqError(w, formatter.MapValidationErr(err))
		return
	}
	err = h.service.Create(userDto)
	if err != nil {
		if errors.Is(err, errortype.ConstraintErrType) {
			response.Fail(w, &response.HttpResponse{Message: err.Error(), Status: http.StatusConflict})
			return
		}
		response.IntServError(w)
		return
	}
	response.Ok(w, &response.HttpResponse{Message: "ok", Status: http.StatusOK})
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.List()
	if err != nil {
		fmt.Println(err)
		response.IntServError(w)
		return
	}
	response.Ok(w, &response.HttpResponse{Message: "success retrieved users", Status: http.StatusOK, Data: users})
}
