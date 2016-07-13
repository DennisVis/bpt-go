package models

type Model interface {
	GetId() int
}

type Question struct {
	Id     int                `json:"id"`
	Name   string             `json:"name"`
	Labels map[string]string  `json:"labels"`
}

func (q Question) GetId() int {
  return q.Id
}
