package main

import (
	"fmt"

	"jaytaylor.com/html2text"
)

func main() {
	inputHTML := `
<html>
  <head>
    <title>My Mega Service</title>
    <link rel=\"stylesheet\" href=\"main.css\">
    <style type=\"text/css\">body { color: #fff; }</style>
  </head>

  <body>
    <div class="logo">
      <a href="http://jaytaylor.com/"><img src="/logo-image.jpg" alt="Mega Service"/></a>
    </div>

    <h1>Welcome to your new account on my service!</h1>

    <p>
      Here is some more information:

      <ul>
        <li>Link 1: <a href="https://example.com">Example.com</a></li>
        <li>Link 2: <a href="https://example2.com">Example2.com</a></li>
        <li>Something else</li>
      </ul>
    </p>

    <table>
      <thead>
        <tr><th>Header 1</th><th>Header 2</th></tr>
      </thead>
      <tfoot>
        <tr><td>Footer 1</td><td>Footer 2</td></tr>
      </tfoot>
      <tbody>
        <tr><td>Row 1 Col 1</td><td>Row 1 Col 2</td></tr>
        <tr><td>Row 2 Col 1</td><td>Row 2 Col 2</td></tr>
      </tbody>
    </table>
  </body>
</html>`

	text, err := html2text.FromString(inputHTML, html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}
