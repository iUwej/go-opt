# go-opt . Optional values in golang.
go-opt is an implementation of optional values in go.\
Optional values can have a number of use cases such as:
- Providing optional parameters to functions
- Representing a look-up value that could be present or absent in particular context.

An optional value implements the following interface
```go
package opt

type Option[T any] interface {

	// Get returns value of type T if present or panic if absent
	Get() T

	// GetOrElse returns value of type T if present else return the value passed
	// as a parameter.
	GetOrElse(T) T

	// GetOrZero returns value of type T if present else returns the zero value of T.
	GetOrZero() T

	// Empty returns true if value is absent else returns false.
	Empty() bool
}
```

This module provides two implementations of the Option[T] interface
- Some[T]
- None[T]

## Installation
`go get github.com/iUwej/go-opt`

## Usage
1) To create an option with a value present, use the Some function.
```go
package main

import "github.com/iUwej/go-opt"

intOpt := opt.Some(5) // creates an int option
strOpt := opt.Some("some string")
custOpt := opt.Some(CustomType())
```
2) To create an option with a value absent, use the None function.
```go
package main

import "github.com/iUwej/go-opt"

intOpt := opt.None[int]()
strOpt := opt.None[string]()
custOpt := opt.None[CustomType]()
```

## Example
```go
package main

import (
	"fmt"
	"github.com/iUwej/go-opt"
	"math/rand"
	"time"
)

type User struct {
	Id   string
	Name string
}

func FindUser(id string) opt.Option[User] {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	val := r1.Intn(100)
	if val%2 == 0 {
		user := User{
			Id:   id,
			Name: "Some Name",
		}

		return opt.Some(user)
	}
	return opt.None[User]()

}

func main() {

	userOpt := FindUser("some_id")
	if userOpt.Empty() {
		// do something about the missing user
	} else {
		user := userOpt.Get()
		fmt.Println("We definitely have a user ", user)
	}

	// we could provide a default user
	user := userOpt.GetOrElse(User{"default_id", "Default Name"})
	fmt.Println(user)

	// we can get the user if present or get zero value
	user = userOpt.GetOrZero()
	fmt.Println(user)

	//we can map the option to another type
	nameOpt := opt.Map(func(t User) string {
		return t.Name
	}, userOpt)
	fmt.Println(nameOpt.GetOrElse("default_name"))

}

```

