package validator

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"motivations-api/pkg/modules/motivations"
	"net/url"
)

type MotivationValidator struct {
	validator *govalidator.Validator
}

func (m *MotivationValidator) New() *MotivationValidator {
	addCustomRules()

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
		"nickname":   []string{"text", "required"},
		"motivation": []string{"text", "required"},
	}

	return m.Validate(rules, motivation)
}

func (m *MotivationValidator) ValidateUpdatedMotivation(motivation *motivations.Motivation) url.Values {
	rules := govalidator.MapData{
		"motivation": []string{"text", "required"},
	}

	return m.Validate(rules, motivation)
}

func addCustomRules() {
	govalidator.AddCustomRule("text", func(field string, rule string, message string, value interface{}) error {
		switch value.(type) {
		case string:
			return nil
		}

		return fmt.Errorf("the field %s can only be a string", field)
	})
}
