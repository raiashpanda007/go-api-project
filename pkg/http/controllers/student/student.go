package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/raiashpanda007/go-api-project/pkg/types"
	"github.com/raiashpanda007/go-api-project/pkg/utils"
)

type Response struct {
	Status string
	Error  string
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func Create() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var student types.Student
		err := json.NewDecoder(request.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			utils.WriteJson(response, http.StatusBadRequest, err.Error())
			return
		}
		slog.Info("Creating a student data ")

		if err != nil {
			utils.WriteJson(response, http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}
		// request validation ::
		err = validator.New().Struct(student)
		if err != nil {
			validateErrs := err.(validator.ValidationErrors)
			utils.WriteJson(response, http.StatusBadRequest, utils.ValidatorResponse(validateErrs))
			return
		}

		utils.WriteJson(response, http.StatusCreated, map[string]string{"success": "ok"})

	}
}
