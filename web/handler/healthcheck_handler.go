package handler

import (
	"hextechdocs-be/model"
	"net/http"
)

func HandleHealthcheck(w http.ResponseWriter) {
	if success, err := model.IsDatabaseAlive(); success {
		w.WriteHeader(http.StatusOK)
	} else if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		_, _ = w.Write([]byte("unknown error"))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
