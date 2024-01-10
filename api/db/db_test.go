package db

import (
	"testing"

	"github.com/OffchainLabs/bold/api"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSqliteDatabase_Edges(t *testing.T) {
	sqlDB, err := sqlx.Connect("sqlite3", ":memory:")
	require.NoError(t, err)
	defer sqlDB.Close()

	_, err = sqlDB.Exec(schema)
	require.NoError(t, err)

	db := &SqliteDatabase{sqlDB: sqlDB}
	assertionsToCreate := []*api.JsonAssertion{
		baseAssertion(),
	}
	require.NoError(t, db.InsertAssertions(assertionsToCreate))
	edgesToCreate := []*api.JsonEdge{
		baseEdge(),
	}
	require.NoError(t, db.InsertEdges(edgesToCreate))

	edges, err := db.GetEdges(WithLimit(1))
	require.NoError(t, err)
	t.Log(edges)
}

func baseAssertion() *api.JsonAssertion {
	return &api.JsonAssertion{
		Hash:                     common.Hash{},
		ConfirmPeriodBlocks:      100,
		RequiredStake:            "1",
		ParentAssertionHash:      common.Hash{},
		InboxMaxCount:            "1",
		AfterInboxBatchAcc:       common.Hash{},
		WasmModuleRoot:           common.Hash{},
		ChallengeManager:         common.Address{},
		CreationBlock:            1,
		TransactionHash:          common.Hash{},
		BeforeStateBlockHash:     common.Hash{},
		BeforeStateSendRoot:      common.Hash{},
		BeforeStateMachineStatus: protocol.MachineStatusFinished,
		AfterStateBlockHash:      common.Hash{},
		AfterStateSendRoot:       common.Hash{},
		AfterStateMachineStatus:  protocol.MachineStatusFinished,
		FirstChildBlock:          nil,
		SecondChildBlock:         nil,
		IsFirstChild:             true,
		Status:                   protocol.AssertionPending,
		ConfigHash:               common.Hash{},
	}
}

func baseEdge() *api.JsonEdge {
	return &api.JsonEdge{
		Id:                common.Hash{},
		ChallengeLevel:    0,
		OriginId:          common.Hash{},
		AssertionHash:     common.Hash{},
		StartHistoryRoot:  common.Hash{},
		StartHeight:       0,
		EndHistoryRoot:    common.Hash{},
		EndHeight:         0,
		MutualId:          common.Hash{},
		ClaimId:           common.Hash{},
		HasChildren:       false,
		LowerChildId:      common.Hash{},
		UpperChildId:      common.Hash{},
		MiniStaker:        common.Address{},
		HasRival:          true,
		Status:            "pending",
		HasLengthOneRival: true,
		CreatedAtBlock:    1,
	}
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
