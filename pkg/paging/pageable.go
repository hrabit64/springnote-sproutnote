package paging

type Pageable struct {
	Page     int
	PageSize int
}

// Offset returns the offset of the current page
func (p *Pageable) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit returns the limit of the current page
func (p *Pageable) Limit() int {
	return p.PageSize
}
