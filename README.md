# Go Dev. to API Client

## Summary

This repo allows easy operation using the [dev. to](https://dev.to) APIs. 

## Installation, 

If using Go modules (Go version >= 11. 1) simply import as needed. 
<pre>
go mod init github. com/yourusername/yourreponame
</pre>

### Older Go versions

<pre>
go get github. com/eftalyurtseven/go-devto
</pre>

## Documentation

### Setup

Client will respect enviroment variables: BaseURL, Key

``` go

    import (
        devtoclient "github. com/eftalyurtseven/go-devto"
    )

    client := devtoclient.NewClient()
	// TODO update configs
	client.UpdateConfig(&devtoclient.ClientConfig{
		BaseURL: "https://dev.to/api",
		Key:     "YOUR_API_KEY",
	})

```

## Methods & Returns

| **Method**    | **Paramaters - Type**              | **Returns**       |
|---------------|------------------------------------|-------------------|
| GetArticles   | ctx - string                       | []Article, error  |
| GetArticle    | postID - int                       | Article, error    |
| NewArticle    | newPost - Article                  | Article, error    |
| UpdateArticle | updated - Article, articleID - int | Article, error    |
| GetComments   | articleID - int                    | []Comment, error  |
| GetComment    | commentID - string                 | Comment, error    |
| GetUser       | ctx - string                       | User, error       |
| GetFollowers  | -                                  | []Follower, error |

## Examples

### GetArticles
GetArticles method serves to get articles from API

``` go
    // get articles //
	articles, err := client.GetArticles("") // you can call with parameters ex client.GetArticles("?page=1&per_page=10")
	if err != nil {
		fmt.Println(err)
	}
	for _, article := range articles {
		fmt.Println(article.Title)
	}
```

### GetArticle

``` go
	article, err := client.GetArticle(343488)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Single Article Title->", article.Title)
```

### AddArticle

``` go
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
```

### UpdateArticle

``` go
	updatedData := devtoclient.Article{
		Title: "Hello from  go-devto API Client Updated!",
	}
	updated, err := client.UpdateArticle(updatedData, 343485)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Updated!", updated.Title)
```

### GetUser

``` go
    user, err := client. GetUser("me") // or you can use client. GetUser(userID string)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
```

### GetFollowers

``` go
followers, err := client. GetFollowers()
	if err != nil {
		panic(err)
	}
	for _, follower := range followers {
		fmt.Println("Follower Name -> ", follower.Name)
	}
```

### GetComments

``` go
	comments, err := client.GetComments(270180) // ArticleID
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println("Comment -> ", comment.BodyHTML)
	}
```

### GetComment

``` go
    comment, err := client. GetComment("m51e")
	if err != nil {
		panic(err)
	}
	fmt.Println("Single comment -> ", comment.BodyHTML)
```

### Data types & structures

#### Post

``` go
type Post struct {
	TypeOf                 string `json:"type_of"` 
	ID                     int `json:"id"` 
	Title                  string `json:"title"` 
	Description            string `json:"description"` 
	ReadablePublishDate    string `json:"readable_publish_date"` 
	Slug                   string `json:"slug"` 
	Path                   string `json:"path"` 
	URL                    string `json:"url"` 
	CommentsCount          int `json:"comments_count"` 
	PositiveReactionsCount int `json:"positive_reactions_count"` 
	CollectionID           interface{} `json:"collection_id"` 
	PublishedTimestamp     time.Time `json:"published_timestamp"` 
	CoverImage             string `json:"cover_image"` 
	SocialImage            string `json:"social_image"` 
	CanonicalURL           string `json:"canonical_url"` 
	CreatedAt              time.Time `json:"created_at"` 
	EditedAt               interface{} `json:"edited_at"` 
	CrosspostedAt          interface{} `json:"crossposted_at"` 
	PublishedAt            time.Time `json:"published_at"` 
	LastCommentAt          time.Time `json:"last_comment_at"` 
	TagList                []string `json:"tag_list"` 
	Tags                   string `json:"tags"` 
	User                   struct {
		Name            string `json:"name"` 
		Username        string `json:"username"` 
		TwitterUsername string `json:"twitter_username"` 
		GithubUsername  string `json:"github_username"` 
		WebsiteURL      interface{} `json:"website_url"` 
		ProfileImage    string `json:"profile_image"` 
		ProfileImage90  string `json:"profile_image_90"` 
	} `json:"user"` 
}
```

#### Article

``` go
// Article struct for add article and returns.
type Article struct {
	Title        string `json:"title"` 
	BodyMarkdown string `json:"body_markdown"` 
	Published    bool `json:"published"` 
	Tags         []string `json:"tags"` 
}
```

#### Comment

``` go
// Comment struct for API returns
type Comment struct {
	TypeOf   string `json:"type_of"` 
	IDCode   string `json:"id_code"` 
	BodyHTML string `json:"body_html"` 
	User     struct {
		Name            string `json:"name"` 
		Username        string `json:"username"` 
		TwitterUsername interface{} `json:"twitter_username"` 
		GithubUsername  interface{} `json:"github_username"` 
		WebsiteURL      interface{} `json:"website_url"` 
		ProfileImage    string `json:"profile_image"` 
		ProfileImage90  string `json:"profile_image_90"` 
	} `json:"user"` 
	Children []struct {
		TypeOf   string `json:"type_of"` 
		IDCode   string `json:"id_code"` 
		BodyHTML string `json:"body_html"` 
		User     struct {
			Name            string `json:"name"` 
			Username        string `json:"username"` 
			TwitterUsername interface{} `json:"twitter_username"` 
			GithubUsername  interface{} `json:"github_username"` 
			WebsiteURL      interface{} `json:"website_url"` 
			ProfileImage    string `json:"profile_image"` 
			ProfileImage90  string `json:"profile_image_90"` 
		} `json:"user"` 
		Children []interface{} `json:"children"` 
	} `json:"children"` 
}
```

#### User

``` go
// User struct for server returns.
type User struct {
	TypeOf          string `json:"type_of"` 
	ID              int `json:"id"` 
	Username        string `json:"username"` 
	Name            string `json:"name"` 
	Summary         string `json:"summary"` 
	TwitterUsername string `json:"twitter_username"` 
	GithubUsername  string `json:"github_username"` 
	WebsiteURL      interface{} `json:"website_url"` 
	Location        string `json:"location"` 
	JoinedAt        string `json:"joined_at"` 
	ProfileImage    string `json:"profile_image"` 
}
```

#### Follower

``` go
// Follower struct for server returns.
type Follower struct {
	TypeOf       string `json:"type_of"` 
	ID           int `json:"id"` 
	Name         string `json:"name"` 
	Path         string `json:"path"` 
	Username     string `json:"username"` 
	ProfileImage string `json:"profile_image"` 
}
```
