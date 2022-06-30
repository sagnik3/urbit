#!/bin/bash
set -e

trap 'killall urbit' SIGINT

cd $(dirname $0)

killall distribkv || true
sleep 0.1

go install -v

urbit -db-location=db1.db -http-addr=127.0.0.2:8080 -config-file=sharding.toml -shard=DB1 &
urbit -db-location=db1-r.db -http-addr=127.0.0.22:8080 -config-file=sharding.toml -shard=DB1 -replica &

urbit -db-location=db2.db -http-addr=127.0.0.3:8080 -config-file=sharding.toml -shard=DB2 &
urbit -db-location=db2-r.db -http-addr=127.0.0.33:8080 -config-file=sharding.toml -shard=DB2 -replica &

urbit -db-location=db3.db -http-addr=127.0.0.4:8080 -config-file=sharding.toml -shard=DB3 &
urbit -db-location=db3-r.db -http-addr=127.0.0.44:8080 -config-file=sharding.toml -shard=DB3 -replica &

urbit -db-location=db4.db -http-addr=127.0.0.5:8080 -config-file=sharding.toml -shard=Tashkent &
urbit -db-location=db4-r.db -http-addr=127.0.0.55:8080 -config-file=sharding.toml -shard=Tashkent -replica &

wait