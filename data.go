package repo

import (
	"fmt"
	"reflect"
	"strings"
)

type Data struct {
	data any
	t    reflect.Type
	v    reflect.Value
}

// NewData creates a data with helper functions to generate SQL statements for the specified data type
func NewData(data any) (*Data, error) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	// type should be a struct
	kind := t.Kind()
	if kind != reflect.Struct {
		return nil, fmt.Errorf("expected struct but got %s", kind)
	}

	return &Data{
		data: data,
		t:    t,
		v:    v,
	}, nil
}

// GetFields returns a slice of strings with the fields of the struct having a `db` struct tag
func (d *Data) GetFields() []string {
	var fields []string
	for i := range d.t.NumField() {
		dbField := d.t.Field(i).Tag.Get("db")
		// skip fields that don't have `db` struct tag
		if dbField == "" {
			continue
		}
		fields = append(fields, dbField)
	}

	return fields
}

// GetFieldList returns a string with the concatenation of the fields separated
// by comma. Used in INSERT to list the fields
func (d *Data) GetFieldList() string {
	return strings.Join(d.GetFields(), ", ")
}

// GetParams is similar to GetFields but prepends a colon to the field name (e.g.
// :field1)
func (d *Data) GetParams() []string {
	var params []string
	for _, field := range d.GetFields() {
		params = append(params, fmt.Sprintf(":%s", field))
	}

	return params
}

// GetParamsList is similar to GetFieldList but prepends a colon to the field
// names (e.g. :field1, :field2). Used in INSERT to specify the values
func (d *Data) GetParamsList() string {
	return strings.Join(d.GetParams(), ", ")
}

// GetValueMap returns a map of field and its corresponding value. (e.g.
// map[string]any{ "name": "John" })
func (d *Data) GetValueMap() map[string]any {
	valueMap := map[string]any{}

	for i := range d.t.NumField() {
		dbField := d.t.Field(i).Tag.Get("db")
		dbValue := d.v.Field(i).Interface()
		// skip fields that don't have `db` struct tag
		if dbField == "" {
			continue
		}
		valueMap[dbField] = dbValue
	}

	return valueMap
}

// GetFieldParamList returns a string with the concatenation of fields and params
// separated by comma. (e.g. field1 = :field1). Used in UPDATE
func (d *Data) GetFieldParamList() string {
	var fieldParams []string
	for _, field := range d.GetFields() {
		fieldParam := fmt.Sprintf("%s = :%s", field, field)
		fieldParams = append(fieldParams, fieldParam)
	}

	return strings.Join(fieldParams, ", ")
}
