package main

import (
	"github.com/lib/pq"
	"github.com/coopernurse/gorp"
	"database/sql"
	"log"
	_ "strings"
	"time"
)

type Sr struct {
    Id   				int32 			`db:"id"`
    Version				int32			`db:"version"`
    ParentId			sql.NullInt64 	`db:"parent_id"`
    Summary				string			`db:"summary"`
    Description			sql.NullString	`db:"descr"`
    
    CreatedById			sql.NullInt64 	`db:"created_by_id"`
    CreatedAt			time.Time		`db:"created_at"`
    
    RequestedById		int32		 	`db:"requested_by_id"`
    LastUpdatedAt		pq.NullTime		`db:"last_updated_at"`
    
    StatusId			int32 			`db:"status_id"`
    SrTypeId			int32			`db:"sr_type_id"`
    
    IsPublic			bool			`db:"is_public"`
    
    StartDate			pq.NullTime		`db:"start_date"`
    
    PercentComplete		sql.NullInt64 	`db:"pct_complete"`
    
    CompletionDate		pq.NullTime		`db:"completion_date"`
    CancellationDate	pq.NullTime		`db:"cancellation_date"`
    
    IsEmergency			bool			`db:"is_emergency"`
    
    WorkSequence		sql.NullInt64 	`db:"work_sequence"`
    RequestorPriority	sql.NullInt64 	`db:"requestor_priority"`
    
    EstimatedHours		sql.NullFloat64	`db:"estimated_hours"`
    TagsAsString 		string			`db:"tags"`
    AssigneesAsString 	string			`db:"assignees"`
    SubscribersAsString string			`db:"subscribers"`
}

func main() {
	//db, err := sql.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable")
	dbmap := initDb()
    defer dbmap.Db.Close()	

	var sr Sr
	err := dbmap.SelectOne(&sr, "select * from sr.srs where id = $1", 3484)
    checkErr(err, "SelectOne failed")
    log.Println("sr row:", sr)
}

func initDb() *gorp.DbMap {
    db, err := sql.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable")
    checkErr(err, "sql.Open failed")

    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

    dbmap.AddTableWithName(Sr{}, "srs").SetKeys(true, "Id")

    return dbmap
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}