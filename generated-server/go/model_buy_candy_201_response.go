// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Candy Server
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi

type BuyCandy201Response struct {
	Thanks string `json:"thanks,omitempty"`
	Change int32  `json:"change,omitempty"`
}

// AssertBuyCandy201ResponseRequired checks if the required fields are not zero-ed
func AssertBuyCandy201ResponseRequired(obj BuyCandy201Response) error {
	return nil
}

// AssertBuyCandy201ResponseConstraints checks if the values respects the defined constraints
func AssertBuyCandy201ResponseConstraints(obj BuyCandy201Response) error {
	return nil
}
