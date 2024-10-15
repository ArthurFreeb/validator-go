package schema

import (
	"fmt"
	"net/url"
	"strings"
)

type StringPointerSchema struct {
	Schema[*string]
}

var _ ISchema = (*StringPointerSchema)(nil)

func StringPointer() *StringPointerSchema {
	return &StringPointerSchema{}
}

func (s *StringPointerSchema) Max(maxLength int) *StringPointerSchema {
	validator := Validator[*string]{
		MessageFunc: func(value *string) string {
			return fmt.Sprintf("String must contain at most %d character(s)", maxLength)
		},
		ValidateFunc: func(value *string) bool {
			return value == nil || len(*value) <= maxLength
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *StringPointerSchema) Min(minLength int) *StringPointerSchema {
	validator := Validator[*string]{
		MessageFunc: func(value *string) string {
			return fmt.Sprintf("String must contain at least %d character(s)", minLength)
		},
		ValidateFunc: func(value *string) bool {
			return value == nil || len(*value) >= minLength
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *StringPointerSchema) Length(length int) *StringPointerSchema {
	validator := Validator[*string]{
		MessageFunc: func(value *string) string {
			return fmt.Sprintf("String must contain exactly %d character(s)", length)
		},
		ValidateFunc: func(value *string) bool {
			return value == nil || len(*value) == length
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *StringPointerSchema) Url() *StringPointerSchema {
	validator := Validator[*string]{
		MessageFunc: func(value *string) string {
			return "Invalid url"
		},
		ValidateFunc: func(value *string) bool {
			uri, err := url.ParseRequestURI(*value)
			return value == nil || (err == nil && uri.Host != "")
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *StringPointerSchema) Includes(str string) *StringPointerSchema {
	validator := Validator[*string]{
		MessageFunc: func(value *string) string {
			return fmt.Sprintf("Invalid input: must include \"%s\"", str)
		},
		ValidateFunc: func(value *string) bool {
			return value == nil || strings.Contains(*value, str)
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *StringPointerSchema) StartsWith(str string) *StringPointerSchema {
	validator := Validator[*string]{
		MessageFunc: func(value *string) string {
			return fmt.Sprintf("Invalid input: must start with \"%s\"", str)
		},
		ValidateFunc: func(value *string) bool {
			return value == nil || strings.HasPrefix(*value, str)
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *StringPointerSchema) EndsWith(str string) *StringPointerSchema {
	validator := Validator[*string]{
		MessageFunc: func(value *string) string {
			return fmt.Sprintf("Invalid input: must end with \"%s\"", str)
		},
		ValidateFunc: func(value *string) bool {
			return value == nil || strings.HasSuffix(*value, str)
		},
	}

	s.validators = append(s.validators, validator)

	return s
}
