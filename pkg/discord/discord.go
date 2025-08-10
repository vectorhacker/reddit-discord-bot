package discord

import (
	"errors"
	"fmt"
	"strings"

	"github.com/vectorhacker/reddit-discord-bot/pkg/reddit"
)

type DiscordInteraction struct {
	threader  DiscordThreader
	channelId string
}

func NewInteraction(threader DiscordThreader, channelId string) *DiscordInteraction {
	return &DiscordInteraction{
		threader:  threader,
		channelId: channelId,
	}
}

func (d DiscordInteraction) ThreadRedditPosts(posts []reddit.Post) error {

	errs := []error{}
	for _, post := range posts {
		content := truncatePost(post)

		message := fmt.Sprintf(
			"%s\n\n%s\n\n%s",
			post.Title,
			content,
			post.URL,
		)
		err := d.threader.CreateThread(d.channelId, post.Title, message)
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func truncatePost(post reddit.Post) string {
	content := post.Text[:min(len(post.Text), 1000)]
	content = strings.TrimSpace(content)
	if len(content) < len(post.Text) {
		content += "..."
	}
	return content
}
