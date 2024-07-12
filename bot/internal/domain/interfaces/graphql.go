package interfaces

import "encoding/json"

type GraphQLClient interface {
	SendGraphQLRequest(url string, request GraphQLRequest) (*GraphQLResponse, error)
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}
 
type GraphQLResponse struct {
	Data   json.RawMessage          `json:"data"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
}