package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://ibrahim:ibrahim@localhost:5432/simple_bank?sslmode=disable"
// )

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}
