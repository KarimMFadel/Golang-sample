package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

var (
	htmlBody  = "<body><a href=\"{{{yyyyy}}}\" style=\"color: #2E6E9E; text-decoration: none;\">click here</a><p width=\"500\" >Germany the recent <a href=\"{{{xxxx}}}?testId=1234\" style=\"color: #2E6E9E; text-decoration: none;\">click here</a></p></body>"
	body      = "<body><p width=\"500\" >Germany <img width=\"500\" type=\"http://www.google.com/Article\" url=\"http://www.google.com/b8ec6435630\">Mario Draghi </img>, the recent </p></body>"
	paragraph = "font-family:arial; "
	pTag      = paragraph + "text-align: center;"
)

func main() {
	insertId("1111111111", htmlBody)
	fmt.Println("****************************************************")
	//bodyParse()
}

func bodyParse() {
	fmt.Println("Body Before updated:", body)
	fmt.Println("==================================================")

	root, err := html.Parse(strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	var updateBody func(*html.Node)
	updateBody = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "p" {
			n.Attr = append(n.Attr, html.Attribute{Key: "style", Val: paragraph})
		}

		if n.Type == html.ElementNode && n.Data == "img" {
			var attrs []html.Attribute
			changed := false
			for _, a := range n.Attr {
				if a.Key == "width" {
					val, _ := strconv.Atoi(a.Val)
					if val > 350 {
						changed = true
						attrs = append(attrs, html.Attribute{
							Key: "width",
							Val: strconv.Itoa(350),
						})
					}
				} else {
					attrs = append(attrs, a)
				}
			}
			if changed {
				n.Attr = attrs
			}
			newPTag := &html.Node{
				Type: html.ElementNode,
				Data: "p",
				Attr: []html.Attribute{{Key: "style", Val: pTag}},
			}
			elem := n.Parent
			elem.RemoveChild(n)
			newPTag.AppendChild(n)
			elem.AppendChild(newPTag)
		}

		// traverses the HTML of the webpage from the first child node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			updateBody(c)
		}
	}
	updateBody(root)

	result, errRend := renderHtml(root)
	if errRend != nil {
		fmt.Println("Error", errRend)
	}

	fmt.Println("Body after updated:", result)
}

func resposeParse() {
	// Send an HTTP GET request to the example.com web page
	resp, err := http.Get("https://www.example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Use the html package to parse the response body from the request
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Find and print all links on the web page
	var links []string
	var link func(*html.Node)
	link = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "p" {
			n.Attr = []html.Attribute{{Key: "class", Val: n.Data}}
		}

		// traverses the HTML of the webpage from the first child node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			link(c)
		}
	}
	link(doc)

	// loops through the links slice
	for _, l := range links {
		fmt.Println("Link:", l)
	}

	var b bytes.Buffer
	errRend := html.Render(&b, doc)
	if errRend != nil {
		fmt.Println("Error", errRend)
	}

	fmt.Println("doc:", b.String())
}

func insertId(sId string, htmlString string) error {
	fmt.Println("htmlString Before updated:", htmlString)
	fmt.Println("==================================================")

	htmlPage, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return err
	}

	err = updateHtml(htmlPage, sId)
	if err != nil {
		return err
	}

	result, errRend := renderHtml(htmlPage)
	if errRend != nil {
		return errRend
	}

	fmt.Println("htmlString after updated:", result)
	return nil
}

func updateHtml(n *html.Node, segmentId string) error {
	if n.Type == html.ElementNode && n.Data == "body" {
		updateHref(n, segmentId)
	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		updateHtml(c, segmentId)
	}

	return nil
}

func updateHref(n *html.Node, segmentId string) error {
	if n.Type == html.ElementNode && n.Data == "a" {
		var attrs []html.Attribute
		changed := false
		for _, a := range n.Attr {
			if a.Key == "href" {
				hrefVal := a.Val
				if !StringIsEmpty(segmentId) {
					var err error
					hrefVal, err = addQueryString(a.Val, "segmentId="+segmentId)
					if err != nil {
						return err
					}
				}
				decodeURI, _ := url.PathUnescape(hrefVal)
				attrs = append(attrs, html.Attribute{
					Key: "href",
					Val: decodeURI,
				})
				changed = true
				break
			} else {
				attrs = append(attrs, a)
			}
		}
		if changed {
			n.Attr = attrs
		}
	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		updateHref(c, segmentId)
	}

	return nil
}

func addQueryString(href, query string) (string, error) {
	link, err := url.Parse(href)
	if err != nil {
		return "", err
	}
	// fmt.Println("link.RawQuery : ", link.RawQuery)
	if StringIsEmpty(link.RawQuery) {
		link.RawQuery = query
	} else {
		link.RawQuery = link.RawQuery + "&" + query
	}
	return link.String(), nil
}

func StringIsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func renderHtml(root *html.Node) (string, error) {
	var b bytes.Buffer
	err := html.Render(&b, root)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
