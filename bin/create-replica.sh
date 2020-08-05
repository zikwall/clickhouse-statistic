#!/bin/bash

for filename in $(pwd)/example/database/*.sql; do
  file=$(basename $filename)

  if [[ "$file" == "distributed.sql" ]]; then
    echo "Skip distributed table"
    continue
  fi

  echo "Import schema from $file"

  sql=$(<$filename)
  docker run --rm --network="clickhouse-net" --link clickhouse-$1:clickhouse-server yandex/clickhouse-client -m --host clickhouse-server --query="$sql"
done