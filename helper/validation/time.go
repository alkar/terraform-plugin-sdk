package validation

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// IsDayOfTheWeek id a SchemaValidateFunc which tests if the provided value is of type string and a valid english day of the week
func IsDayOfTheWeek(ignoreCase bool) schema.SchemaValidateFunc {
	return StringInSlice([]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}, ignoreCase)
}

// IsMonth id a SchemaValidateFunc which tests if the provided value is of type string and a valid english month
func IsMonth(ignoreCase bool) schema.SchemaValidateFunc {
	return StringInSlice([]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}, ignoreCase)
}

// IsRFC3339Time is a SchemaValidateFunc which tests if the provided value is of type string and a valid RFC33349Time
func IsRFC3339Time(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return
	}

	if _, err := time.Parse(time.RFC3339, v); err != nil {
		errors = append(errors, fmt.Errorf("expected %q to be a valid RFC3339 date, got %q: %+v", k, i, err))
	}

	return warnings, errors
}

// ValidateRFC3339TimeString is a ValidateFunc that ensures a string parses as time.RFC3339 format
//
// Deprecated: use IsRFC3339Time() instead
func ValidateRFC3339TimeString(v interface{}, k string) (ws []string, errors []error) {
	return IsRFC3339Time(v, k)
}