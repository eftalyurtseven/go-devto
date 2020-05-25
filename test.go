package devtoclient

import "fmt"

type Post struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	MinSize string `json:"min_size"`
}

func (c *Client) GetPosts() ([]Post, error) {
	var posts []Post
	url := fmt.Sprintf("/articles")
	_, err := c.Request("GET", url, nil, &posts)
	return posts, err
}
