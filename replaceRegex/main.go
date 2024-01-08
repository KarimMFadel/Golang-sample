package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
)

var (
	htmlcontent = "<a href=\"@@@emailPermalink@@@\" style=\"color: #2E6E9E;\" target=\"_blank\">click here</a>"
)

func main() {

	secondMethod()

	url := "asdasdasdwerwerwerwerwe  $RANDOMr"

	replaceRegex := regexp.MustCompile(`\$RANDOM`)
	randString := strconv.FormatInt(int64(rand.Float64()*10e12), 10)

	fmt.Println(replaceRegex.ReplaceAllString(url, randString))

	convertTemplates(htmlcontent)
}

func convertTemplates(htmlcontent string) {
	fmt.Println("htmlcontent Before updated:", htmlcontent)

	replaceRegex := regexp.MustCompile(`@@@(.*)@@@`)
	replaceRegexSec := regexp.MustCompile(`{{(.*)}}`)
	randString := `%recipient.$1%`

	result := replaceRegex.ReplaceAllString(htmlcontent, randString)
	result = replaceRegexSec.ReplaceAllString(result, randString)
	fmt.Println("htmlcontent Before updated:", result)
}

func secondMethod() {

	example := "\\\"#GoLangCode!$!\\\""

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("\\\"")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(example, "")

	fmt.Printf("%s \n  %s \n", example, processedString)
}
