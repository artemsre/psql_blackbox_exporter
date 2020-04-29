package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	psql_query_state := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "psql_query_state",
		Help: "psql query by user with state",
	}, []string{"dbcon", "state"})
	psql_query_errors_total := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "psql_query_errors_total",
		Help: "psql query error by db connections",
	}, []string{"dbcon"})
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)
	for {
		for _, e := range os.Environ() {
			envLine := strings.SplitN(e, "=", 2)
			if strings.HasPrefix(envLine[0], "psql_") {
				//				fmt.Println(envLine[0])
				//				fmt.Println(envLine[1])
				db, err := sql.Open("postgres", envLine[1])
				if err != nil {
					psql_query_errors_total.WithLabelValues(envLine[0]).Inc()
					continue
				}
				defer db.Close()
				rows, err := db.Query("SELECT count(state),state FROM pg_stat_activity where state<>'idle' group by state;")
				if err != nil {
					psql_query_errors_total.WithLabelValues(envLine[0]).Inc()
					continue
				}
				defer rows.Close()
				for rows.Next() {
					var count int
					var name string
					err = rows.Scan(&count, &name)
					if err != nil {
						print(err.Error())
					}
					psql_query_state.WithLabelValues(envLine[0], name).Set(float64(count))
					fmt.Println(name, count)
				}
			}
		}
		time.Sleep(60 * time.Second)
	}
}
