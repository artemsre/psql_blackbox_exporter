# psql_blackbox_exporter
The goal is to have cheap ans easy check for psql credentials and postgresql server availability. 

The service iterate thought enviropment variables and uses any with name 'psql_.*' as connection string. 

Run 'SELECT count(state),state FROM pg_stat_activity where state<>'idle' group by state;' for each of them and provide gauge metrics like this:

psql_query_state{dbcon="psql_users_app_pre",state="active"} 1
psql_query_state{dbcon="psql_action_app_stg",state="active"} 1
psql_query_state{dbcon="psql_action_app_stg",state="idle in transaction"} 2

#todo 
- create vector for query errors count

