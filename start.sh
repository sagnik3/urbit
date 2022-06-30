#!/bin/bash
set -e

trap 'killall distribkv' SIGINT

cd $(dirname $0)

killall distribkv || true
sleep 0.1

go install -v

urbit -db-location=moscow.db -http-addr=127.0.0.2:8080 -config-file=sharding.toml -shard=Moscow &
urbit -db-location=moscow-r.db -http-addr=127.0.0.22:8080 -config-file=sharding.toml -shard=Moscow -replica &

urbit -db-location=minsk.db -http-addr=127.0.0.3:8080 -config-file=sharding.toml -shard=Minsk &
urbit -db-location=minsk-r.db -http-addr=127.0.0.33:8080 -config-file=sharding.toml -shard=Minsk -replica &

urbit -db-location=kiev.db -http-addr=127.0.0.4:8080 -config-file=sharding.toml -shard=Kiev &
urbit -db-location=kiev-r.db -http-addr=127.0.0.44:8080 -config-file=sharding.toml -shard=Kiev -replica &

urbit -db-location=tashkent.db -http-addr=127.0.0.5:8080 -config-file=sharding.toml -shard=Tashkent &
urrbit -db-location=tashkent-r.db -http-addr=127.0.0.55:8080 -config-file=sharding.toml -shard=Tashkent -replica &

wait