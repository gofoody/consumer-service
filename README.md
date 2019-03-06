# consumer-service

Run - start consumer service at `localhost:8071`
```
go run main.go
```


Check status - return consumer service status
```
curl localhost:8071/api/status
```

Get consumer - return consumer with given consumerId
```
curl http://localhost:8071/api/consumers/1
```

Create consumer - create new consumer with given consumer json and return consumerId
```
curl -X POST localhost:8071/api/consumers -d '{"name":"hemal"}'
```
