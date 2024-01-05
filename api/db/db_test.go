package db

import (
	"testing"

	"github.com/OffchainLabs/bold/api/server"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestGetEdges(t *testing.T) {
	sqlDb := setupTestDB(t)
	defer sqlDb.Close()
	db := &Database{sqlDB: sqlDb}

	e := &server.JsonEdge{}
	// err here is not nil because there are no field destinations for columns in `place`
	err := sqlDb.Get(e, "SELECT * FROM Edges LIMIT 1;")
	require.NoError(t, err)
	t.Logf("%+v\n", e)

	edges, err := db.GetEdges(WithLimit(2), WithOrderBy("CreatedAtBlock DESC"))
	require.NoError(t, err)
	require.Len(t, edges, 1) // Adjust according to your fake data
	t.Logf("%+v\n", edges)
	t.Fatal("Error")
}

func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	require.NoError(t, err)

	// Execute schema creation
	_, err = db.Exec(schema)
	require.NoError(t, err, "failed here")

	// Populate the database with fake data
	populateFakeData(t, db)

	return db
}

func populateFakeData(t *testing.T, db *sqlx.DB) error {
	// Insert data into Challenges
	challengeData := []string{
		"assertion_hash_1",
		"assertion_hash_2",
		"assertion_hash_3",
	}
	for _, hash := range challengeData {
		if _, err := db.Exec("INSERT INTO Challenges (AssertionHash) VALUES (?)", hash); err != nil {
			return err
		}
	}

	// Insert data into Edges
	edgeData := []server.JsonEdge{
		{
			Id:                common.BytesToHash([]byte("foobar")),
			ChallengeLevel:    0,
			OriginId:          common.BytesToHash([]byte("origin")),
			AssertionHash:     common.BytesToHash([]byte("assertion_hash")),
			StartHistoryRoot:  common.BytesToHash([]byte("start")),
			StartHeight:       0,
			EndHistoryRoot:    common.BytesToHash([]byte("end")),
			EndHeight:         32,
			MutualId:          common.BytesToHash([]byte("mutual")),
			ClaimId:           common.BytesToHash([]byte("claim")),
			HasChildren:       false,
			LowerChildId:      common.Hash{},
			UpperChildId:      common.Hash{},
			MiniStaker:        common.Address{},
			HasRival:          true,
			Status:            "confirmed",
			HasLengthOneRival: true,
			CreatedAtBlock:    230239,
		},
	}
	for _, e := range edgeData {
		_, err := db.NamedExec("INSERT INTO Edges (Id, ChallengeLevel, OriginId, AssertionHash, StartHistoryRoot, StartHeight, EndHistoryRoot, EndHeight, MutualId, ClaimId, HasChildren, LowerChildId, UpperChildId, MiniStaker, HasRival, Status, HasLengthOneRival, CreatedAtBlock) VALUES (:Id, :ChallengeLevel, :OriginId, :AssertionHash, :StartHistoryRoot, :StartHeight, :EndHistoryRoot, :EndHeight, :MutualId, :ClaimId, :HasChildren, :LowerChildId, :UpperChildId, :MiniStaker, :HasRival, :Status, :HasLengthOneRival, :CreatedAtBlock)", e)
		require.NoError(t, err)
	}

	// // Insert data into Assertions
	// assertionData := []struct {
	// 	Hash                string
	// 	ConfirmPeriodBlocks int
	// 	// Add other fields as necessary
	// }{
	// 	{"assertion_hash_1", 100},
	// 	{"assertion_hash_2", 200},
	// 	// Add more records as necessary
	// }
	// for _, a := range assertionData {
	// 	if _, err := db.Exec("INSERT INTO Assertions (Hash, ConfirmPeriodBlocks) VALUES (?, ?)",
	// 		a.Hash, a.ConfirmPeriodBlocks); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

// func TestGetAllChildren(t *testing.T) {
// 	db := setupTestDB(t)
// 	defer db.Close()

// 	// Assuming there are children for the provided ID in your fake data
// 	children, err := GetAllChildren(db, "parent_id_example")
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, children) // Check if children are returned
// }

// func TestGetEdgeWithAssertion(t *testing.T) {
// 	db := setupTestDB(t)
// 	defer db.Close()

// 	edgeWithAssertion, err := GetEdgeWithAssertion(db, "edge_id_example")
// 	assert.NoError(t, err)
// 	assert.NotNil(t, edgeWithAssertion) // Check if the result is not nil
// }
