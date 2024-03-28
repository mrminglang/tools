package htmls

import (
	"golang.org/x/net/html"
	"strings"
)

// HTML富文本转纯文本
func ConvertToText(htmlContent string) string {
	if htmlContent == "" {
		return htmlContent
	}

	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return ""
	}

	var textContent string
	var extractText func(*html.Node)
	extractText = func(n *html.Node) {
		if n.Type == html.TextNode {
			textContent += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractText(c)
		}
	}

	extractText(doc)
	// Remove white spaces and new lines
	textContent = strings.ReplaceAll(textContent, "\n", "")
	textContent = strings.ReplaceAll(textContent, "\t", "")
	textContent = strings.ReplaceAll(textContent, " ", "")

	return textContent
}
