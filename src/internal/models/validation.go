package models

import (
	"fmt"
	"regexp"
)

type ValidationError map[string][]string

func notEmpty(property string, value string, validationError *ValidationError) *ValidationError {
	if value == "" {
		message := fmt.Sprintf("%s is required", property)
		(*validationError)[property] = append((*validationError)[property], message)
	}
	return validationError
}

// float value between -180 and 180
func inLatitudeBoundary(property string, value float32, validationError *ValidationError) *ValidationError {
	if value < -180 || value > 180 {
		message := fmt.Sprintf("%s is not in latitude boundary, between -90 and 90", property)
		(*validationError)[property] = append((*validationError)[property], message)
	}
	return validationError
}

// float value between -90 and 90
func inLongitudeBoundary(property string, value float32, validationError *ValidationError) *ValidationError {
	if value > 90 || value < -90 {
		message := fmt.Sprintf("%s is not in longitude boundary, between -180 an 180", property)
		(*validationError)[property] = append((*validationError)[property], message)
	}
	return validationError
}

func isUrlFormat(property string, value string, validationError *ValidationError) *ValidationError {
	urlRegex := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)

	if urlRegex.Match([]byte(value)) {
		message := fmt.Sprintf("%s is not in url format, http(s)://xxx.xx", property)
		(*validationError)[property] = append((*validationError)[property], message)
	}
	return validationError
}
