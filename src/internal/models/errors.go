package models

type ErrorDTO struct{
	Errors *ValidationError `json:"errors"`
}