package action

import (
	"fmt"
	"net/http"
)

func IndexAction(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Hello world!")
}
