package utils

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	baseQuery  string
	conditions []string
	parameters []interface{}
	paramCount int
	sortBy     string
	limit      int
	offset     int
}

func NewQueryBuilder(baseQuery string) *QueryBuilder {
	return &QueryBuilder{
		baseQuery:  baseQuery,
		conditions: []string{},
		parameters: []interface{}{},
		paramCount: 1,
	}
}

func (qb *QueryBuilder) AddCondition(condition string, value interface{}) {
	// Add condition
	qb.conditions = append(qb.conditions, fmt.Sprintf(condition, qb.paramCount))
	// Add query value
	qb.parameters = append(qb.parameters, value)
	// Increment the param count
	qb.paramCount++
}

func (qb *QueryBuilder) SetLimit(limit int) {
	qb.limit = limit
}

func (qb *QueryBuilder) SetOffset(offset int) {
	qb.offset = offset
}

func (qb *QueryBuilder) SetSortBy(sortBy string) {
	qb.sortBy = sortBy
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
	// Get the base query
	query := qb.baseQuery

	// Add conditions
	if len(qb.conditions) > 0 {
		query += " WHERE " + strings.Join(qb.conditions, " AND ")
	}

	// Add sort by
	if qb.sortBy != "" {
		query += " ORDER BY " + qb.sortBy
	}

	// Add limit and offset
	if qb.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d ", qb.limit)
	}
	if qb.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d ", qb.offset)
	}

	return query, qb.parameters
}
