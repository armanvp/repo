package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repo[T any] struct {
	table string
	tx    *sqlx.Tx
}

// New returns a repo
func New[T any](table string, tx *sqlx.Tx) *Repo[T] {
	return &Repo[T]{
		table: table,
		tx:    tx,
	}
}

// Get returns a record by ID
func (r *Repo[T]) Get(id any) (*T, error) {

	sql := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", r.table)

	var row T
	err := r.tx.QueryRowx(sql, id).StructScan(&row)
	if err != nil {
		return nil, fmt.Errorf("failed to query row: %w", err)
	}

	return &row, nil
}

// GetByParam returns a slice of records by param
func (r *Repo[T]) GetByParam(params Params) ([]*T, error) {

	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s", r.table, params.GetWhere())

	rows, err := r.tx.NamedQuery(sql, params.GetValueMap())
	if err != nil {
		return nil, fmt.Errorf("failed to query rows: %w", err)
	}

	var data []*T
	for rows.Next() {
		var row T
		err := rows.StructScan(&row)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		data = append(data, &row)
	}

	return data, nil
}

// Create a record
func (r *Repo[T]) Create(data *T) error {
	rec, err := NewData(*data)
	if err != nil {
		return fmt.Errorf("failed new record: %w", err)
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", r.table, rec.GetFieldList(), rec.GetParamsList())

	_, err = r.tx.NamedExec(sql, rec.GetValueMap())
	if err != nil {
		return fmt.Errorf("failed to insert records: %w", err)
	}

	return nil
}

// Update a record by ID
func (r *Repo[T]) Update(id any, data *T) error {
	rec, err := NewData(*data)
	if err != nil {
		return fmt.Errorf("failed new record: %w", err)
	}

	param := Params{
		{Field: "id", Operator: Equal, Value: id},
	}

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", r.table, rec.GetFieldParamList(), param.GetWhere())

	// merge value map of the data struct and params
	valueMap := param.GetValueMap()
	for k, v := range rec.GetValueMap() {
		valueMap[k] = v
	}

	_, err = r.tx.NamedExec(sql, valueMap)
	if err != nil {
		return fmt.Errorf("failed to update records: %w", err)
	}

	return nil
}

// Delete a record by id
func (r *Repo[T]) Delete(id any) error {
	param := Params{
		{Field: "id", Operator: Equal, Value: id},
	}

	sql := fmt.Sprintf("DELETE FROM %s WHERE %s", r.table, param.GetWhere())

	_, err := r.tx.NamedExec(sql, param.GetValueMap())
	if err != nil {
		return fmt.Errorf("failed to delete records: %w", err)
	}

	return nil
}
