#!/bin/sh

# Set pass and user based on .envfile
REDIS_USER=$(grep '^REDIS_USER' .env | cut -d '=' -f2)
REDIS_PASS=$(grep '^REDIS_PASS' .env | cut -d '=' -f2)

echo "$REDIS_USER, $REDIS_PASS"
sed -i "1015i\user $REDIS_USER on >$REDIS_PASS +@ALL ~*" redis.conf
