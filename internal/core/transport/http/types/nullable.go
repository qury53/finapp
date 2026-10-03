package core_http_types

import (
	"encoding/json"

	"github.com/qury53/finapp/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	n.Set = true

	if string(data) == "null" {
		n.Value = nil
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	n.Value = &v
	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{
		Set:   n.Set,
		Value: n.Value,
	}
}
