# Golang API

## Quick Setup

### Install Dependent Packages

Install AIR;
```sh
go install github.com/cosmtrek/air@latest
```

Install SWAG;
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

### Install APP

Install packages;
```sh
go install
``` 

Generate API documentations;
```sh
swag init
``` 

### Run APP

```sh
air
``` 
