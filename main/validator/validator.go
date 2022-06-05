package main

import (
	"fmt"
	"goLearningProject/main/validator/comparisons/num_comparison"
	constants "goLearningProject/main/validator/const"
	"reflect"
	"sync"
)

type Student struct {
	Age   int    `validate:"gt=10"`
	Hobby string `validate:"eq=basketball"`
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
		Age:   10,
		Hobby: "football",
	}

	validator := NewEmptyValidator()
	validate, err := validator.validate(student)
		if err != nil {
		fmt.Println("errors:" + err.Error())
	}
	fmt.Println(validate)
}
