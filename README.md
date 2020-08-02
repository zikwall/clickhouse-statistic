## Clickhouse Container statistic project

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
    --topic MyTopicName \
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
    --group jacek-japila-pl \
    --topic t1
```

- [x] Run Kafka console producer

```shell script
docker exec -it clickhouse-kafka \
  kafka-console-producer.sh \
    --broker-list :9092 \
    --topic MyTopic1
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

### Following manuals

- [x] https://github.com/zikwall/clickhouse-docs
- [x] [Advanced manual](./ADVANCED.md)