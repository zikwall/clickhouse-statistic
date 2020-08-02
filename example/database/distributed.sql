CREATE TABLE main
(
    `user_id` UInt32,
    `app` String,
    `host` String,
    `event` String,
    `ip` String,
    `guid` String,
    `created_at` DateTime('Europe/London')
)
ENGINE = Distributed(cluster_1, default, main)