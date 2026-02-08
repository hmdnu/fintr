package response

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Data    any
	Message string
	Status  int
	Error   any
}

func Ok(w http.ResponseWriter, res *HttpResponse) {
	json := map[string]any{
		"message": res.Message,
		"success": true,
		"data":    res.Data,
		"status":  res.Status,
	}
	writeResponse(w, json)
}

func Fail(w http.ResponseWriter, res *HttpResponse) {
	if res.Error != nil {
		jsonWithError := map[string]any{
			"message": res.Message,
			"success": false,
			"status":  res.Status,
			"error":   res.Error,
		}
		writeResponse(w, jsonWithError)
	}
	json := map[string]any{
		"message": res.Message,
		"success": false,
		"status":  res.Status,
	}
	writeResponse(w, json)
}

func writeResponse(w http.ResponseWriter, res map[string]any) {
	jsonByte, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}
	statusString := res["status"]
	statusInt := statusString.(int)
	w.WriteHeader(statusInt)
	w.Write(jsonByte)
}

func IntServError(w http.ResponseWriter) {
	json := map[string]any{
		"message": "internal server error",
		"success": false,
		"status":  500,
	}
	writeResponse(w, json)
}

func BadReqError(w http.ResponseWriter, validatorErr map[string]string) {
	json := map[string]any{
		"message": "validation error",
		"success": false,
		"error":   validatorErr,
		"status":  http.StatusBadRequest,
	}
	writeResponse(w, json)
}
