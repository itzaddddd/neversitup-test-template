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
|  ├── constant                       # Constant value
|  |  ├── constant.go
|  ├── server
|  |  ├── server.go
|  ├── api
|  |  ├── middleware
|  |  |  ├── middleware.go
|  |  ├── http                        # Define api route
|  |  |  ├── order.go
|  ├── handler                        # Handler is for handling http request and response
|  |  ├── order_handler.go
|  |  ├── handler.go
|  ├── service                        # Service is core business logic
|  |  ├── order_service.go 
|  |  ├── service.go
|  ├── repository                     # Repository is for working with db
|  |  ├── order_repository.go
|  |  ├── repository.go
|  ├── db
|  |  ├── postgresql
|  |  |  ├── connection.go
|  |  |  ├── migration
|  |  |  |  ├── 000001_up.go
|  |  |  |  ├── 000001_down.go
|  ├── external                       # Calling the third-party service
|  |  ├── payment_service
|  |  |  ├── payment_service.go
|  ├── util                           # Utility function
|  |  ├── util.go
|  ├── docs
|  |  ├── email_template
|  |  |  ├── mail_1.html
|  ├── test                           # Testing script
|  |  ├── loadtest
|  |  |  ├── test.go
```