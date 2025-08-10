package reddit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ScrapeSorting string

const (
	Best   ScrapeSorting = "best"
	Hot    ScrapeSorting = "hot"
	New    ScrapeSorting = "new"
	Top    ScrapeSorting = "top"
	Rising ScrapeSorting = "rising"
)

type Post struct {
	Title       string
	URL         string
	Ups         int32
	Downs       int32
	UpvoteRatio float32
	Text        string
	Preview     ScrapePreview
}

type ScrapePreview struct {
	Images []struct {
		Url string
		Id  string
	}
}

type ScrapeSubredditResult struct {
	Posts  []Post
	After  string
	Before string
}

type ScrapeSubredditFunc func(
	sub string,
	sorting ScrapeSorting,
) (ScrapeSubredditResult, error)

func ScrapeSubreddit(
	sub string,
	sorting ScrapeSorting,
) (*ScrapeSubredditResult, error) {

	result := ScrapeSubredditResult{}

	response, err := http.Get(fmt.Sprintf("https://www.reddit.com/r/%s/%s/.json", sub, sorting))

	if err != nil {
		return nil, fmt.Errorf("unable to retrieve reddit page: %w", err)
	}
	defer response.Body.Close()

	data := RedditResponse{}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("problem unmarshalling body: %w", err)
	}

	for _, child := range data.Data.Children {
		preview := ScrapePreview{}

		if child.Data.Preview != nil {
			for _, image := range child.Data.Preview.Images {
				preview.Images = append(preview.Images, struct {
					Url string
					Id  string
				}{

					Url: image.Source.URL,
					Id:  image.ID,
				})
			}
		}

		result.Posts = append(result.Posts, Post{
			Title:       child.Data.Title,
			Downs:       int32(child.Data.Downs),
			URL:         child.Data.URL,
			UpvoteRatio: float32(child.Data.UpvoteRatio),
			Ups:         int32(child.Data.Ups),
			Text:        child.Data.Selftext,
			Preview:     preview,
		})
	}

	return &result, nil
}
