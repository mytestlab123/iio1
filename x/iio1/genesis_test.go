package iio1_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iio1/testutil/keeper"
	"github.com/mytestlab123/iio1/testutil/nullify"
	"github.com/mytestlab123/iio1/x/iio1"
	"github.com/mytestlab123/iio1/x/iio1/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Iio1Keeper(t)
	iio1.InitGenesis(ctx, *k, genesisState)
	got := iio1.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
