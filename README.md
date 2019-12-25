# Real time metrics
Visualizing real-time metrics from golang server (using WebSockets and plain-javascript)


## Backend
It is a service that exposes certain endpoints to get data from sources like database.
In this case we are exposing "/stream" endpoint which gives us current cpu usage. 
Here it is returning randomly generated data. 
This is written in MVC style for better division of responsibilities.

#### To start backend:
``` bash
// Starts server on :8082, To change port edit backend/stream.go, proxy/main.go (since proxy needs to know backend location) 
go run cmd/main.go
```

## Frontend: 
It is a service that serves html/js files to the client. 
This service talks to backend for pulling dynamic data, in this case live CPU usage. (connection goes through proxy)

#### To start frontend:
``` bash
// Starts server on :8080, To change port edit frontend/main.go
go run main.go
```


## Proxy
It is a service which behaves as proxy/gateway to all the backend services.
This service can perform various functions such as authenticating user, auditing sessions, routing requests to
particular backend service. Here it is forwarding requests to backend.

#### To start proxy:
``` bash
// Starts server on :8081, To change port edit proxy/main.go, frontend/views/index.html (since frontend needs to know proxy location) 
go run main.go
```