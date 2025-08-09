package dtos

import "math"

type PaginationMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	IsLastPage bool  `json:"is_last_page"`
}

func CreatePaginationMeta(page, pageSize int, total int64) PaginationMeta {
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	isLastPage := page >= totalPages

	return PaginationMeta{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
		IsLastPage: isLastPage,
	}
}
