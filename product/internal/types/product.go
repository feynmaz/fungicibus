package types

type ProductListFilters struct {
	Limit       int
	Page        int
	InStockOnly bool
	Category    string
}
