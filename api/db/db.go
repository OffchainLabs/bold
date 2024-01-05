// Package db handles the interface to an underlying database of BOLD data
// for easy querying of information used by the BOLD API.
package db

import (
	"os"
	"strings"

	"github.com/OffchainLabs/bold/api/server"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
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

type AssertionQuery struct {
	filters []string
	args    []interface{}
	limit   int
	offset  int
	orderBy string
}

func NewAssertionQuery(opts ...AssertionOption) *AssertionQuery {
	query := &AssertionQuery{}
	for _, opt := range opts {
		opt(query)
	}
	return query
}

type AssertionOption func(*AssertionQuery)

// Options for Assertions (similar to EdgeOptions)
func WithAssertionHash(hash string) AssertionOption {
	return func(q *AssertionQuery) {
		q.filters = append(q.filters, "Hash = ?")
		q.args = append(q.args, hash)
	}
}

func (q *AssertionQuery) ToSQL() (string, []interface{}) {
	baseQuery := "SELECT * FROM Assertions"
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

func (d *Database) GetAssertions(opts ...AssertionOption) ([]*server.JsonAssertion, error) {
	query := NewAssertionQuery(opts...)
	sql, args := query.ToSQL()
	assertions := make([]*server.JsonAssertion, 0)
	err := d.sqlDB.Select(&assertions, sql, args...)
	if err != nil {
		return nil, err
	}
	return assertions, nil
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

func (d *Database) GetEdges(opts ...EdgeOption) ([]*server.JsonEdge, error) {
	query := NewEdgeQuery(opts...)
	sql, args := query.ToSQL()
	edges := make([]*server.JsonEdge, 0)
	err := d.sqlDB.Select(&edges, sql, args...)
	if err != nil {
		return nil, err
	}
	return edges, nil
}

func (d *Database) GetAllChildren(edgeId common.Hash) ([]*server.JsonEdge, error) {
	var allChildren []*server.JsonEdge
	err := d.getChildrenRecursive(edgeId, allChildren)
	if err != nil {
		return nil, err
	}
	return allChildren, nil
}

func (d *Database) getChildrenRecursive(parentID common.Hash, allChildren []*server.JsonEdge) error {
	var children []*server.JsonEdge
	query := `SELECT * FROM Edges WHERE LowerChildID = ? OR UpperChildID = ?`
	err := d.sqlDB.Select(&children, query, parentID, parentID)
	if err != nil {
		return err
	}

	for _, child := range children {
		allChildren = append(allChildren, child)
		err := d.getChildrenRecursive(child.Id, allChildren)
		if err != nil {
			return err
		}
	}
	return nil
}

type AssertionWithInfo struct {
	protocol.AssertionCreatedInfo
}

func (d *Database) InsertAssertions(assertions []*AssertionWithInfo) error {
	for _, a := range assertions {
		if err := d.InsertAssertion(a); err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) InsertAssertion(a *AssertionWithInfo) error {
	return nil
}

func (d *Database) InsertEdges(edges []protocol.SpecEdge) error {
	for _, e := range edges {
		if err := d.InsertEdge(e); err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) InsertEdge(edge protocol.SpecEdge) error {
	return nil
}
