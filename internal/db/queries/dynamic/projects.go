package queries

import (
	"github.com/Masterminds/squirrel"
)

// BuildBaseProjectQuery creates the base query for project CRUD operations.
func BuildBaseProjectQuery() squirrel.SelectBuilder {
	return squirrel.Select("id", "name", "description", "created_at").From("projects")
}

type ProjectListQueryOpts struct {
	SortBy    string
	SortOrder string
	Limit     uint64
	Offset    uint64
}

func BuildProjectListQuery(ops ProjectListQueryOpts) squirrel.SelectBuilder {
	query := BuildBaseProjectQuery()

	query = ApplySorting(query, []string{ops.SortBy}, []string{ops.SortOrder})

	query = ApplyPagination(query, ops.Limit, ops.Offset)

	return query

}

// helper
func ApplySorting(query squirrel.SelectBuilder, sortBy []string, sortOrder []string) squirrel.SelectBuilder {

	for i, column := range sortBy {
		order := "ASC"
		if i < len(sortOrder) && sortOrder[i] == "desc" {
			order = "DESC"
		}
		query = query.OrderBy(column + " " + order)
	}
	return query
}

func ApplyPagination(query squirrel.SelectBuilder, limit, offset uint64) squirrel.SelectBuilder {
	if limit > 0 {
		query = query.Limit(limit)

		if offset > 0 {
			query = query.Offset(offset)
		}
	}

	return query
}
