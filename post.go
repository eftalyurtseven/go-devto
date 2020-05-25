package devtoclient

import "fmt"

type Post struct {
	type_of                  string `json:"type_of"`
	id                       int    `json:"id"`
	title                    string `json:"title"`
	description              string `json:"description"`
	readable_publish_date    string `json:"readable_publish_date"`
	slug                     string `json:"slug"`
	path                     string `json:"path"`
	url                      string `json:"url"`
	comments_count           int    `json:"comments_count"`
	positive_reactions_count int    `json:"positive_reactions_count"`
}

func (c *Client) GetPosts() ([]Post, error) {
	var posts []Post
	url := fmt.Sprintf("/articles")
	_, err := c.Request("GET", url, nil, &posts)
	return posts, err
}
