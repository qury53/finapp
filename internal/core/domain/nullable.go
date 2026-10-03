package domain

type Nullable[T any] struct {
	Value *T   `json:"value"`
	Set   bool `json:"set"`
}
