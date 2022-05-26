package types

import "encoding/json"

type Pipeline struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Stages      []Stage `json:"stages"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type Stage struct {
	Description string `json:"description"`
}

type LambdaEvent struct {
	Arguments json.RawMessage `json:"arguments"`
}
