/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedOn   time.Time `json:"createdOn,omitempty"`
	CreatedBy   string    `json:"createdBy,omitempty"`
	ModifiedOn  time.Time `json:"modifiedOn,omitempty"`
	ModifiedBy  string    `json:"modifiedBy,omitempty"`
	Priority    float32   `json:"priority"`
	Action      string    `json:"action"`
	Filter      []string  `json:"filter,omitempty"`
	External    []string  `json:"external,omitempty"`
	Internal    []string  `json:"internal,omitempty"`
}