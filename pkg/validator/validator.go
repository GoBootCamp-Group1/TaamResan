package validator

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
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

func (v *Validator) Date() *Validator {
	re := regexp.MustCompile(`^d{4}-d{2}-d{2}$`)
	if !re.MatchString(v.value) {
		v.errors = append(v.errors, fmt.Sprintf("%s is invalid date YYYY-mm-dd", v.value))
	}
	return v
}

func (v *Validator) Password() *Validator {
	if len(v.value) < 8 {
		v.errors = append(v.errors, "password must be at least 8 characters long")
	}

	hasUpper := false
	hasSpecial := false
	for _, char := range v.value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		v.errors = append(v.errors, "password must have at least one uppercase letter")
	}
	if !hasSpecial {
		v.errors = append(v.errors, "password must have at least one special character")
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

func (v *Validator) IsFloat() *Validator {
	if _, err := strconv.ParseFloat(v.value, 10); err != nil || v.value == "" {
		v.errors = append(v.errors, fmt.Sprintf("%s is invalid float", v.value))
	}
	return v
}
