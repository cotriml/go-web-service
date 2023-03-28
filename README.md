# go-web-service

### This is an example of REST interfaces built using only net/http. 


To execute the application, just navigate to its root directory and run the command below: 

```
go run .
```

## GET /healthz

Status code 200
```
{
    "status": "OK"
}
```

## GET /users

Status code 200
```
[
    {
        "id": "123",
        "firstName": "Lucas",
        "lastName": "Cotrim",
        "age": 28
    },
    {
        "id": "456",
        "firstName": "Lucas",
        "lastName": "Machado",
        "age": 40
    }
]
```

## GET /users/{id}

Status code 200
```
{
    "id": "123",
    "firstName": "Lucas",
    "lastName": "Cotrim",
    "age": 28
}
```

## POST /users
Status code 201

## PUT /users/{id}
Status code 204

## DELETE /users/{id}
Status code 204

