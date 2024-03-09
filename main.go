package main


import (
   "stuService/router"
   "fmt"
)

func main(){
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("error:", r)
        }
    }()
    r := router.Router()
    r.Run(":8080") 
}
