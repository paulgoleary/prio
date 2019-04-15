package mpc

import (
	"github.com/henrycg/prio/config"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func getSimpleField() config.Field {
	return config.Field{Name:"F1", Type:config.TypeInt, IntBits:63}
}

func TestBasic(t *testing.T) {
	cfg := new(config.Config)
	cfg.Servers = []config.ServerAddress{{"PUB1", "PRIV1"}, {"PUB2", "PRIV2"}}

	cfg.Fields = make([]config.Field, 1)
	cfg.Fields[0] = getSimpleField()

	crs := RandomRequest(cfg, 0)
	require.Equal(t, len(crs), 2)
	require.Equal(t, len(cfg.Circuit.Outputs()), 1)

	checkers := make([]*Checker, 2)
	checkOut := new(big.Int)
	for x, _ := range checkers {
		checkers[x] = NewChecker(cfg, x, 0)
		checkers[x].SetReq(crs[x])

		require.Equal(t, len(checkers[x].Outputs()), 1)
		checkOut.Add(checkOut, checkers[x].Outputs()[0].WireValue)
	}
	checkOut.Mod(checkOut, checkers[0].mod)
	require.True(t, cfg.Circuit.Outputs()[0].WireValue.Cmp(checkOut) == 0, "additive shares should combine to original output value")
}
