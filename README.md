# go-mongorm
This is a MongoDB ORM written with generics

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
_, err := mongorm.New(mongorm.NewConfig().SetDBName("emp_local").SetURI("mongo://127.0.0.1:27017"))
```
3. Run a Find Query on users:
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

## TODO:
- Complete CRUD methods