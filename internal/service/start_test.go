package service_test

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"

	"go-framework/pkg/boot"
)

var booted *boot.Booted
var server *gin.Engine

func TestMain(m *testing.M) {
	var err error
	if booted, err = boot.Boot(
		boot.WithConfigFile(os.Getenv("LGO_TEST_FILE")),
	); err != nil {
		panic(err)
	}
	server = booted.Server
	m.Run()
}
