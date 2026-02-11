package auth

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

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

var validate = validator.New()

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var authDto AuthDto
	json.NewDecoder(r.Body).Decode(&authDto)
	err := validate.Struct(authDto)
	if err != nil {
		fmt.Println(err)
		response.BadReqError(w, formatter.MapValidationErr(err))
		return
	}
	token, err := h.service.Login(authDto)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, errortype.CredInvalidErr) {
			response.Fail(w, &response.HttpResponse{Message: err.Error(), Status: http.StatusUnauthorized})
			return
		}
		response.IntServError(w)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   3600,
	})
	response.Ok(w, &response.HttpResponse{Message: "login success", Status: http.StatusOK})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Path:     "/",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	response.Ok(w, &response.HttpResponse{Message: "logout success", Status: http.StatusOK})
}
