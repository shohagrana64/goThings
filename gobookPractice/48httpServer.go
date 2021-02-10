package main

import (
	"io"
	"net/http"
)

func hello4(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
<html>
  <head>
      <title>Hello World</title>
  </head>
  <body>
      Whatsup Prithibi!
  </body>
</html>`,
	)
}
func main() {
	http.HandleFunc("/hello", hello4)
	http.ListenAndServe(":9000", nil)
}
