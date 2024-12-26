package repo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecord_GetFields(t *testing.T) {
	r := require.New(t)

	testCases := map[string]struct {
		data any
		exp  []string
		err  error
	}{
		"struct": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
			}{},
			exp: []string{"field1", "field2"},
		},
		"struct with some db fields": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
				Field3 string
			}{},
			exp: []string{"field1", "field2"},
		},
		"non-struct": {
			data: "test",
			err:  fmt.Errorf("expected struct but got string"),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			rec, err := NewData(tc.data)
			if tc.err == nil {
				act := rec.GetFields()
				r.NoError(err, "GetFields() should not return an error")
				r.Equal(tc.exp, act, "should have the expected fields")
				return
			}

			// error cases
			r.ErrorContains(err, tc.err.Error(), "should have the expected error")
		})
	}
}

func TestRecord_GetFieldList(t *testing.T) {
	r := require.New(t)

	testCases := map[string]struct {
		data any
		exp  string
		err  error
	}{
		"struct": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
			}{},
			exp: "field1, field2",
		},
		"struct with some db fields": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
				Field3 string
			}{},
			exp: "field1, field2",
		},
		"non-struct": {
			data: "test",
			err:  fmt.Errorf("expected struct but got string"),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			rec, err := NewData(tc.data)
			if tc.err == nil {
				act := rec.GetFieldList()
				r.NoError(err, "GetFieldList() should not return an error")
				r.Equal(tc.exp, act, "should have the expected fields")
				return
			}

			// error cases
			r.ErrorContains(err, tc.err.Error(), "should have the expected error")
		})
	}
}

func TestRecord_GetParams(t *testing.T) {
	r := require.New(t)

	testCases := map[string]struct {
		data any
		exp  []string
		err  error
	}{
		"struct": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
			}{},
			exp: []string{":field1", ":field2"},
		},
		"struct with some db fields": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
				Field3 string
			}{},
			exp: []string{":field1", ":field2"},
		},
		"non-struct": {
			data: "test",
			err:  fmt.Errorf("expected struct but got string"),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			rec, err := NewData(tc.data)
			if tc.err == nil {
				act := rec.GetParams()
				r.NoError(err, "GetParams() should not return an error")
				r.Equal(tc.exp, act, "should have the expected fields")
				return
			}

			// error cases
			r.ErrorContains(err, tc.err.Error(), "should have the expected error")
		})
	}
}

func TestRecord_GetParamsList(t *testing.T) {
	r := require.New(t)

	testCases := map[string]struct {
		data any
		exp  string
		err  error
	}{
		"struct": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
			}{},
			exp: ":field1, :field2",
		},
		"struct with some db fields": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
				Field3 string
			}{},
			exp: ":field1, :field2",
		},
		"non-struct": {
			data: "test",
			err:  fmt.Errorf("expected struct but got string"),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			rec, err := NewData(tc.data)
			if tc.err == nil {
				act := rec.GetParamsList()
				r.NoError(err, "GetParamsList() should not return an error")
				r.Equal(tc.exp, act, "should have the expected fields")
				return
			}

			// error cases
			r.ErrorContains(err, tc.err.Error(), "should have the expected error")
		})
	}
}

func TestRecord_GetValueMap(t *testing.T) {
	r := require.New(t)

	testCases := map[string]struct {
		data any
		exp  map[string]any
		err  error
	}{
		"struct": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 int    `db:"field2"`
			}{
				Field1: "James",
				Field2: 40,
			},
			exp: map[string]any{
				"field1": "James",
				"field2": 40,
			},
		},
		"struct with some db fields": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 int    `db:"field2"`
				Field3 string
			}{
				Field1: "James",
				Field2: 40,
				Field3: "Doe",
			},
			exp: map[string]any{
				"field1": "James",
				"field2": 40,
			},
		},
		"non-struct": {
			data: "test",
			err:  fmt.Errorf("expected struct but got string"),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			rec, err := NewData(tc.data)
			if tc.err == nil {
				act := rec.GetValueMap()
				r.NoError(err, "GetValueMap() should not return an error")
				r.Equal(tc.exp, act, "should have the expected fields")
				return
			}

			// error cases
			r.ErrorContains(err, tc.err.Error(), "should have the expected error")
		})
	}
}

func TestRecord_GetFieldParamList(t *testing.T) {
	r := require.New(t)

	testCases := map[string]struct {
		data any
		exp  string
		err  error
	}{
		"struct": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
			}{},
			exp: "field1 = :field1, field2 = :field2",
		},
		"struct with some db fields": {
			data: struct {
				Field1 string `db:"field1"`
				Field2 string `db:"field2"`
				Field3 string
			}{},
			exp: "field1 = :field1, field2 = :field2",
		},
		"non-struct": {
			data: "test",
			err:  fmt.Errorf("expected struct but got string"),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			rec, err := NewData(tc.data)
			if tc.err == nil {
				act := rec.GetFieldParamList()
				r.NoError(err, "GetParamsList() should not return an error")
				r.Equal(tc.exp, act, "should have the expected fields")
				return
			}

			// error cases
			r.ErrorContains(err, tc.err.Error(), "should have the expected error")
		})
	}
}
