package reddit_test

import (
	"strings"
	"testing"

	"github.com/vectorhacker/reddit-discord-bot/pkg/reddit"
)

func TestScrapeReddit(t *testing.T) {
	tests := []struct {
		sorting reddit.ScrapeSorting
	}{
		{reddit.Best},
		{reddit.Hot},
		{reddit.New},
		{reddit.Rising},
		{reddit.Top},
	}

	for _, test := range tests {
		t.Run(string(test.sorting), func(t *testing.T) {
			t.Parallel()
			result, err := reddit.ScrapeSubreddit("artificial", test.sorting)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if len(result.Posts) == 0 {
				t.Fatalf("Got no posts")
			}

			for _, post := range result.Posts {
				if strings.TrimSpace(post.Title) == "" {
					t.Fatalf("Expected title to be populated")
				}
			}
		})
	}
}
