package schema

import (
    "fmt"
)

type IntPointerSchema struct {
    Schema[*int]
}

var _ ISchema = (*IntPointerSchema)(nil)

func IntPointer() *IntPointerSchema {
    return &IntPointerSchema{}
}

func (s *IntPointerSchema) Lt(max int) *IntPointerSchema {
    validator := Validator[*int]{
        MessageFunc: func(value *int) string {
            return fmt.Sprintf("Int must be less than %d", max)
        },
        ValidateFunc: func(value *int) bool {
            return value != nil && *value < max
        },
    }

    s.validators = append(s.validators, validator)

    return s
}

func (s *IntPointerSchema) Lte(max int) *IntPointerSchema {
    validator := Validator[*int]{
        MessageFunc: func(value *int) string {
            return fmt.Sprintf("Int must be less than or equal to %d", max)
        },
        ValidateFunc: func(value *int) bool {
            return value != nil && *value <= max
        },
    }

    s.validators = append(s.validators, validator)

    return s
}

func (s *IntPointerSchema) Gt(min int) *IntPointerSchema {
    validator := Validator[*int]{
        MessageFunc: func(value *int) string {
            return fmt.Sprintf("Int must be greater than %d", min)
        },
        ValidateFunc: func(value *int) bool {
            return value != nil && *value > min
        },
    }

    s.validators = append(s.validators, validator)

    return s
}

func (s *IntPointerSchema) Gte(min int) *IntPointerSchema {
    validator := Validator[*int]{
        MessageFunc: func(value *int) string {
            return fmt.Sprintf("Int must be greater than or equal to %d", min)
        },
        ValidateFunc: func(value *int) bool {
            return value != nil && *value >= min
        },
    }

    s.validators = append(s.validators, validator)

    return s
}

func (s *IntPointerSchema) Positive() *IntPointerSchema {
	validator := Validator[*int]{
		MessageFunc: func(value *int) string {
			return "Int must be greater than 0"
		},
		ValidateFunc: func(value *int) bool {
			return value != nil && *value > 0
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *IntPointerSchema) Nonnegative() *IntPointerSchema {
	validator := Validator[*int]{
		MessageFunc: func(value *int) string {
			return "Int must be greater than or equal to 0"
		},
		ValidateFunc: func(value *int) bool {
			return value != nil && *value >= 0
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *IntPointerSchema) Negative() *IntPointerSchema {
	validator := Validator[*int]{
		MessageFunc: func(value *int) string {
			return "Int must be less than 0"
		},
		ValidateFunc: func(value *int) bool {
			return value != nil && *value < 0
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *IntPointerSchema) Nonpositive() *IntPointerSchema {
	validator := Validator[*int]{
		MessageFunc: func(value *int) string {
			return "Int must be less than or equal to 0"
		},
		ValidateFunc: func(value *int) bool {
			return value != nil && *value <= 0
		},
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *IntPointerSchema) MultipleOf(factor int) *IntPointerSchema {
	validator := Validator[*int]{
		MessageFunc: func(value *int) string {
			return fmt.Sprintf("Int must be a multiple of %d", factor)
		},
		ValidateFunc: func(value *int) bool {
			return value != nil && *value%factor == 0
		},
	}

	s.validators = append(s.validators, validator)

	return s
}