package discord_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vectorhacker/reddit-discord-bot/pkg/discord"
	"github.com/vectorhacker/reddit-discord-bot/pkg/reddit"
)

type fakeThreader struct {
	t         *testing.T
	calls     int
	channelId string
}

// CreateThread implements discord.DiscordThreader.
func (f *fakeThreader) CreateThread(
	channel string,
	title string,
	message string,
) (err error) {
	f.calls += 1

	assert.Equal(f.t, f.channelId, channel)
	assert.NotEmpty(f.t, title)
	assert.NotEmpty(f.t, message)

	return nil
}

func TestThreadingPosts(t *testing.T) {
	channelId := "test"

	threader := &fakeThreader{
		t:         t,
		channelId: channelId,
	}
	scrape, _ := reddit.ScrapeSubreddit("omscs", reddit.Best)

	interaction := discord.NewInteraction(threader, channelId)

	err := interaction.ThreadRedditPosts(scrape.Posts)
	assert.Nil(t, err)
	assert.Equal(t, len(scrape.Posts), threader.calls)
}
