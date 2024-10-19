docker run -d --name some-clickhouse-server -p18123:8123 -p19000:9000 --ulimit nofile=262144:262144 clickhouse/clickhouse-server
