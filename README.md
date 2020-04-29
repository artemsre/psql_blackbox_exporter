# psql_blackbox_exporter
The goal is to have cheap ans easy check for psql credentials and postgresql server availability. 

The service iterate thought enviropment variables and uses any with name 'psql_.*' as connection string. 

Run 'SELECT count(state),state FROM pg_stat_activity where state<>'idle' group by state;' for each of them and provide gauge metrics like this:

```# HELP psql_query_errors_total psql query error by db connections
# TYPE psql_query_errors_total counter
psql_query_errors_total{dbcon="psql_users_app_pre"} 1
psql_query_errors_total{dbcon="psql_action_app_stg"} 1
# HELP psql_query_state psql query by user with state
# TYPE psql_query_state gauge
psql_query_state{dbcon="psql_users_app_pre",state="active"} 1
psql_query_state{dbcon="psql_action_app_stg",state="active"} 1
psql_query_state{dbcon="psql_action_app_stg",state="idle in transaction"} 2
```
