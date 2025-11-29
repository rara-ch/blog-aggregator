package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	request.Header.Add("User-Agent", "gator")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return &RSSFeed{}, err
	}
	body := response.Body
	defer response.Body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return &RSSFeed{}, err
	}

	var rssFeed RSSFeed
	rssFeedPointer := &rssFeed
	err = xml.Unmarshal(data, rssFeedPointer)
	if err != nil {
		return &RSSFeed{}, err
	}

	cleanFeed(rssFeedPointer)

	return rssFeedPointer, err
}

func cleanFeed(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
