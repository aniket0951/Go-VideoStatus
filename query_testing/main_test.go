package querytesting

import (
	"database/sql"
	"log"
	"os"
	"testing"

	sqlc_lib "github.com/aniket0951/video_status/sqlc_lib"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/video_status"
)

var testQueries *sqlc_lib.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db : ", err)
	}

	testQueries = sqlc_lib.New(testDB)
	os.Exit(m.Run())
}
