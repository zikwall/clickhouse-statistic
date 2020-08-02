## Advanced manual for Clickhouse Statisitc project

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