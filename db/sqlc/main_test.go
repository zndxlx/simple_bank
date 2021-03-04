package db

import (
    "testing"
    "database/sql"
     _ "github.com/lib/pq"
    "log"
    "os"
)

const (
    DBDriver = "postgres"
    DBSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueies *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
    var err error
    testDB, err = sql.Open(DBDriver, DBSource)
    if err != nil {
        log.Fatalf("cannot connet to db:%v", DBSource)
    }
    testQueies = New(testDB)
    
    r := m.Run()
    os.Exit(r)
}
