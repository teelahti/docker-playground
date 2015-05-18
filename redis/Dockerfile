FROM redis

COPY redis.conf /usr/local/etc/redis/redis.conf
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]

# There is no EXPOSE on purpose: all traffic is via internal links
# https://docs.docker.com/examples/running_redis_service/