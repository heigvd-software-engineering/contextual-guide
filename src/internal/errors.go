package internal

type ErrorDTO struct{
	Errors *ValidationError `json:"errors"`
}