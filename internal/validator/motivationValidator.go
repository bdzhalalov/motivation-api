package validator

import (
	"github.com/thedevsaddam/govalidator"
	"motivations-api/pkg/modules/motivations"
	"net/url"
)

//TODO: create a validator object when initializing the app for reuse

func ValidateCreatedMotivation(motivation *motivations.Motivation) url.Values {
	rules := govalidator.MapData{
		//TODO: create custom rule that supports cyrillic symbols
		"nickname":   []string{"required", "alpha_space"},
		"motivation": []string{"required", "alpha_space"},
	}

	opts := govalidator.Options{
		Data:  motivation,
		Rules: rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()

	if len(e) > 0 {
		return e
	}
	return nil
}

func ValidateUpdatedMotivation(motivation *motivations.Motivation) url.Values {
	rules := govalidator.MapData{
		//TODO: create custom rule that supports cyrillic symbols
		"motivation": []string{"required", "alpha_space"},
	}

	opts := govalidator.Options{
		Data:  motivation,
		Rules: rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()

	if len(e) > 0 {
		return e
	}
	return nil
}
