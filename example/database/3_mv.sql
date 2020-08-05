CREATE MATERIALIZED VIEW mainconsumer TO main
    AS SELECT `user_id`, `app`, `host`, `event`, `ip`, `guid`, `created_at`
FROM queue;