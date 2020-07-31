## Clickhouse compose project

### Configuration

- [x] Основной конфигурационный файл сервера - `config.xml`. Он расположен в директории `/etc/clickhouse-server/`.
- [x] Отдельные настройки могут быть переопределены в файлах `*.xml` и `*.conf` из директории config.d рядом с конфигом.
- [x] Также в конфиге могут быть указаны «подстановки». Если у элемента присутствует атрибут `incl`, то в качестве значения будет использована соответствующая подстановка из файла. 

```xml
<zookeeper incl="zookeeper-servers" optional="true" />
```

По умолчанию, путь к файлу с подстановками - `/etc/metrika.xml`. Он может быть изменён в конфигурации сервера в элементе `include_from`.

- [x] В `config.xml` может быть указан отдельный конфиг с настройками пользователей, профилей и квот.  
Относительный путь к нему указывается в элементе `users_config`. 
По умолчанию - `users.xml`. Если users_config не указан, то настройки пользователей, профилей и квот, указываются непосредственно в `config.xml`.

```yaml
volumes:
      - ./configuration/config.xml:/etc/clickhouse-server/config.xml
      - ./configuration/macroses/macros-01.xml:/etc/clickhouse-server/config.d/macros.xml
      - ./configuration/metrika.xml:/etc/clickhouse-server/metrika.xml
      # Для users_config могут также существовать переопределения в файлах из директории users_config.d (например, users.d) и подстановки. 
      - ./configuration/users.xml:/etc/clickhouse-server/users.xml
```

### Compose

- [x] Create clickhouse network `$ docker network create clickhouse-net`
- [x] Check `$ docker network ls`
- [x] Run `$ docker-compose up -d`

**Output**

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

- [x] Check `$ docker container ls -a`

**Output**

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

### Connect to one of cluster server

- [x] `$ docker run -it --rm --network="clickhouse-net" --link clickhouse-01:clickhouse-server yandex/clickhouse-client --host clickhouse-server`
- [x] Check: `SELECT * FROM system.clusters;`

**Output**

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

### Distributed

`// todo`

### Following manuals

- [x] https://github.com/zikwall/clickhouse-docs