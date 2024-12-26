package repo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParams_GetWhere(t *testing.T) {
	r := require.New(t)

	testCases := map[string]struct {
		params Params
		exp    string
	}{
		"single": {
			params: Params{
				{
					Field:    "name",
					Operator: Equal,
					Value:    "John",
				},
			},
			exp: "name = :param_name",
		},
		"multiple": {
			params: Params{
				{
					Field:    "name",
					Operator: Equal,
					Value:    "John",
				},
				{
					Field:    "age",
					Operator: Equal,
					Value:    40,
				},
			},
			exp: "name = :param_name AND age = :param_age",
		},
		"operator - not equal": {
			params: Params{
				{
					Field:    "name",
					Operator: NotEqual,
					Value:    "John",
				},
			},
			exp: "name <> :param_name",
		},
		"operator - less than": {
			params: Params{
				{
					Field:    "age",
					Operator: LessThan,
					Value:    40,
				},
			},
			exp: "age < :param_age",
		},
		"operator - less than equal to": {
			params: Params{
				{
					Field:    "age",
					Operator: LessThanEqual,
					Value:    40,
				},
			},
			exp: "age <= :param_age",
		},
		"operator - greater than": {
			params: Params{
				{
					Field:    "age",
					Operator: GreaterThan,
					Value:    40,
				},
			},
			exp: "age > :param_age",
		},
		"operator - greater than equal to": {
			params: Params{
				{
					Field:    "age",
					Operator: GreaterThanEqual,
					Value:    40,
				},
			},
			exp: "age >= :param_age",
		},
		"default operator": {
			params: Params{
				{
					Field: "name",
					Value: "John",
				},
			},
			exp: "name = :param_name",
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			r.Equal(tc.exp, tc.params.GetWhere())
		})
	}
}

func TestParam_GetValueMap(t *testing.T) {
	r := require.New(t)

	testCases := map[string]struct {
		params Params
		exp    map[string]any
	}{
		"single": {
			params: Params{
				{Field: "name", Operator: Equal, Value: "John"},
			},
			exp: map[string]any{
				"param_name": "John",
			},
		},
		"multiple": {
			params: Params{
				{Field: "name", Operator: Equal, Value: "John"},
				{Field: "age", Operator: Equal, Value: 40},
			},
			exp: map[string]any{
				"param_name": "John",
				"param_age":  40,
			},
		},
		"non-equal operator": {
			params: Params{
				{Field: "name", Operator: NotEqual, Value: "John"},
			},
			exp: map[string]any{
				"param_name": "John",
			},
		},
		"default operator": {
			params: Params{
				{Field: "name", Value: "John"},
			},
			exp: map[string]any{
				"param_name": "John",
			},
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			r.Equal(tc.exp, tc.params.GetValueMap())
		})
	}
}
