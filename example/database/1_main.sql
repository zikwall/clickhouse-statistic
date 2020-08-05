CREATE TABLE main
(
    `user_id` UInt32,
    `app` String,
    `host` String,
    `event` String,
    `ip` String,
    `guid` String,
    `created_at` DateTime('Europe/London')
) ENGINE = ReplicatedReplacingMergeTree('/clickhouse/tables/{layer}-{shard}/table_name', '{replica}')
PARTITION BY toYYYYMM(created_at)
ORDER BY (event, created_at, intHash32(user_id))
SAMPLE BY intHash32(user_id);
