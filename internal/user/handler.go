package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
	var userDto CreateUserDto
	json.NewDecoder(r.Body).Decode(&userDto)
	err := validate.Struct(userDto)
	if err != nil {
		response.BadReqError(w, formatter.MapValidationErr(err))
		return err
	}
	err = h.service.Create(userDto)
	if err != nil {
		if errors.Is(err, errortype.ConstraintErrType) {
			response.Fail(w, &response.HttpFail{Message: err.Error(), Status: http.StatusConflict})
			return err
		}
		response.IntServError(w)
		return err
	}
	response.Ok(w, &response.HttpOk{Message: "ok", Status: http.StatusOK})
	return nil
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) error {
	users, err := h.service.List()
	if err != nil {
		fmt.Println(err)
		response.IntServError(w)
		return err
	}
	response.Ok(w, &response.HttpOk{Message: "success retrieved users", Status: http.StatusOK, Data: users})
	return nil
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) error {
	idInt, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.BadReqError(w, map[string]string{"id": "must be int"})
		return err
	}
	user, err := h.service.Get(idInt)
	if err != nil {
		response.IntServError(w)
		return err
	}
	response.Ok(w, &response.HttpOk{Data: user, Message: "user retrieved", Status: http.StatusOK})
	return nil
}
