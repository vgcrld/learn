package funcs

/*
	Need to Port forward the native client port 9000

	kubectl port-forward --namespace default svc/core-clickhouse-default-1 9000

*/

import (
	"fmt"
	"log"
	"time"

	"github.com/vgcrld/all-things-go/cmd/queries"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

func RunClickhouse() {

	connect, err := sqlx.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		log.Println("Error on connection")
		log.Fatal(err)
	}
	defer closeDB(connect)

	var items []struct {
		Item_id              uint64    `db:"item_id"`
		Name                 string    `db:"name"`
		Tags                 []string  `db:"tags"`
		Key                  string    `db:"key"`
		Type                 string    `db:"type"`
		File_key             string    `db:"file_key"`
		First_epoch_ns       int64     `db:"first_epoch_ns"`
		Last_epoch_ns        int64     `db:"last_epoch_ns"`
		First_trend_epoch_ns int64     `db:"first_trend_epoch_ns"`
		Last_trend_epoch_ns  int64     `db:"last_trend_epoch_ns"`
		Alias                string    `db:"alias"`
		Reporting            uint8     `db:"reporting"`
		Loading              uint8     `db:"loading"`
		Visible              uint8     `db:"visible"`
		Insert_ts            time.Time `db:"insert_ts"`
		Update_ts            time.Time `db:"update_ts"`
		First_ts             time.Time `db:"first_ts"`
		Last_ts              time.Time `db:"last_ts"`
		First_trend_ts       time.Time `db:"first_trend_ts"`
		Last_trend_ts        time.Time `db:"last_trend_ts"`
		First_epoch          int64     `db:"first_epoch"`
		Last_epoch           int64     `db:"last_epoch"`
		First_trend_epoch    int64     `db:"first_trend_epoch"`
		Last_trend_epoch     int64     `db:"last_trend_epoch"`
	}

	if err := connect.Select(&items, queries.Items); err != nil {
		log.Fatal(err)
	}

	formatString := "%-10v %-25v %-50v %-5v %-25v %-25s\n"
	fmt.Printf(formatString, "ID", "Alias", "Name", "Vis", "Type", "Tags")
	for _, v := range items {
		fmt.Printf(formatString, v.Item_id, v.Alias, v.Name, v.Visible, v.Type, v.Tags)
	}

}

func closeDB(db *sqlx.DB) {
	log.Println("Closing the DB.")
	db.Close()
}
