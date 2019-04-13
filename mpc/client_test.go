package mpc

import (
	"github.com/henrycg/prio/config"
	"github.com/stretchr/testify/require"
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
}
