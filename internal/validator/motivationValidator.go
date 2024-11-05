package validator

import (
	"github.com/thedevsaddam/govalidator"
	"motivations-api/pkg/modules/motivations"
	"net/url"
)

type MotivationValidator struct {
	validator *govalidator.Validator
}

func (m *MotivationValidator) New() *MotivationValidator {
	return &MotivationValidator{
		validator: &govalidator.Validator{
			Opts: govalidator.Options{},
		},
	}
}

func (m *MotivationValidator) Validate(rules govalidator.MapData, data interface{}) url.Values {
	m.validator.Opts.Data = data
	m.validator.Opts.Rules = rules

	e := m.validator.ValidateStruct()

	if len(e) > 0 {
		return e
	}
	return nil
}

func (m *MotivationValidator) ValidateCreatedMotivation(motivation *motivations.Motivation) url.Values {
	rules := govalidator.MapData{
		//TODO: create custom rule that supports cyrillic symbols
		"nickname":   []string{"required", "alpha_space"},
		"motivation": []string{"required", "alpha_space"},
	}

	return m.Validate(rules, motivation)
}

func (m *MotivationValidator) ValidateUpdatedMotivation(motivation *motivations.Motivation) url.Values {
	rules := govalidator.MapData{
		//TODO: create custom rule that supports cyrillic symbols
		"motivation": []string{"required", "alpha_space"},
	}

	return m.Validate(rules, motivation)
}
