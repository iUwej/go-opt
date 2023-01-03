# go-opt
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

import "github.com/iUwej/go-opt"



type User struct {
	Id     int64
	Name   string
	Emails []string
}

func FindUser(id string)opt.Option[User]  {
	user := User{
		Id : id,
		Name:   "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	if true{
		return opt.Some(user)
    } else{
		return opt.None[User]()
    }
	
}

func main() {
    userOpt := FindUser("some_id")
	// check if user present
	if userOpt.Empty(){
		//do something if userOpt is empty
    }
	// get the user if present or provide a default
	user := userOpt.GetOrElse(User{id:"back_up",Name: "Some Name"})
	
	//map the optional user into an optional name
	nameOpt := opt.Map(func(t User) string {
        return t.Name
	},userOpt)
}

```

