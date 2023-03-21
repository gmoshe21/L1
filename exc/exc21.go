package main

import (
	"fmt"
)

type Target interface {
    Request() string
}

type Adaptee interface {
    SpecificRequest() string
}

type adapteeImpl struct {}

func (a *adapteeImpl) SpecificRequest() string {
    return "Adaptee request"
}

type adapter struct {
    adaptee Adaptee
}

func (a *adapter) Request() string {
    return "Adapter: " + a.adaptee.SpecificRequest()
}

func main() {
    adaptee := &adapteeImpl{}
    target := &adapter{
        adaptee: adaptee,
    }

    fmt.Println(target.Request())
}