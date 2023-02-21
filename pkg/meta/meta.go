package meta

import (
	"os"
	"strconv"
)

type Meta struct {
	TotalCount int `json:"total_count"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
	Page       int `json:"page"`
}

func New(page, perPage, total int) (*Meta, error) {

	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(os.Getenv("PAGINATOR_LIMIT_DEFAULT"))

		if err != nil {
			return nil, err
		}

	}

	pageCount := 0

	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage

		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		TotalCount: total,
		PerPage:    perPage,
		PageCount:  pageCount,
		Page:       page,
	}, nil
}

func (p *Meta) Offset() int {
	return (p.Page - 1) * p.PerPage
}

func (p *Meta) Limit() int {
	return p.PerPage
}
