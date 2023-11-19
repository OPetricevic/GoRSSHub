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
