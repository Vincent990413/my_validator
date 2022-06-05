package num_comparison

import (
	"errors"
	defaultcomparison "goLearningProject/main/validator/comparisons"
	constants "goLearningProject/main/validator/const"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type NumComparison struct {
	defaultComparison defaultcomparison.DefaultComparison
	myLock            sync.Mutex
}

type NumComparator interface {
	defaultcomparison.DefaultComparator
}

func NewEmptyComparison() *NumComparison {
	return &NumComparison{}
}

func (numC *NumComparison) Compare(tagString string, obj constants.Object) (bool, error) {
	numC.myLock.Lock()
	defer numC.myLock.Unlock()

	objectType := reflect.TypeOf(obj)
	objectValues := reflect.ValueOf(obj)

	for i := 0; i < objectValues.NumField(); i++ {
		fieldValue := objectValues.Field(i)
		tagContent := objectType.Field(i).Tag.Get("validate")
		fieldKind := fieldValue.Kind()

		//确保只有一个验证条件
		//TODO 这里可以想办法加逗号，类似required, gt=20这样的多重条件
		switch fieldKind {
		case reflect.Int:
			splitStringArr := strings.Split(tagContent, constants.Equal)
			//eq = 10 || gt = 12
			operatorDesc := splitStringArr[0]
			operatorNum := splitStringArr[1]
			valInt := fieldValue.Int()                                                //valInt 为实际的值
			parseInt, err := strconv.ParseInt(operatorNum, constants.DecimalType, 64) //parseInt 为标签中的值
			if err != nil {
				return false, err
			}

			switch operatorDesc {
			case constants.ComparisonTagEq:
				return compareEqInternal(parseInt, valInt)

			case constants.ComparisonTagGt:
				return compareGtInternal(parseInt, valInt)

			case constants.ComparisonTagLt:
				return compareLtInternal(parseInt, valInt)

			case constants.ComparisonTagNe:
				return compareNeInternal(parseInt, valInt)
			}

		}
	}

	return true, nil
}

func compareEqInternal(parseInt int64, valInt int64) (bool, error) {
	if valInt != parseInt {
		errMsg := "validate int failed, tag's value is not " + strconv.FormatInt(parseInt, 10) + " but " + strconv.FormatInt(valInt, 10)
		return false, errors.New(errMsg)
	}
	return true, nil
}

func compareGtInternal(parseInt int64, valInt int64) (bool, error) {
	if valInt <= parseInt {
		errMsg := "validate int failed, tag's value is " + strconv.FormatInt(parseInt, 10) + " but not greater than " + strconv.FormatInt(valInt, 10)
		return false, errors.New(errMsg)
	}
	return true, nil
}

func compareLtInternal(parseInt int64, valInt int64) (bool, error) {
	if valInt >= parseInt {
		errMsg := "validate int failed, tag's value is " + strconv.FormatInt(parseInt, 10) + " but not less than " + strconv.FormatInt(valInt, 10)
		return false, errors.New(errMsg)
	}
	return true, nil
}

func compareNeInternal(parseInt int64, valInt int64) (bool, error) {
	if valInt == parseInt {
		errMsg := "validate int failed, tag's value is " + strconv.FormatInt(parseInt, 10) + " but equal to " + strconv.FormatInt(valInt, 10)
		return false, errors.New(errMsg)
	}
	return true, nil
}

