package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	gexast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type ExtractedData struct {
	Questions []Question `json:"questions"`
}

type Answer struct {
	Content   string `json:"content"`
	IsCorrect bool   `json:"isCorrect"`
}

type Question struct {
	Content     string   `json:"content"`
	Answers     []Answer `json:"answers"`
	Explanation string   `json:"explanation"`
}

func extractData(markdownContent []byte) (*ExtractedData, error) {
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithExtensions(
			extension.TaskList,
		),
	)

	reader := text.NewReader(markdownContent)
	parsed := md.Parser().Parse(reader)
	// parsed.Dump(markdownContent, 2)

	// Extract data from the parsed Markdown AST (Abstract Syntax Tree)
	data := ExtractedData{
		Questions: make([]Question, 0),
	}

	cnt := -1

	walker := func(node gast.Node, entering bool) (gast.WalkStatus, error) {
		if !entering {
			return gast.WalkContinue, nil
		}

		// Check the type of the node and perform the necessary extraction logic
		switch n := node.(type) {
		case *gast.Heading:
			// Extract heading data
			title := string(n.Text(markdownContent))
			if n.Level < 4 && title != "" {
				return gast.WalkContinue, nil
			}
			q := Question{
				Content: title,
				Answers: make([]Answer, 0),
			}
			fmt.Println("\n===Starting new question: ", title)
			data.Questions = append(data.Questions, q)
			cnt++

		case *gast.TextBlock:
			if n.ChildCount() <= 1 {
				return gast.WalkContinue, nil
			}
			ch1 := n.FirstChild()
			ch2 := n.FirstChild().NextSibling()

			checkBox, ok1 := ch1.(*gexast.TaskCheckBox)
			text, ok2 := ch2.(*gast.Text)
			if !ok1 || !ok2 {
				return gast.WalkContinue, nil
			}

			// fmt.Printf("TextBlock %d children. Checked: %v, Text: %s\n", n.ChildCount(), checkBox.IsChecked, text.Text(markdownContent))

			q := &data.Questions[cnt]
			q.Answers = append(q.Answers, Answer{
				Content:   string(text.Text(markdownContent)),
				IsCorrect: checkBox.IsChecked,
			})
		case *gast.Paragraph:
			// Extract paragraph data
			text := string(n.Text(markdownContent))
			fmt.Println("Paragraph text: ", text)

			list, ok := n.NextSibling().(*gast.List)
			if !ok {
				return gast.WalkContinue, nil
			}

			fmt.Println("List: ", list.ChildCount())

		default:
			// Handle other node types if necessary
		}

		return gast.WalkContinue, nil
	}

	if err := gast.Walk(parsed, walker); err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	source, err := os.ReadFile("./static/aws-lambda-quiz-full.md")
	log.Printf("Read %d bytes", len(source))
	if err != nil {
		log.Panic(err)
	}

	// Extract data from Markdown
	extractedData, err := extractData(source)
	if err != nil {
		fmt.Println("Error extracting data from Markdown:", err)
		return
	}

	// Convert data to JSON
	jsonData, err := json.MarshalIndent(extractedData, "", "  ")
	if err != nil {
		fmt.Println("Error encoding data to JSON:", err)
		return
	}

	// jsonData = []byte{}
	fmt.Println(string(jsonData))

	// Save JSON data to a file
	jsonFile := "output.json"
	err = os.WriteFile(jsonFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error saving JSON file:", err)
		return
	}

	fmt.Println("Data extracted and saved to", jsonFile)
}
