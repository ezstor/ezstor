package info_test

import (
	"encoding/json"
	"os"
	"sort"
	"testing"

	"github.com/ezstor/ezstor/pkg/info"
	"github.com/stretchr/testify/require"
)

type TestData struct {
	Controllers []string
}

func TestControllers(t *testing.T) {
	fileBytes, err := os.ReadFile("./testdata/testdata.json")
	require.Nil(t, err)

	testData := TestData{}
	err = json.Unmarshal(fileBytes, &testData)
	require.Nil(t, err)

	conts, err := info.Controllers()
	require.Nil(t, err)

	sort.Strings(testData.Controllers)
	sort.Strings(conts)

	require.Equal(t, testData.Controllers, conts)
}
