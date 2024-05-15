package ndfc

import (
	"log"
	"os"
	"testing"
)

var ut_client *NDFC

func TestMain(m *testing.M) {
	host := "http://localhost:3001"
	user := "admin"
	password := "admin"
	domain := "local"
	timeout := int64(1000)
	var err error
	log.Print("Running TestMain for NDFC package")
	ut_client, err = NewNDFCClient(host, user, password, domain, true, timeout)
	if err != nil {
		log.Fatal("Failed to create new NDFC client")
		return
	}
	os.Exit(m.Run())
}
