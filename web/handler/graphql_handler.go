package handler

import (
	"encoding/json"
	"hextechdocs-be/web/graphql"
	"net/http"
)

func HandleGraphQL(w http.ResponseWriter, r *http.Request) {
	request := graphQLRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeBadRequest(w)
		return
	}

	result := graphql.ExecuteQuery(request.OperationName, request.Query, request.Variables)

	resBytes, err := json.Marshal(result)

	if err != nil {
		writeInternalServerError(w)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resBytes)
}

type graphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}
