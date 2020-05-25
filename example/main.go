package main

import (
	"fmt"

	devtoclient "github.com/eftalyurtseven/go-devto"
)

func main() {
	client := devtoclient.NewClient()
	// TODO update configs
	client.UpdateConfig(&devtoclient.ClientConfig{
		BaseURL: "https://dev.to/api",
		Key:     "YOUR_API_KEY",
	})
	// get articles //
	articles, err := client.GetArticles("") // you can call with parameters ex client.GetArticles("?page=1&per_page=10")
	if err != nil {
		fmt.Println(err)
	}
	for _, article := range articles {
		fmt.Println(article.Title)
	}
	// get specific article //
	article, err := client.GetArticle(343488)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Single Article Title->", article.Title)

	// add an article //
	data := devtoclient.Article{
		Title:        "Hello from  go-devto API Client",
		BodyMarkdown: "Hey, this post published by go-devto API Client cheers!",
		Published:    true,
		Tags:         []string{"golang", "devto-client", "restful-api"},
	}
	added, err := client.AddArticle(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(added.Title)

	// update an article //
	updatedData := devtoclient.Article{
		Title: "Hello from  go-devto API Client Updated!",
	}
	updated, err := client.UpdateArticle(updatedData, 343485)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Updated!", updated.Title)
}
