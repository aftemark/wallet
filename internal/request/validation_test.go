package request

import (
	"context"
	"fmt"
	"testing"
	"wallet/internal/postgres"

	"github.com/stretchr/testify/assert"
)

func TestValidateExists(t *testing.T) {
	db, err := postgres.NewTestConn()
	assert.NoError(t, err)
	defer db.Close(context.Background())

	table := "wallets"
	var okSndr, errSndr uint32 = 1, 6
	assert.NoError(t, validateExists(db, table)(okSndr))
	assert.EqualError(t, validateExists(db, table)(errSndr), "wallet not found")
}

func TestValidateDeposit(t *testing.T) {
	db, err := postgres.NewTestConn()
	assert.NoError(t, err)
	defer db.Close(context.Background())

	col := "amount"
	table := "wallets"
	var okRcvr, errRcvr uint32 = 3, 2
	var okAmnt, errAmnt float64 = 1, 0.01
	errTxt := fmt.Sprintf("too much money! Result amount must be no greater than 1e+%v", numericPrecision)
	assert.NoError(t, validateDeposit(db, table, col, okRcvr)(okAmnt))
	assert.EqualError(t, validateDeposit(db, table, col, errRcvr)(errAmnt), errTxt)
}

func TestValidateWithdrawal(t *testing.T) {
	db, err := postgres.NewTestConn()
	assert.NoError(t, err)
	defer db.Close(context.Background())

	col := "amount"
	table := "wallets"
	var sndr uint32 = 1
	var okAmnt, errAmnt float64 = 1, 1000
	assert.NoError(t, validateWithdrawal(db, table, col, sndr)(okAmnt))
	assert.EqualError(t, validateWithdrawal(db, table, col, sndr)(errAmnt), "insufficient funds")
}
