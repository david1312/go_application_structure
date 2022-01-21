# go_application_structure

Learning Go

# Compiling go application into different OS Example
GOOS = For operating system
GOARCH = For architecture or cpu type

Compiling for windows
```shell script
GOOS=windows GOARCH=amd64 go build -o winapp.exe
```

Compiling for linux
```shell script
GOOS=linux GOARCH=amd64 go build -o linuxapp
```

Compiling for mac
```shell script
GOOS=darwin GOARCH=amd64 go build -o macapp
```