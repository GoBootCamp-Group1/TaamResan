package validator

import "testing"

func TestValidator(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"ValidEmail", "test@example.com", false},
		{"InvalidEmail", "invalid-email", true},
		{"ValidMinLength", "hello", false},
		{"InvalidMinLength", "hi", true},
		{"ValidMaxLength", "hello", false},
		{"InvalidMaxLength", "hello world", true},
		{"NonEmpty", "non-empty", false},
		{"Empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Validate(tt.value)
			switch tt.name {
			case "ValidEmail", "InvalidEmail":
				v = v.Email()
			case "ValidMinLength", "InvalidMinLength":
				v = v.MinLength(3)
			case "ValidMaxLength", "InvalidMaxLength":
				v = v.MaxLength(10)
			case "NonEmpty", "Empty":
				v = v.NonEmpty()
			}

			if (len(v.Errors()) > 0) != tt.wantErr {
				t.Errorf("got errors = %v, wantErr %v", v.Errors(), tt.wantErr)
			}
		})
	}
}
