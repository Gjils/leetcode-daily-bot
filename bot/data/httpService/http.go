package httpService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GraphQLRequest представляет GraphQL запрос
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
 }
 
 // GraphQLResponse представляет GraphQL ответ
 type GraphQLResponse struct {
	Data   json.RawMessage            `json:"data"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
 }

func SendGraphQLRequest(url string, request GraphQLRequest) (*GraphQLResponse, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
	 return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
 
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
	 return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
 
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
	 return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()
 
	body, err := io.ReadAll(resp.Body)
	if err != nil {
	 return nil, fmt.Errorf("failed to read response body: %w", err)
	}
 
	var response GraphQLResponse
	if err := json.Unmarshal(body, &response); err != nil {
	 return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
 
	return &response, nil
 }