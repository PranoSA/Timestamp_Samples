package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

/**
	 *  CREATE TABLE go_times (
id SERIAL PRIMARY KEY,
inserted_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
timer TIMESTAMP,
timer_tz TIMESTAMP WITH TIME ZONE
);

	 *
*/

//add http listener that takes json request pOst request
//parse the json request and insert the timestamp into the database
//query the database for the timestamp

	 func insertTimestamp(w http.ResponseWriter, r *http.Request) {
		// Parse the JSON request
		var request struct {
			Timestamp time.Time `json:"timestamp"`
		}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insert the timestamp into the database
		_, err = db.Exec(context.Background(), "INSERT INTO go_times (timer, timer_tz) VALUES ($1, $2)", request.Timestamp, request.Timestamp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}

	//Now we need to query the database for the timestamps
	func getTimestamp(w http.ResponseWriter, r *http.Request) {


		type result struct {
			Id int `json:"id"`
			Timestamp time.Time `json:"timestamp"`
			Timestamp_TZ time.Time `json:"timestamp_tz"`
			Created time.Time `json:"created"`
		}

		var results []result

		//queery all rows 
		rows, err := db.Query(context.Background(), "SELECT id, timer, timer_tz, inserted_at FROM go_times")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		for rows.Next() {
			var result result
			err := rows.Scan(&result.Id, &result.Timestamp, &result.Timestamp_TZ, &result.Created)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//convert the timestamp to UTC
			//result.Created = result.Created.UTC()
			result.Timestamp_TZ = result.Timestamp_TZ.UTC()


			results = append(results, result)
		}
	

		json.NewEncoder(w).Encode(results)
	}



func main(){

	// start connection to postgres
	//host = localhost, port = 5432, user = postgres, password = password, dbname = test
	//db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=test sslmode=disable")


	//postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
	conn_string := "postgres://username:password@localhost:5432/spring_timestamp?sslmode=disable&options=-c%20timezone=America%2New_York"

	//using pgxpool
	//config, err := pgxpool.ParseConfig("host=localhost port=5432 user=username password=password dbname=spring_timestamp sslmode=disable options='-c timezone=UCT'")
	config, err := pgxpool.ParseConfig(conn_string)
	if err != nil {
		panic(err)
	}

	config.ConnConfig.RuntimeParams["timezone"] = "UTC"
	config.ConnConfig.RuntimeParams["TimeZone"] = "UTC"
	config.ConnConfig.Config.RuntimeParams["timezone"] = "UTC"
	config.ConnConfig.Config.RuntimeParams["TimeZone"] = "UTC"
	//still PDT for some reason
	//config.ConnConfig.RuntimeParams["timezone"] = "PDT"
	config.ConnConfig.Config.RuntimeParams["timezone"] = "UTC"
	

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
	db = pool
	//defer db.Close()

	//test connection 
	err = db.Ping(context.Background())



	if err != nil {
		panic(err)
	}

	// Create A Timestamp, using PDT
	timestamp := time.Date(2021, time.March, 14, 1, 59, 59, 0, time.FixedZone("PDT", -7*60*60))


	// start listening to the http server
	http.HandleFunc("/insert", insertTimestamp)
	http.HandleFunc("/get", getTimestamp)

	http.ListenAndServe(":8080", nil)


	// Insert the timestamp into the database
	_, err = pool.Exec(context.Background(), "INSERT INTO test_table (timestamp) VALUES ($1)", timestamp)

	if err != nil {
		panic(err)
	}

	// Query the database for the timestamp
	var result time.Time

	err = pool.QueryRow(context.Background(), "SELECT timestamp FROM test_table").Scan(&result)

	if err != nil {
		panic(err)
	}

}

/*

*/