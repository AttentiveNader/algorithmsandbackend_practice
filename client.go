package client

import (
	"net/http"
	"fmt"
)

func main() {
	client := &http.Client{}
	fmt.Println(client)	
}