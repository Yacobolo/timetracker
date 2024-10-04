package table

import (
	"fmt"
	"reflect"
)

type Table struct {
	Columns []string
	Rows    []Row
}

type Row struct {
	Values []string
}

func NewTableFromStructs(data interface{}) (Table, error) {
	val := reflect.ValueOf(data)

	// Check if the input is a slice
	if val.Kind() != reflect.Slice {
		return Table{}, fmt.Errorf("invalid input: expected a slice, got %s", val.Kind())
	}

	// Check if the slice is empty
	if val.Len() == 0 {
		return Table{}, nil
	}

	// Check if the slice elements are structs
	elemType := val.Type().Elem() // Store element type for reuse
	if elemType.Kind() != reflect.Struct {
		return Table{}, fmt.Errorf("invalid input: expected a slice of structs, got slice of %s", elemType.Kind())
	}

	numFields := elemType.NumField() // Store number of fields for reuse

	// Initialize TableData with pre-allocated capacity for Rows
	table := Table{
		Columns: make([]string, numFields),
		Rows:    make([]Row, 0, val.Len()),
	}

	// Extract column names from struct field names
	for i := 0; i < numFields; i++ {
		table.Columns[i] = elemType.Field(i).Name
	}

	// Extract row values from each struct
	for i := 0; i < val.Len(); i++ {
		row := Row{
			Values: make([]string, numFields),
		}

		for j := 0; j < numFields; j++ {
			field := val.Index(i).Field(j)
			row.Values[j] = fmt.Sprintf("%v", field.Interface()) // Handle any field type
		}
		table.Rows = append(table.Rows, row)
	}

	return table, nil
}
