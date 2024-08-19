## Steps setup the project

1) Use Sqlc to generte type safe code for users-service:
```
$ cd users-service

$ sqlc generate
```

if you have a problem with GO_PATH, try:

```ssh
$ export GO_PATH=~/go

$ export PATH=$PATH:/$GO_PATH/bin
```

2) Generate JWT keys:

```ssh
$ openssl genrsa -out jwtRSA256-private.pem 2048

$ openssl rsa -in jwtRSA256-private.pem -pubout -outform PEM -out jwtRSA256-public.pem
```
