package domain

// ResourceList is a generic container for listing resources.
// Note: ApiVersion and Kind are API-layer concerns, not domain concerns.
// They are set by the converter layer when converting to API types.
type ResourceList[T any] struct {
	Items      []T
	Pagination Pagination
}

// Pagination holds pagination state for list results.
type Pagination struct {
	// Continue is the opaque pagination token.
	// nil means no more items available.
	Continue *string

	// RemainingItemCount is the estimated count of remaining items.
	RemainingItemCount *int64
}

// HasMore returns true if there are more pages available.
func (p Pagination) HasMore() bool {
	return p.Continue != nil
}

// NewResourceList creates a new ResourceList.
func NewResourceList[T any](items []T, pagination Pagination) ResourceList[T] {
	return ResourceList[T]{
		Items:      items,
		Pagination: pagination,
	}
}

// SetContinue sets pagination continuation data.
func (r *ResourceList[T]) SetContinue(cont *string, numRemaining *int64) {
	if cont != nil {
		r.Pagination.Continue = cont
		r.Pagination.RemainingItemCount = numRemaining
	}
}

// EmptyResourceList creates an empty ResourceList.
func EmptyResourceList[T any]() ResourceList[T] {
	return ResourceList[T]{
		Items:      []T{},
		Pagination: Pagination{},
	}
}

// ResourceListParams contains common list query parameters for resources.
// This is a domain type that is independent of API versions.
type ResourceListParams struct {
	Continue      *string
	LabelSelector *string
	FieldSelector *string
	Limit         *int32
}

// SortOrder represents the sort order for list queries.
type SortOrder string

const (
	SortAsc  SortOrder = "asc"
	SortDesc SortOrder = "desc"
)
