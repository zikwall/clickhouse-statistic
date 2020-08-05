<div align="center">
  <h1>Clickhouse container statistic project</h1>
  <h4>Your mini assistant in future problems :)</h4>
  <h5>This guide contains step by step instructions with working examples of deploying Clickhouse, Apache Kafka, Zookeeper in Docker</h5>
</div>

### Getting Started

- [x] `$ mkdir -p /shared/ch/{clickhouse,zookeeper,kafka}`
- [x] First step create clickhouse network `$ docker network create clickhouse-net`
- [x] You can check result of prev. step `$ docker network ls`
- [x] If OK, then run `$ docker-compose up -d`

<details>
  <summary>Output</summary>
  
  ```shell script
    msi@msi clickhouse-compose # docker-compose up -d
    Starting clickhouse-zookeeper ... done
    Recreating clickhouse-04      ... done
    Recreating clickhouse-05      ... done
    Recreating clickhouse-01      ... done
    Recreating clickhouse-02      ... done
    Recreating clickhouse-06      ... done
    Recreating clickhouse-03      ... done
  ```
</details>

- [x] Again check `$ docker container ls -a`

<details>
  <summary>Output</summary>
  
  ```shell script
    CONTAINER ID        IMAGE                      COMMAND                  CREATED             STATUS              PORTS                                                            NAMES
    442a79a43f3a        yandex/clickhouse-server   "/entrypoint.sh"         2 minutes ago       Up 2 minutes        8123/tcp, 9009/tcp, 0.0.0.0:9003->9000/tcp                       clickhouse-03
    f5279aec0e37        yandex/clickhouse-server   "/entrypoint.sh"         2 minutes ago       Up 2 minutes        8123/tcp, 9009/tcp, 0.0.0.0:9006->9000/tcp                       clickhouse-06
    3a783ee75502        yandex/clickhouse-server   "/entrypoint.sh"         2 minutes ago       Up 2 minutes        8123/tcp, 9009/tcp, 0.0.0.0:9002->9000/tcp                       clickhouse-02
    ace4df988157        yandex/clickhouse-server   "/entrypoint.sh"         2 minutes ago       Up 2 minutes        8123/tcp, 9009/tcp, 0.0.0.0:9001->9000/tcp                       clickhouse-01
    a40ac11a5194        yandex/clickhouse-server   "/entrypoint.sh"         2 minutes ago       Up 2 minutes        8123/tcp, 9009/tcp, 0.0.0.0:9005->9000/tcp                       clickhouse-05
    23495201a490        yandex/clickhouse-server   "/entrypoint.sh"         2 minutes ago       Up 2 minutes        8123/tcp, 9009/tcp, 0.0.0.0:9004->9000/tcp                       clickhouse-04
    8de765edf713        zookeeper                  "/docker-entrypoint.…"   4 minutes ago       Up 2 minutes        2888/tcp, 3888/tcp, 0.0.0.0:2181-2182->2181-2182/tcp, 8080/tcp   clickhouse-zookeeper
    ... other own containers
  ```
</details>

- [x] For stopping all containers `$ docker-compose stop`

### Connect to one of cluster server

- [x] `$ docker run -it --rm --network="clickhouse-net" --link clickhouse-01:clickhouse-server yandex/clickhouse-client --host clickhouse-server`
- [x] Check: `SELECT * FROM system.clusters;`

```shell script
clickhouse-01 :) SELECT * FROM system.clusters;

SELECT *
FROM system.clusters

┌─cluster─────────────────────┬─shard_num─┬─shard_weight─┬─replica_num─┬─host_name─────┬─host_address─┬─port─┬─is_local─┬─user────┬─default_database─┬─errors_count─┬─estimated_recovery_time─┐
│ cluster_1                   │         1 │            1 │           1 │ clickhouse-01 │ 172.19.0.6   │ 9000 │        1 │ default │                  │            0 │                       0 │
│ cluster_1                   │         1 │            1 │           2 │ clickhouse-06 │ 172.19.0.7   │ 9000 │        0 │ default │                  │            0 │                       0 │
│ cluster_1                   │         2 │            1 │           1 │ clickhouse-02 │ 172.19.0.8   │ 9000 │        0 │ default │                  │            0 │                       0 │
│ cluster_1                   │         2 │            1 │           2 │ clickhouse-03 │ 172.19.0.4   │ 9000 │        0 │ default │                  │            0 │                       0 │
│ cluster_1                   │         3 │            1 │           1 │ clickhouse-04 │ 172.19.0.3   │ 9000 │        0 │ default │                  │            0 │                       0 │
│ cluster_1                   │         3 │            1 │           2 │ clickhouse-05 │ 172.19.0.5   │ 9000 │        0 │ default │                  │            0 │                       0 │
│ test_shard_localhost        │         1 │            1 │           1 │ localhost     │ 127.0.0.1    │ 9000 │        1 │ default │                  │            0 │                       0 │
│ test_shard_localhost_secure │         1 │            1 │           1 │ localhost     │ 127.0.0.1    │ 9440 │        0 │ default │                  │            0 │                       0 │
└─────────────────────────────┴───────────┴──────────────┴─────────────┴───────────────┴──────────────┴──────┴──────────┴─────────┴──────────────────┴──────────────┴─────────────────────────┘
```

