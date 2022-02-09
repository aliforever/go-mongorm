# go-mongorm
This is a MongoDB ORM written with generics

## Requirement
```
go install golang.org/dl/go1.18beta2@latest
go1.18beta2 download
```

## Install
```go get -u github.com/aliforever/go-mongorm```

## Use
1. Define your struct as a Model interface (it should have a Collection() method):
```go
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	FullName string             `bson:"full_name"`
	Username string             `bson:"username"`
}

func (User) Collection() string {
	return "users"
}
```
2. Define your settings
```go
_, err := mongorm.New(mongorm.NewConfig().SetDBName("mydb").SetURI("mongo://127.0.0.1:27017"))
```
3. Write a Find Query on users:
```go
user, err := mongorm.FindOneWithFilter[User](bson.M{
    "username": "admin",
})
if err != nil {
    fmt.Println(err)
    return
}

fmt.Printf("%#v\n", user)
```
4. Run the project:
```shell
go1.18beta2 main.go
```

## TODO:
- Complete CRUD methods