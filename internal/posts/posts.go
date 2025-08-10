package posts

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vectorhacker/reddit-discord-bot/pkg/discord"
	"github.com/vectorhacker/reddit-discord-bot/pkg/reddit"

	_ "modernc.org/sqlite"
)

type Posts struct {
	d         *discord.DiscordInteraction
	subreddit string
	sorting   reddit.ScrapeSorting
	lastPage  string
	db        *sql.DB
}

const table = `
CREATE TABLE IF NOT EXISTS checkpoints (
	subreddit text,
	sorting text,
	lastPage text,
	PRIMARY KEY (subreddit, sorting)
);
`

func New(d *discord.DiscordInteraction, subreddit string, sorting reddit.ScrapeSorting) (*Posts, error) {
	db, err := sql.Open("sqlite", ".checkpoints.sqlite")
	if err != nil {
		return nil, fmt.Errorf("unable to open sqlite: %v", err)
	}
	_, err = db.Exec(table)
	if err != nil {
		return nil, fmt.Errorf("could not get checkpoint %v", err)
	}

	rows, err := db.Query("SELECT lastPage FROM checkpoints WHERE subreddit = ? AND sorting = ?", subreddit, sorting)
	if err != nil {
		return nil, fmt.Errorf("could not get checkpoint %v", err)
	}
	defer rows.Close()

	var lastPage string
	if rows.Next() {
		err = rows.Scan(&lastPage)
		if err != nil {
			return nil, fmt.Errorf("could not get checkpoint %v", err)
		}
	}

	return &Posts{
		d,
		subreddit,
		sorting,
		lastPage,
		db,
	}, nil
}

func (p *Posts) Run() error {
	scrape, err := reddit.ScrapeSubreddit(p.subreddit, p.sorting, p.lastPage)
	if err != nil {
		return err
	}

	err = p.d.ThreadRedditPosts(scrape.Posts)
	if err != nil {
		return fmt.Errorf("problem posting: %v", err)
	}

	result, err := p.db.Exec("INSERT INTO checkpoints(subreddit, sorting, lastPage) VALUES($1, $2, $3) ON CONFLICT(subreddit, sorting) DO UPDATE SET lastPage=$3 WHERE subreddit=$1 AND sorting=$2;", p.subreddit, p.sorting, scrape.After)
	if rows, err := result.RowsAffected(); err != nil {

		return err
	} else {

		log.Printf("%v", rows)
	}
	if err != nil {
		return fmt.Errorf("problem checkpointing: %v", err)
	}

	return err
}
