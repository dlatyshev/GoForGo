package apiclient


type Phone struct {
	ID   string        `json:"id"`
	Name string        `json:"name"`
	Color string       `json:"data.color"`
	Capacity string    `json:"data.capacity"`
}