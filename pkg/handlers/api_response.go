package handlers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGetewayProxyResponse, error) {
   
	resp := events.APIGetewayProxyResponse{Headers: map[string]string["Content-Type":"application/json"]}
	 
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil 

}
