CREATE TABLE queue (
    `user_id` UInt32,
    `app` String,
    `host` String,
    `event` String,
    `ip` String,
    `guid` String,
    `created_at` DateTime('Europe/London')
) ENGINE = Kafka('0.0.0.0:9092', 'MyTopic', 'my-group', 'JSONEachRow') SETTINGS kafka_num_consumers = 2;