### Create first data

- [x] Use `.sql` files for create tables from folder [example/database](./example/database)

### How work with Apache Kafka in Docker

- [x] Create new topic

```shell script
docker exec -t clickhouse-kafka \
  kafka-topics.sh \
    --bootstrap-server :9092 \
    --create \
    --topic ClickhouseTopic \
    --partitions 3 \
    --replication-factor 1
```

- [x] Print out the topics

```shell script
docker exec -t clickhouse-kafka \
  kafka-topics.sh \
    --bootstrap-server :9092 \
    --list
```

- [x] Describe

```shell script
docker exec -t clickhouse-kafka \
  kafka-topics.sh \
    --bootstrap-server :9092 \
    --describe \
    --topic MyTopic1
```

- [x] Run Kafka console consumer (run in another console)

```shell script
docker exec -t clickhouse-kafka \
  kafka-console-consumer.sh \
    --bootstrap-server :9092 \
    --group my-group \
    --topic MyTopic1
```

- [x] Run Kafka console producer

```shell script
docker exec -it clickhouse-kafka \
  kafka-console-producer.sh \
    --broker-list :9092 \
    --topic MyTopic1
```

- [x] Get count messages of topic

```shell script
docker exec -it clickhouse-kafka \
  kafka-run-class.sh \
  kafka.tools.GetOffsetShell \
    --broker-list :9092 \
    --topic MyTopic1 \
    --time -1
```

- [x] Drop topic

```shell script
docker exec -it clickhouse-kafka \
  kafka-topics.sh \
    --bootstrap-server :2181 \
    --topic MyTopic1 \
    --delete
```

**after put messages in producer console && to see messages printed out in second terminal, where run Kafka CLI consumer**

