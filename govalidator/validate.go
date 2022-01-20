package govalidator

import "github.com/go-playground/validator/v10"

func Struct(st interface{}) error {
	if err := cacheValidate.Struct(st); err != nil {
		if er, ok := err.(*validator.InvalidValidationError); ok {
			return er
		}
		errs := ValidateStructErrors{}
		for _, er := range err.(validator.ValidationErrors) {
			e := InvalidValidateStruct{
				Namespace:   er.Namespace(),
				Field:       er.Field(),
				StructField: er.StructField(),
				Tag:         er.Tag(),
				ActualTag:   er.ActualTag(),
				Param:       er.Param(),
				Value:       er.Value(),
				Type:        er.Type(),
				Kind:        er.Kind(),
			}
			e.MsgError = e.Error()
			errs = append(errs, e)
		}
		return errs
	}
	return nil
}

func StructWithValidate(v *validator.Validate, st interface{}) error {
	if err := v.Struct(st); err != nil {
		if er, ok := err.(*validator.InvalidValidationError); ok {
			return er
		}
		errs := ValidateStructErrors{}
		for _, er := range err.(validator.ValidationErrors) {
			e := InvalidValidateStruct{
				Field:       er.Field(),
				StructField: er.StructField(),
				Tag:         er.Tag(),
				ActualTag:   er.ActualTag(),
				Param:       er.Param(),
				Value:       er.Value(),
				Type:        er.Type(),
				Kind:        er.Kind(),
			}
			errs = append(errs, e)
		}
		return errs
	}
	return nil
}
