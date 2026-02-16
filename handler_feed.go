package main

import (
	"aggreGATOR/internal/database"
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	feed := RSSFeed{}
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &feed, err
	}
	// Set headers
	req.Header.Add("User-Agent", "gator")

	resp, err := client.Do(req)
	if err != nil {
		return &feed, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return &feed, err
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i] = item
	}
	return &feed, nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("aggregate command needs a single argument, the url")
	}
	_, err := fetchFeed(context.Background(), "")
	if err != nil {
		return err
	}
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("addfeed command needs a two arguments, first the name, second the url")
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		Url: sql.NullString{
			String: url,
			Valid:  true,
		},
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Name:   name,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("listfeeds command needs no argument")
	}

	ctx := context.Background()
	feeds, err := s.db.ListAllFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		u, err := s.db.GetUserById(ctx, feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Name: %s, URL: %v, Owner: %s\n", feed.Name, feed.Url, u.Name)
	}

	return nil
}
