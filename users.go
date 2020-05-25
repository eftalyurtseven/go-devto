package devtoclient

import "fmt"

type User struct {
	TypeOf          string      `json:"type_of"`
	ID              int         `json:"id"`
	Username        string      `json:"username"`
	Name            string      `json:"name"`
	Summary         string      `json:"summary"`
	TwitterUsername string      `json:"twitter_username"`
	GithubUsername  string      `json:"github_username"`
	WebsiteURL      interface{} `json:"website_url"`
	Location        string      `json:"location"`
	JoinedAt        string      `json:"joined_at"`
	ProfileImage    string      `json:"profile_image"`
}

type Follower struct {
	TypeOf       string `json:"type_of"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image"`
}

func (c *Client) GetUser(ctx string) (User, error) {
	var user User
	url := fmt.Sprintf("/users/&s", ctx)
	_, err := c.Request("GET", url, nil, &user)
	return user, err
}

func (c *Client) GetFollowers() ([]Follower, error) {
	var followers []Follower
	url := fmt.Sprintf("/followers/users")
	_, err := c.Request("GET", url, nil, &followers)
	return followers, err
}
