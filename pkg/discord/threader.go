package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordThreader interface {
	CreateThread(channel string, title string, message string) (err error)
}

type DiscordGoThreader struct {
	session *discordgo.Session
}

func (d *DiscordGoThreader) CreateThread(
	channel string,
	title string,
	message string,
) error {

	discordMessage, err := d.session.ChannelMessageSend(channel, message)
	if err != nil {
		return fmt.Errorf("unable to create message on channel %s due to: %v", channel, err)
	}

	_, err = d.session.MessageThreadStart(channel, discordMessage.ID, title, 24*60)
	if err != nil {
		return fmt.Errorf("failed to create thread using message id: %s due to: %v", discordMessage.ID, err)
	}

	return nil
}
