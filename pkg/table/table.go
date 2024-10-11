package table

import (
	"fmt"
	"reflect"
)

type Table struct {
	Columns []Column
	Rows    []Row
}

type Column struct {
	ID    string
	Label string
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
		Columns: make([]Column, numFields),
		Rows:    make([]Row, 0, val.Len()),
	}

	// Extract column names from struct field tags or names
	for i := 0; i < numFields; i++ {
		field := elemType.Field(i)
		dbTag := field.Tag.Get("db")

		// Use the db tag as the column ID if available, otherwise use the field name
		columnID := field.Name
		if dbTag != "" {
			columnID = dbTag
		}

		column := Column{
			ID:    columnID,
			Label: field.Name, // You can change how the label is set if needed
		}

		table.Columns[i] = column
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

func NewRowFromStruct(data interface{}) (Row, error) {
	val := reflect.ValueOf(data)

	// Check if the input is a struct
	if val.Kind() != reflect.Struct {
		return Row{}, fmt.Errorf("invalid input: expected a struct, got %s", val.Kind())
	}

	numFields := val.NumField() // Store number of fields for reuse

	// Initialize RowData with pre-allocated capacity for Values
	row := Row{
		Values: make([]string, numFields),
	}

	// Extract values from struct fields
	for i := 0; i < numFields; i++ {
		field := val.Field(i)
		row.Values[i] = fmt.Sprintf("%v", field.Interface()) // Handle any field type
	}

	return row, nil
}
