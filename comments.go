package devtoclient

import "fmt"

// Comment struct for API returns
type Comment struct {
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
	Children []struct {
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
	} `json:"children"`
}

// GetComments method serve to get comments for an article
func (c *Client) GetComments(articleID int) ([]Comment, error) {
	var comments []Comment
	url := fmt.Sprintf("/comments?a_id=%d", articleID)
	_, err := c.Request("GET", url, nil, &comments)
	return comments, err
}

// GetComment method serve to get comment for an article
func (c *Client) GetComment(commentID string) (Comment, error) {
	var comment Comment
	url := fmt.Sprintf("/comments/%s", commentID)
	_, err := c.Request("GET", url, nil, &comment)
	return comment, err
}
