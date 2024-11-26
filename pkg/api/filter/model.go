package filter

import "fmt"

const (
	DataTypeStr  = "string"
	DataTypeDate = "date"

	OperatorEq            = "="
	OperatorNotEq         = "neq"
	OperatorLowerThan     = "lt"
	OperatorLowerThanEq   = "lte"
	OperatorGreaterThat   = "gt"
	OperatorGreaterThatEq = "gte"
	OperatorBetween       = "between"
	OperatorLike          = "like"
)

type options struct {
	limit  int
	offset int
	fields []Field
}

func NewOptions(limit int, offset int) Options {
	return &options{
		limit:  limit,
		offset: offset,
	}
}

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}

type Options interface {
	Limit() int
	Offset() int
	AddField(name, operator, value, dtype string) error
	Fields() []Field
}

func (o *options) Limit() int {
	return o.limit
}

func (o *options) Offset() int {
	return o.offset
}

func (o *options) AddField(name, operator, value, dtype string) error {
	if err := validateOperator(operator); err != nil {
		return err
	}
	o.fields = append(o.fields, Field{
		Name:     name,
		Value:    value,
		Operator: operator,
		Type:     dtype,
	})
	return nil
}

func (o *options) Fields() []Field {
	return o.fields
}

func validateOperator(operator string) error {
	switch operator {
	case OperatorEq:
	case OperatorNotEq:
	case OperatorLowerThan:
	case OperatorLowerThanEq:
	case OperatorGreaterThat:
	case OperatorGreaterThatEq:
	case OperatorBetween:
	case OperatorLike:
	default:
		return fmt.Errorf("wrong operator")
	}
	return nil
}
