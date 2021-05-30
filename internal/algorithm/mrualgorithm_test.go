package MRUApp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeTestInput(t *testing.T) {
	mru := MakeTestInput()
	require.NotNil(t, mru)
	fmt.Printf("\n%#v\n", mru)
	require.True(t, false)
}
