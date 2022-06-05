package defaultcomparison

import constants "goLearningProject/main/validator/const"

type DefaultComparator interface {
	Compare(tagString string, obj constants.Object) (bool, error)
}
