package main

import (
	"errors"
	"fmt"
	"goLearningProject/main/validator/comparisons/num_comparison"
	"goLearningProject/main/validator/const"
	"reflect"
	"strings"
	"sync"
)

type Student struct {
	Age    int    `validate:"gt=10"`
	Hobby  string `validate:"eq=basketball"`
	Gender string `validate:"required"`
}

type Validator struct {
	objectValidate constants.Object
	myLock         sync.Mutex
}

func NewEmptyValidator() *Validator {
	return &Validator{}
}

func (v *Validator) validate(obj interface{}) (bool, error) {
	v.myLock.Lock()
	defer v.myLock.Unlock()

	objectType := reflect.TypeOf(obj)
	objectValues := reflect.ValueOf(obj)

	for i := 0; i < objectValues.NumField(); i++ {
		fieldValue := objectValues.Field(i)
		tagContent := objectType.Field(i).Tag.Get("validate")
		fieldKind := fieldValue.Kind()

		if strings.Contains(tagContent, "required") {
			fieldRequired := objectValues.FieldByName(objectType.Field(i).Name)
			if strings.Compare(fieldRequired.String(), "") == 0 {
				return false, errors.New("validate failed, the field " + objectType.Field(i).Name + " must be required!")
			}
		}

		if tagContent == "nil" {
			continue
		}

		switch fieldKind {
		case reflect.Int: //如果当前是int类型的域
			comparison := num_comparison.NewEmptyComparison()
			eq, err := comparison.Compare(tagContent, obj)
			if !eq || err != nil {
				return false, err
			}

		}

	}
	return true, nil
}

func main() {
	var student = Student{
		Age:   11,
		Hobby: "football",
	}

	validator := NewEmptyValidator()
	validate, err := validator.validate(student)
	if err != nil {
		fmt.Println("errors:" + err.Error())
	}
	fmt.Println(validate)
}
