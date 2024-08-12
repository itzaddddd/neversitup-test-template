### Project structure

```
├── go.mod
├── go.sum
├── README.md
├── Dockerfile
├── docker-compose.yml
├── .gitignore
├── cmd                               # Command to run main function
|  ├── app
|  |  ├── main.go
|  ├── migrate
|  |  ├── main.go
├── script                            # Script to run work
|  ├── run_app.sh
|  ├── migrate.sh
├── dev.env
├── sit.env
├── uat.env
├── src                               # Source files
|  ├── config
|  |  ├── config.go                   # Load .env config to go struct
|  |  ├── config_test.go
|  ├── constant                       # Constant value
|  |  ├── constant.go
|  ├── server
|  |  ├── server.go
|  |  ├── server_test.go
|  ├── api
|  |  ├── middleware
|  |  |  ├── middleware.go
|  |  |  ├── middleware_test.go
|  |  ├── http                        # Define api route
|  |  |  ├── order.go
|  |  |  ├── order_test.go
|  ├── handler                        # Handler is for handling http request and response
|  |  ├── order_handler.go
|  |  ├── order_handler_test.go
|  |  ├── handler.go
|  |  ├── handler_test.go
|  ├── service                        # Service is core business logic
|  |  ├── order_service.go
|  |  ├── order_service_test.go
|  |  ├── service.go
|  |  ├── service_test.go
|  ├── repository                     # Repository is for working with db
|  |  ├── order_repository.go
|  |  ├── order_repository_test.go
|  |  ├── repository.go
|  |  ├── repository_test.go
|  ├── db
|  |  ├── postgresql
|  |  |  ├── connection.go
|  |  |  ├── connection_test.go
|  |  |  ├── migration
|  |  |  |  ├── 000001_up.go
|  |  |  |  ├── 000001_down.go
|  ├── external                       # Calling the third-party service
|  |  ├── payment_service
|  |  |  ├── payment_service.go
|  |  |  ├── payment_service_test.go
|  ├── util                           # Utility function
|  |  ├── util.go
|  ├── docs
|  |  ├── email_template
|  |  |  ├── mail.html
|  ├── test                           # Testing script
|  |  ├── loadtest
|  |  |  ├── test.go
```
