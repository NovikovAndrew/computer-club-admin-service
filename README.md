# Admin compiki

Admin compiki is service for BTB users and admins of comps 

## Run Docker Containers

Before setup project make sure you start mongodb and postgres containers is running

## Setup project

If you want to run service use command in terminal from PHONY list 

```bash
run-admin-service
```

or 

```bash
cd cmd/main
go run main.go
```

App dependencies

```
github.com/sirupsen/logrus
go.mongodb.org/mongo-driver/mongo
github.com/julienschmidt/httprouter
```


## License

[Andrey Novikov]() [no Novitsky pidarasik]()
