package main

type errResponse struct {
	Error string `json:"error"`
}

type parameters struct {
	Name string `json:"name"`
}