- [x] Show full logs fro Kafka run: `$ docker logs -f clickhouse-kafka`
- [x] [Kafkacat](https://github.com/edenhill/kafkacat)

### How wotk with Zookeeper

- [x] Connect container `$ docker exec -it clickhouse-zookeeper bash`
- [x] Connect server `$ bin/zkCli.sh -server 127.0.0.1:2181`
- [x] List root `$ ls /`

**Output**

```shell script
[admin, brokers, clickhouse, cluster, config, consumers, controller, controller_epoch, isr_change_notification, latest_producer_id_block, log_dir_event_notification, zookeeper]
```

- [x] `$ ls brokers/` after `$ ls brokers/topics`

**Output**

```shell script
[MyTopic, MyTopic1, __consumer_offsets]
```

- [x] `$ ls /consumers`

### Tests

- [x] First add next line `127.0.0.1    clickhouse-kafka` in `/etc/hosts`, for DEV, why? [1](https://ealebed.github.io/posts/2018/docker-%D1%81%D0%BE%D0%B2%D0%B5%D1%82-28-%D0%BA%D0%B0%D0%BA-%D0%B8%D1%81%D0%BF%D1%80%D0%B0%D0%B2%D0%B8%D1%82%D1%8C-%D0%BE%D1%88%D0%B8%D0%B1%D0%BA%D1%83-connection-reset-by-peer/), [2](https://github.com/grafana/metrictank/issues/1286), [3](https://github.com/wurstmeister/kafka-docker/issues/424)
- [x] Create topic `ClickhouseTopic` if already is not created
- [x] Run consumer in another terminal
- [x] Run app from example/backend `$ go run .`

<details>
  <summary>Output in Go terminal</summary>
  
  ```shell script
      Send message to broker: user 23, time 2020-08-04 13:58:14
      Send message to broker: user 16, time 2020-08-04 13:58:15
      Send message to broker: user 29, time 2020-08-04 13:58:16
      Send message to broker: user 11, time 2020-08-04 13:58:17
      Send message to broker: user 22, time 2020-08-04 13:58:18
      Send message to broker: user 25, time 2020-08-04 13:58:19
      Send message to broker: user 15, time 2020-08-04 13:58:20
      Send message to broker: user 20, time 2020-08-04 13:58:21
      Send message to broker: user 17, time 2020-08-04 13:58:22
      message at topic/partition/offset MyTopic/0/189:  = {"user_id":23,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:14"}
      message at topic/partition/offset MyTopic/0/190:  = {"user_id":16,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:15"}
      message at topic/partition/offset MyTopic/0/191:  = {"user_id":29,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:16"}
      message at topic/partition/offset MyTopic/0/192:  = {"user_id":11,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:17"}
      message at topic/partition/offset MyTopic/0/193:  = {"user_id":22,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:18"}
      message at topic/partition/offset MyTopic/0/194:  = {"user_id":25,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:19"}
      message at topic/partition/offset MyTopic/0/195:  = {"user_id":15,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:20"}
      message at topic/partition/offset MyTopic/0/196:  = {"user_id":20,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:21"}
      message at topic/partition/offset MyTopic/0/197:  = {"user_id":17,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:22"}
      Send message to broker: user 19, time 2020-08-04 13:58:23
      Send message to broker: user 18, time 2020-08-04 13:58:24
      Send message to broker: user 28, time 2020-08-04 13:58:25
  ```
</details>


<details>
  <summary>Output in consumer terminal</summary>
  
  ```shell script
      {"user_id":19,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:23"}
      {"user_id":18,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:24"}
      {"user_id":28,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:25"}
      {"user_id":14,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:26"}
      {"user_id":13,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:27"}
      {"user_id":14,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:28"}
      {"user_id":17,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:29"}
      {"user_id":22,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:30"}
      {"user_id":22,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:31"}
      {"user_id":26,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:32"}
      {"user_id":29,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:33"}
      {"user_id":10,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:34"}
      {"user_id":16,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:35"}
      {"user_id":21,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:36"}
      {"user_id":26,"app":"","host":"","event":"","ip":"","guid":"","created_at":"2020-08-04 13:58:37"}
  ```
</details>

- [x] Another terminal `make cluster-client` for connect `ch-01` server
- [x] `SELECT * from main;`

**Output**

```shell script
clickhouse-01 :) select * from main;

┌─user_id─┬─app─┬─host─┬─event─┬─ip─┬─guid─┬──────────created_at─┐
│      10 │     │      │       │    │      │ 2020-08-04 15:48:12 │
│      29 │     │      │       │    │      │ 2020-08-04 15:48:13 │
│      24 │     │      │       │    │      │ 2020-08-04 15:48:14 │
│      14 │     │      │       │    │      │ 2020-08-04 15:48:15 │
│      10 │     │      │       │    │      │ 2020-08-04 15:48:16 │
│      21 │     │      │       │    │      │ 2020-08-04 15:48:17 │
│      11 │     │      │       │    │      │ 2020-08-04 15:48:18 │
│      27 │     │      │       │    │      │ 2020-08-04 15:48:19 │
└─────────┴─────┴──────┴───────┴────┴──────┴─────────────────────┘
```

### Cluster

- [x] Create `main`, `queue` and `mainconsumer` tables each hosts `ch-`: 01, 02, 03, 04, 05, you can use `$ bin/create-replica.sh 03`
- [x] Create distributed table on last host `ch-06` from `example/database/distributed.sql`, or use `$ bin/create-distributed.sh 06`
- [x] connect to `ch-06`, u can see Makefile command `make cluster-client` and replace `clickhouse-01` to `clickhouse-06`
- [x] `SELECT COUNT() FROM main_distributed`;

**Output**

```shell script
clickhouse-06 :) select count() from main_distributed;

SELECT count()
FROM main_distributed

┌─count()─┐
│    1016 │
└─────────┘
```

- [x] Check connect one of server, example `ch-01`, `ch-02` and `SELECT count(*) from main;`

**Output**

```shell script
clickhouse-01 :) select count() from main;

SELECT count()
FROM main

┌─count()─┐
│     945 │
└─────────┘
```

**Output**

```shell script
clickhouse-02 :) select count() from main;

SELECT count()
FROM main

┌─count()─┐
│      71 │
└─────────┘
```

Happy! ^_^

### Following manuals

- [x] [Hard manual](https://github.com/zikwall/clickhouse-docs)
- [x] [Advanced manual](./ADVANCED.md)