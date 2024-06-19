package validator

import (
	"fmt"
	"regexp"
)

// Validator struct to hold value and errors.
type Validator struct {
	value  string
	errors []string
}

// Validate creates a new Validator instance.
func Validate(value string) *Validator {
	return &Validator{value: value, errors: []string{}}
}

// Email adds an email validation to the chain.
func (v *Validator) Email() *Validator {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !re.MatchString(v.value) {
		v.errors = append(v.errors, fmt.Sprintf("invalid email"))
	}
	return v
}

// MinLength adds a min length validation to the chain.
func (v *Validator) MinLength(minLength int) *Validator {
	if len(v.value) < minLength {
		v.errors = append(v.errors, fmt.Sprintf("%s is shorter than %d", v.value, minLength))
	}
	return v
}

// MaxLength adds a max length validation to the chain.
func (v *Validator) MaxLength(maxLength int) *Validator {
	if len(v.value) > maxLength {
		v.errors = append(v.errors, fmt.Sprintf("%s is longer than %d", v.value, maxLength))
	}
	return v
}

// NonEmpty adds a non-empty validation to the chain.
func (v *Validator) NonEmpty() *Validator {
	if len(v.value) == 0 {
		v.errors = append(v.errors, fmt.Sprintf("%s is empty", v.value))
	}
	return v
}

// Phone adds a phone number validation to the chain.
func (v *Validator) Phone() *Validator {
	re := regexp.MustCompile(`^09\d{9}$`)
	if !re.MatchString(v.value) {
		v.errors = append(v.errors, fmt.Sprintf("%s is invalid phone number", v.value))
	}
	return v
}

// Errors returns the validation errors.
func (v *Validator) Errors() []string {
	return v.errors
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}
