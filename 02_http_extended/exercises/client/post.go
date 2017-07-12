package main

import "net/http"


type User struct {
	Name string

}
func main() {

	http.Post("localhost:10000", "application/json", )
}
