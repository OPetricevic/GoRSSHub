package main

import "github.com/google/uuid"

type errResponse struct {
	Error string `json:"error"`
}

type parameters struct {
	Name   string    `json:"name"`
	URL    string    `json:"url"`
	FeedID uuid.UUID `json:"feed_id"`
}

type ApiResponse struct {
	Message string `json:"message"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
