# helm install -f ./helm-charts/values/values.mysql.yaml temporaldev -n temporal-starship ./helm-charts

```azure
#in https://github.com/temporalio/temporal git repo dir
export SQL_PLUGIN=mysql
export SQL_HOST=mysql_host
export SQL_PORT=3306
export SQL_USER=mysql_user
export SQL_PASSWORD=mysql_password

./temporal-sql-tool create-database -database temporal
SQL_DATABASE=temporal ./temporal-sql-tool setup-schema -v 0.0
SQL_DATABASE=temporal ./temporal-sql-tool update -schema-dir schema/mysql/v57/temporal/versioned

./temporal-sql-tool create-database -database temporal_visibility
SQL_DATABASE=temporal_visibility ./temporal-sql-tool setup-schema -v 0.0
SQL_DATABASE=temporal_visibility ./temporal-sql-tool update -schema-dir schema/mysql/v57/visibility/versioned
```

