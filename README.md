# Rabbitmq data copier using golang
A golang project which facilitates copying from one rabbitmq server queue to a separate rabbitmq queue

To run:<br>
1- Edit docker-compose file and change the environment variables according to your needs.<br>
```
environment:
  - TARGET_USER=admin
  - TARGET_PASSWORD=admin
  - TARGET_HOST=127.0.0.1
  - TARGET_V_HOST=/live
  - TARGET_EXCHANGE_NAME=app
  - TARGET_ROUTING_KEY=products.*
  - SOURCE_USER=admin
  - SOURCE_PASSWORD=admin
  - SOURCE_HOST=127.0.0.1
  - SOURCE_V_HOST=/live
  - SOURCE_QUEUE_NAME=app2.products
```

2- run `docker-compose up`
