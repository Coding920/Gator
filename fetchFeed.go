package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return &RSSFeed{}, fmt.Errorf("Server response: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	var rssContent RSSFeed
	err = xml.Unmarshal(data, &rssContent)
	if err != nil {
		return &RSSFeed{}, err
	}
	rssContent.Channel.Title = html.UnescapeString(rssContent.Channel.Title)
	rssContent.Channel.Description = html.UnescapeString(rssContent.Channel.Description)

	for _, item := range rssContent.Channel.Item {
		item.Title = html.EscapeString(item.Title)
		item.Description = html.EscapeString(item.Description)
	}

	return &rssContent, nil
}
