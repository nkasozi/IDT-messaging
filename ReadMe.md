# IDT Messaging Api

This API functionality for creating and retrieving users from the IDT backend.

The Go-Kit library which provides industry standard building blocks for microservices was used extensively

## Routes Exposed
| Route                         | Method | Sample Request  | Sample Response                |
|-------------------------------|--------| ----------------| ---------------
| localhost:8080/v1/users            | PUT (Set User)  | {"id": "test-user-2","name": "test-user-2","signUpTime": 6000} |{"id": "test-user-2","name": "test-user-2","signUpTime": 6000}
| localhost:8080/v1/users            | GET (List Users)   |  |{"users":[{"id": "test-user-2","name": "test-user-2","signUpTime": 6000}]}
| localhost:8080/v1/users/{users-id}           | GET (Get User)   |  |{"id": "test-user-2","name": "test-user-2","signUpTime": 6000}

**NB: All Endpoints need an authorization header:** 
- ``` {"Header":"Authorization","value":"dummy-hard-coded-api-key"}] ```
## How to Run it Locally

You need to install docker first. Then you can run

```
docker build --tag idt-messaging-core-app .

docker run -p 127.0.0.1:8080:8080/tcp idt-messaging-core-app
```
## How to Run the Unit Tests
```
go test -v ./...
```
