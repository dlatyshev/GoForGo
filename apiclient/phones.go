package apiclient

type Phone struct {
	ID   string        `json:"id"`
	Name string        `json:"name"`
	Color string       `json:"data.color"`
	Capacity string    `json:"data.capacity"`
}

func (p *Phone) ToString() string {
	return p.Name + " " + p.Color + " " + p.Capacity
}