package web

type BookResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Quantity int    `json:"quantity"`
}
