package devtoclient

import (
	"fmt"
	"time"
)

type Post struct {
	TypeOf                 string      `json:"type_of"`
	ID                     int         `json:"id"`
	Title                  string      `json:"title"`
	Description            string      `json:"description"`
	ReadablePublishDate    string      `json:"readable_publish_date"`
	Slug                   string      `json:"slug"`
	Path                   string      `json:"path"`
	URL                    string      `json:"url"`
	CommentsCount          int         `json:"comments_count"`
	PositiveReactionsCount int         `json:"positive_reactions_count"`
	CollectionID           interface{} `json:"collection_id"`
	PublishedTimestamp     time.Time   `json:"published_timestamp"`
	CoverImage             string      `json:"cover_image"`
	SocialImage            string      `json:"social_image"`
	CanonicalURL           string      `json:"canonical_url"`
	CreatedAt              time.Time   `json:"created_at"`
	EditedAt               interface{} `json:"edited_at"`
	CrosspostedAt          interface{} `json:"crossposted_at"`
	PublishedAt            time.Time   `json:"published_at"`
	LastCommentAt          time.Time   `json:"last_comment_at"`
	TagList                []string    `json:"tag_list"`
	Tags                   string      `json:"tags"`
	User                   struct {
		Name            string      `json:"name"`
		Username        string      `json:"username"`
		TwitterUsername string      `json:"twitter_username"`
		GithubUsername  string      `json:"github_username"`
		WebsiteURL      interface{} `json:"website_url"`
		ProfileImage    string      `json:"profile_image"`
		ProfileImage90  string      `json:"profile_image_90"`
	} `json:"user"`
}

func (c *Client) GetPosts() ([]Post, error) {
	var posts []Post
	url := fmt.Sprintf("/articles")
	_, err := c.Request("GET", url, nil, &posts)
	return posts, err
}
