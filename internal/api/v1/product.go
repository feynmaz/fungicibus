package v1

import (
	"net/http"

	"github.com/feynmaz/fungicibus/internal/types"
)

// List all available products
// (GET /products)
func (api *API) GetProducts(w http.ResponseWriter, r *http.Request, params GetProductsParams) {
	productListFilters := api.parseProductParams(params)

	_ = productListFilters
}

func (api *API) parseProductParams(params GetProductsParams) types.ProductListFilters {
	filters := types.ProductListFilters{
		Limit:       int(GetProductsParamsLimitDefault),
		Page:        1,
		InStockOnly: false,
		Category:    "",
	}

	if params.Limit != nil {
		if *params.Limit >= GetProductsParamsLimitMin && *params.Limit <= GetProductsParamsLimitMax {
			filters.Limit = int(*params.Limit)
		}
	}

	if params.Page != nil && *params.Page > 0 {
		filters.Page = int(*params.Page)
	}

	if params.InStockOnly != nil {
		filters.InStockOnly = *params.InStockOnly
	}

	if params.Category != nil {
		switch *params.Category {
		case GetProductsParamsCategoryCulinary:
			filters.Category = string(GetProductsParamsCategoryCulinary)
		case GetProductsParamsCategoryExotic:
			filters.Category = string(GetProductsParamsCategoryExotic)
		case GetProductsParamsCategoryMedicinal:
			filters.Category = string(GetProductsParamsCategoryMedicinal)
		}
	}

	return filters
}
