# golang-with-kafka


# Preparando Kafka

Levantar toda la estructura de Kafka y Zookeeper.

    $ docker-compose up

 Con esta intrucción se va a crear el topic fogo-chat con 4 particiones y un factor de replicación de 2.

    $ docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic fogo-chat --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181

