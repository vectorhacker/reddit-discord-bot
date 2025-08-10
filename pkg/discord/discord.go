package discord

import (
	"errors"

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
		err := d.threader.CreateThread(d.channelId, post.Title, post.Text)
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
