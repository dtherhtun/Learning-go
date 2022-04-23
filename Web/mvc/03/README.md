# CRUD mongodb

### CREATE user
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"James Bond","gender":"male","age":32,"id":"1"}' http://localhost:8080/user
```

### READ user
```
curl http://localhost:8080/user/1
```

### UPDATE user
```
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Dylan","gender":"male","age":32,"id":"1"}' http://localhost:8080/user/1
```

### DELETE user
```
 curl -X DELETE http://localhost:8080/user/1
```