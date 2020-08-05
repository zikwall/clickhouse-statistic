single:
	docker run -d --rm --name clickhouse-server -p 9000:9000 --ulimit nofile=262144:262144 yandex/clickhouse-server

client:
	docker run -it --rm --link clickhouse-server:clickhouse-server yandex/clickhouse-client  --host clickhouse-server

cluster-client:
	docker run -it --rm --network="clickhouse-net" --link clickhouse-$(ch):clickhouse-server yandex/clickhouse-client --host clickhouse-server

cluster-client-credentials:
	docker run -it --rm --network="clickhouse-net" --link clickhouse-01:clickhouse-server yandex/clickhouse-client --host clickhouse-server -u root_example --password 123456