#!/bin/bash

distributed=$(<(pwd)/example/database/distrubuted.sql)
docker run --rm --network="clickhouse-net" --link clickhouse-$1:clickhouse-server yandex/clickhouse-client -m --host clickhouse-server --query="$distributed"