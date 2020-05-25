package devtoclient

import "fmt"

// User struct for server returns.
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

// Follower struct for server returns.
type Follower struct {
	TypeOf       string `json:"type_of"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image"`
}

// GetUser method serves get a user with parameters from API
// example context: "me" || "1"
func (c *Client) GetUser(ctx string) (User, error) {
	var user User
	url := fmt.Sprintf("/users/%s", ctx)
	fmt.Println(url)
	_, err := c.Request("GET", url, nil, &user)
	return user, err
}

// GetFollowers method servers get auth user followers from API
func (c *Client) GetFollowers() ([]Follower, error) {
	var followers []Follower
	url := fmt.Sprintf("/followers/users")
	_, err := c.Request("GET", url, nil, &followers)
	return followers, err
}
