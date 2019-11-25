# gRPC Server written in Golang
To run the server, use the following commands
```
go build main.go
./main
```

To connect to server you need to install 'evans':
```
brew tap ktr0731/evans 
brew install evans 
```

After that, open another terminal and enter:
```
evans protos/uuid.proto -p 8080
call IdSend
```

###### Response type: {uuid: <uuid>}
