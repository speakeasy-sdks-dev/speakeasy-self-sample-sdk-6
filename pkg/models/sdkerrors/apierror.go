// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type APIError struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}

var _ error = &APIError{}

func (e *APIError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
