package discord

import (
	"errors"
	"fmt"

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
		message := fmt.Sprintf(
			"%s\n\n%s\n\n%s",
			post.Title,
			post.Text[:min(len(post.Text), 1000)],
			post.URL,
		)
		err := d.threader.CreateThread(d.channelId, post.Title, message)
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
