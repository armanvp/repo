package repo

import "fmt"

const (
	Equal            = "="
	NotEqual         = "<>"
	LessThan         = "<"
	LessThanEqual    = "<="
	GreaterThan      = ">"
	GreaterThanEqual = ">="
	In               = "IN" // TODO implement later
)

type Param struct {
	Field    string
	Operator string
	Value    any
}

type Params []Param

// GetWhere returns the where string of a params (e.g. field1 = :param_field1 AND field2 = :param_field2)
func (p Params) GetWhere() string {
	var (
		where    string
		joiner   string
		operator string
	)

	for i, param := range p {
		if i > 0 {
			joiner = " AND "
		}

		operator = param.Operator
		// set default operator
		if param.Operator == "" {
			operator = Equal
		}

		where = fmt.Sprintf("%s%s%s %s :param_%s", where, joiner, param.Field, operator, param.Field)
	}

	return where
}

// GetValueMap returns a map of field and its corresponding value. (e.g.
// map[string]any{ "param_name": "John" })
func (p Params) GetValueMap() map[string]any {
	values := make(map[string]any, len(p))
	for _, param := range p {
		values[fmt.Sprintf("param_%s", param.Field)] = param.Value
	}

	return values
}
