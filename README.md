


#Run directly
Open one terminal:
```bash
go get -u -v golang.org/x/net/websocket
go run main.go
```


#Dockerized

```bash
docker build -t tanjinfu/go-websocket:v1 .
docker run -i -t -p 8080:8080 tanjinfu/go-websocket:v1
```