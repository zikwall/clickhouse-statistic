CREATE TABLE queue (
    `user_id` UInt32,
    `app` String,
    `host` String,
    `event` String,
    `ip` String,
    `guid` String,
    `created_at` DateTime('Europe/London')
) ENGINE = Kafka('clickhouse-kafka:9092', 'ClickhouseTopic', 'my-group-ch', 'JSONEachRow') SETTINGS kafka_num_consumers = 2;