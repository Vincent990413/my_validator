package constants

type Object interface{}

const Equal = "="

const DecimalType = 10

//以下是各个比较中的标签描述 比如"等于"就是"eq"
const (
	ComparisonTagEq string = "eq"
	ComparisonTagGt string = "gt"
	ComparisonTagLt string = "lt"
	ComparisonTagNe string = "ne"
)
