package main

import (
	"log"
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
		log.Fatalf("Cannot create bot session: %v", err)
	}

	threader := discord.NewDiscordGoThreader(session)
	interaction := discord.NewInteraction(
		threader,
		channelId,
	)

	p, err := posts.New(
		interaction,
		subreddit,
		reddit.Best,
	)
	if err != nil {
		log.Fatalf("Cannot create posts: %v", err)
	}

	if err := p.Run(); err != nil {
		log.Fatalf("Failure to run: %v", err)
	}
}
