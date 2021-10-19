package handlers

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	indexContent := `
<!DOCTYPE html>
<html>
  <head>
    <title>Homepage - Project Name</title>
    <link rel="stylesheet" href="static/css/styles.css" />
    <script async src="static/js/index.js"></script>
  </head>
  <body>
    <h1>Echorand Corp. This is the homepage for project Project Name.</h1>
  </body>
</html>
`
	fmt.Fprintf(w, indexContent)
}
