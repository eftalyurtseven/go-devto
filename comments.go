package devtoclient

import "fmt"

type Comment []struct {
	TypeOf   string `json:"type_of"`
	IDCode   string `json:"id_code"`
	BodyHTML string `json:"body_html"`
	User     struct {
		Name            string      `json:"name"`
		Username        string      `json:"username"`
		TwitterUsername interface{} `json:"twitter_username"`
		GithubUsername  interface{} `json:"github_username"`
		WebsiteURL      interface{} `json:"website_url"`
		ProfileImage    string      `json:"profile_image"`
		ProfileImage90  string      `json:"profile_image_90"`
	} `json:"user"`
	Children []interface{} `json:"children"`
}

func (c *Client) getComments(articleId int) ([]Comment, error) {
	var comments []Comment
	url := fmt.Sprintf("/comments?a_id=%d", articleId)
	_, err := c.Request("GET", url, nil, &comments)
	return comments, err
}

func (c *Client) getComment(commentId int) (Comment, error) {
	var comment Comment
	url := fmt.Sprintf("/comments/%d", commentId)
	_, err := c.Request("GET", url, nil, &comment)
	return comment, err
}
