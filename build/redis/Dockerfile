FROM redis:7.2.3

WORKDIR /usr/local/etc/redis

# Copy a custom configuration file to the image
COPY redis.conf redis.conf

# Copy env file from project
# This sh file is automate to add redis user
# based on env file of the project
COPY .env .env
COPY redis-add-user.sh redis-add-user.sh
RUN chmod +x redis-add-user.sh
RUN ./redis-add-user.sh

CMD ["redis-server", "redis.conf"]

EXPOSE 6379
