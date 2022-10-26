package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

var feedbackHtml string = `
<section class="feedback-section">
    <div class="container">
        <div class="row">
            <div class="col-md-12 text-center">
                <h1 class="single-line">The Customers’ Verdict</h1>
            </div>
        </div><br>
        <div class="row">
            <div class="col-md-3">
                <div class="feedback-head">
                    <a href="https://www.google.com/search?q=THE+MARINA+MALL&rlz=1C1CHBD_enIN908IN908&oq=th+&aqs=chrome.1.69i57j69i59l2j69i60l5.2415j0j7&sourceid=chrome&ie=UTF-8#lrd=0x3a525a5ed3d3509d:0x51ba8d5c2f099ebb,1,,," target="_blank">
                        <img src="http://marinamallchennai.com/wp-content/uploads/2020/09/Googlebusiness.png">
                    </a>
                    <p><i class="fa fa-star"></i> <i class="fa fa-star"></i> <i class="fa fa-star"></i> <i class="fa fa-star"></i> <i class="fa fa-star"></i></p><br>
                </div>
            </div>
      
        </div>
       
    </div>
</section>
`

func Query(n *html.Node, query string) *html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return &html.Node{}
	}
	return cascadia.Query(n, sel)
}

func QueryAll(n *html.Node, query string) []*html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return []*html.Node{}
	}
	return cascadia.QueryAll(n, sel)
}

func AttrOr(n *html.Node, attrName, or string) string {
	for _, a := range n.Attr {

		if a.Key == attrName {
			return a.Val
		}
	}
	return or
}

// main function
func main() {
	doc, err := html.Parse(strings.NewReader(feedbackHtml))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of URLS:\n\n")
	for _, p := range QueryAll(doc, "section.feedback-section") {
		reviewUrl := AttrOr(Query(p, "div a "), "href", "")
		imageUrl := AttrOr(Query(p, "div  a img"), "src", "")
		fmt.Println("Review url", reviewUrl, "\n", "Image URl", imageUrl)
	}
}
