package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/vectorhacker/reddit-discord-bot/internal/posts"
	"github.com/vectorhacker/reddit-discord-bot/pkg/discord"
	"github.com/vectorhacker/reddit-discord-bot/pkg/reddit"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	token := os.Getenv("DISCORD_TOKEN")
	channelId := os.Getenv("DISCORD_CHANNEL_ID")
	subreddit := os.Getenv("SUBREDDIT")

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	threader := discord.NewDiscordGoThreader(session)
	interaction := discord.NewInteraction(
		threader,
		channelId,
	)

	p := posts.New(
		interaction,
		subreddit,
		reddit.Top,
	)

	if err := p.Run(); err != nil {
		panic(err)
	}
}
