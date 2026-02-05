package service

import (
	"fmt"

	"github.com/flightctl/flightctl/internal/domain"
	"github.com/flightctl/flightctl/internal/store"
	"github.com/flightctl/flightctl/internal/store/selector"
	"github.com/samber/lo"
)

func prepareListParams(params domain.ResourceListParams) (*store.ListParams, domain.Status) {
	cnt, err := store.ParseContinueString(params.Continue)
	if err != nil {
		return nil, domain.StatusBadRequest(fmt.Sprintf("failed to parse continue parameter: %v", err))
	}

	var fieldSelector *selector.FieldSelector
	if params.FieldSelector != nil {
		if fieldSelector, err = selector.NewFieldSelector(*params.FieldSelector); err != nil {
			return nil, domain.StatusBadRequest(fmt.Sprintf("failed to parse field selector: %v", err))
		}
	}

	var labelSelector *selector.LabelSelector
	if params.LabelSelector != nil {
		if labelSelector, err = selector.NewLabelSelector(*params.LabelSelector); err != nil {
			return nil, domain.StatusBadRequest(fmt.Sprintf("failed to parse label selector: %v", err))
		}
	}

	listParams := &store.ListParams{
		Limit:         int(lo.FromPtr(params.Limit)),
		Continue:      cnt,
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	if listParams.Limit == 0 {
		listParams.Limit = MaxRecordsPerListRequest
	} else if listParams.Limit > MaxRecordsPerListRequest {
		return nil, domain.StatusBadRequest(fmt.Sprintf("limit cannot exceed %d", MaxRecordsPerListRequest))
	} else if listParams.Limit < 0 {
		return nil, domain.StatusBadRequest("limit cannot be negative")
	}

	return listParams, domain.StatusOK()
}
