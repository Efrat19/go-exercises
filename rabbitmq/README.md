# rabbitmq-go

-  set the rabbitmq server:
    ```
    docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
    ```
- run the program:
    ```
    go build . && ./rabbitmq
    ```