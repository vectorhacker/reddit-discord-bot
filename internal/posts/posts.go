package posts

import (
	"github.com/vectorhacker/reddit-discord-bot/pkg/discord"
	"github.com/vectorhacker/reddit-discord-bot/pkg/reddit"
)

type Posts struct {
	d         *discord.DiscordInteraction
	subreddit string
	sorting   reddit.ScrapeSorting
}

func New(d *discord.DiscordInteraction, subreddit string, sorting reddit.ScrapeSorting) *Posts {
	return &Posts{
		d,
		subreddit,
		sorting,
	}
}

func (p *Posts) Run() error {
	scrape, err := reddit.ScrapeSubreddit(p.subreddit, p.sorting)
	if err != nil {
		return err
	}

	return p.d.ThreadRedditPosts(scrape.Posts)
}
