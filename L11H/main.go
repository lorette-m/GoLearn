package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	//"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"os"
	"strings"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`

func visit(n *html.Node, words, pics *int) {

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))

	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)

	return words, pics
}

func main() {
	/*c := colly.NewCollector(
	colly.AllowedDomains("https://en.wikipedia.org/wiki/Ice_eggs"))*/

	//url := "https://en.wikipedia.org/wiki/Ice_eggs"
	url := "https://go-colly.org/"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	// Read the response body
	bytess, _ := io.ReadAll(resp.Body)

	// Print the body as a string
	//fmt.Println("HTML:\n\n", string(bytess))

	//doc, err := html.Parse(bytes.NewReader([]byte(raw)))
	doc, err := html.Parse(bytes.NewReader(bytess))

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)

	fmt.Printf("words: %d, pics: %d\n", words, pics)
}
