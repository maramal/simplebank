package api

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/maramal/simplebank/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

type Obj interface {
	db.Account | db.User
}

func requireBodyMatch[T Obj](t *testing.T, body *bytes.Buffer, obj T) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotObj T
	err = json.Unmarshal(data, &gotObj)
	require.NoError(t, err)
	require.Equal(t, obj, gotObj)
}
