#!/bin/bash

echo "GRANT ALL ON \`api\`.* TO '"$MYSQL_USER"'@'%' ;" | "${mysql[@]}"
echo 'FLUSH PRIVILEGES ;' | "${mysql[@]}"

"${mysql[@]}" < /docker-entrypoint-initdb.d/1_schema.sql_
