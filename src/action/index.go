package action

import (
	"fmt"
	"net/http"
)

func Index(writer http.ResponseWriter, r http.Request) {
	fmt.Println("API REST EXAMPLE")
}
