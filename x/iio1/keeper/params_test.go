package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iio1/testutil/keeper"
	"github.com/mytestlab123/iio1/x/iio1/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Iio1Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
