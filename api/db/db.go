// Package db handles the interface to an underlying database of BOLD data
// for easy querying of information used by the BOLD API.
package db

import (
	"os"
	"strings"

	"github.com/OffchainLabs/bold/api/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	sqlDB               *sqlx.DB
	currentTableVersion int
}

func NewDatabase(path string) (*Database, error) {
	if _, err := os.Stat(path); err != nil {
		_, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	}
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &Database{
		sqlDB:               db,
		currentTableVersion: -1,
	}, nil
}

type EdgeQuery struct {
	filters []string
	args    []interface{}
	limit   int
	offset  int
	orderBy string
}

func NewEdgeQuery(opts ...EdgeOption) *EdgeQuery {
	query := &EdgeQuery{}
	for _, opt := range opts {
		opt(query)
	}
	return query
}

// Define similar function for Assertions

type EdgeOption func(e *EdgeQuery)

// EdgeOptions
func WithOriginID(originID string) EdgeOption {
	return func(q *EdgeQuery) {
		q.filters = append(q.filters, "OriginId = ?")
		q.args = append(q.args, originID)
	}
}

// Limit option
func WithLimit(limit int) EdgeOption {
	return func(q *EdgeQuery) {
		q.limit = limit
	}
}

// Offset option
func WithOffset(offset int) EdgeOption {
	return func(q *EdgeQuery) {
		q.offset = offset
	}
}

// OrderBy option
func WithOrderBy(orderBy string) EdgeOption {
	return func(q *EdgeQuery) {
		q.orderBy = orderBy
	}
}

func (q *EdgeQuery) ToSQL() (string, []interface{}) {
	baseQuery := "SELECT * FROM Edges"
	if len(q.filters) > 0 {
		baseQuery += " WHERE " + strings.Join(q.filters, " AND ")
	}
	if q.orderBy != "" {
		baseQuery += " ORDER BY " + q.orderBy
	}
	if q.limit > 0 {
		baseQuery += " LIMIT ?"
		q.args = append(q.args, q.limit)
	}
	if q.offset > 0 {
		baseQuery += " OFFSET ?"
		q.args = append(q.args, q.offset)
	}
	return baseQuery, q.args
}

// Define similar ToSQL method for Assertions

func GetEdges(db *sqlx.DB, opts ...EdgeOption) ([]*server.JsonEdge, error) {
	query := NewEdgeQuery(opts...)
	sql, args := query.ToSQL()
	edges := make([]*server.JsonEdge, 0)
	err := db.Select(&edges, sql, args...)
	if err != nil {
		return nil, err
	}
	return edges, nil
}

// func GetAllChildren(db *sqlx.DB, parentID string) ([]Edge, error) {
// 	var allChildren []Edge
// 	err := getChildrenRecursive(db, parentID, &allChildren)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return allChildren, nil
// }

// func getChildrenRecursive(db *sqlx.DB, parentID string, allChildren *[]Edge) error {
// 	var children []Edge
// 	query := `SELECT * FROM Edges WHERE LowerChildID = ? OR UpperChildID = ?`
// 	err := db.Select(&children, query, parentID, parentID)
// 	if err != nil {
// 		return err
// 	}

// 	for _, child := range children {
// 		*allChildren = append(*allChildren, child)
// 		err := getChildrenRecursive(db, child.ID, allChildren)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// type EdgeWithAssertion struct {
// 	Edge
// 	Assertion
// }

// func GetEdgeWithAssertion(db *sqlx.DB, edgeID string) (*EdgeWithAssertion, error) {
// 	var edgeWithAssertion EdgeWithAssertion
// 	query := `SELECT e.*, a.* FROM Edges e
//               JOIN Assertions a ON e.AssertionHash = a.Hash
//               WHERE e.ID = ?`

// 	err := db.Get(&edgeWithAssertion, query, edgeID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &edgeWithAssertion, nil
// }
