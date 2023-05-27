package sched

import (
	"testing"
	"zakirullin/dumpbot/pkg/str"

	"github.com/stretchr/testify/require"
)

func TestUcfirst(t *testing.T) {
	r := require.New(t)

	res := str.Ucfirst("abc")

	r.Equal("Abc", res)
}

func TestUcfirstRu(t *testing.T) {
	r := require.New(t)

	res := str.Ucfirst("абв")

	r.Equal("Абв", res)
}
