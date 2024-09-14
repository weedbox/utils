package pagination

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Meta struct {
	Page    int   `json:"page"`
	PerPage int   `json:"perPage"`
	Total   int64 `json:"total"`
}

type Pagination struct {
	Page    int
	PerPage int
	Offset  int
}

func NewPagination(page int, perPage int) (*Pagination, error) {

	opts := &Pagination{
		PerPage: perPage,
		Page:    page,
		Offset:  0,
	}

	opts.Offset = opts.PerPage * (opts.Page - 1)

	err := validation.ValidateStruct(opts,
		validation.Field(&opts.PerPage, validation.Min(-1), validation.Max(100)),
		validation.Field(&opts.Page, validation.Min(1)),
	)

	if err != nil {
		return nil, err
	}

	return opts, nil
}
