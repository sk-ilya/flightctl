package model

import (
	"encoding/json"
	"fmt"

	"github.com/flightctl/flightctl/internal/domain"
	"github.com/flightctl/flightctl/internal/util"
	"github.com/samber/lo"
)

type Event struct {
	Resource
	Reason             string                          `gorm:"type:string;index" selector:"reason"`
	SourceComponent    string                          `gorm:"type:string"`
	Actor              string                          `gorm:"type:string" selector:"actor"`
	Type               string                          `gorm:"type:string;index" selector:"type"`
	Message            string                          `gorm:"type:text"`
	Details            *JSONField[domain.EventDetails] `gorm:"type:jsonb"`
	InvolvedObjectName string                          `gorm:"type:string;index:idx_involved_object" selector:"involvedObject.name"`
	InvolvedObjectKind string                          `gorm:"type:string;index:idx_involved_object" selector:"involvedObject.kind"`
}

func (e Event) String() string {
	val, _ := json.Marshal(e)
	return string(val)
}

func NewEventFromDomain(resource *domain.Event) (*Event, error) {
	if resource == nil {
		return &Event{}, nil
	}
	details := domain.EventDetails{}
	if resource.Details != nil {
		details = *resource.Details
	}
	return &Event{
		Resource: Resource{
			Name:        *resource.Metadata.Name,
			Annotations: lo.FromPtrOr(resource.Metadata.Annotations, make(map[string]string)),
		},
		Reason:             string(resource.Reason),
		SourceComponent:    resource.Source.Component,
		Actor:              resource.Actor,
		Type:               string(resource.Type),
		Message:            resource.Message,
		Details:            MakeJSONField(details),
		InvolvedObjectName: resource.InvolvedObject.Name,
		InvolvedObjectKind: resource.InvolvedObject.Kind,
	}, nil
}

func EventAPIVersion() string {
	return fmt.Sprintf("%s/%s", domain.APIGroup, domain.EventAPIVersion)
}

func (e *Event) ToDomain(opts ...APIResourceOption) (*domain.Event, error) {
	if e == nil {
		return &domain.Event{}, nil
	}

	var details *domain.EventDetails
	if e.Details != nil {
		details = &e.Details.Data
	}

	return &domain.Event{
		ApiVersion: EventAPIVersion(),
		Kind:       domain.EventKind,
		Metadata: domain.ObjectMeta{
			Name:              lo.ToPtr(e.Name),
			Annotations:       lo.ToPtr(util.EnsureMap(e.Resource.Annotations)),
			CreationTimestamp: lo.ToPtr(e.CreatedAt.UTC()),
		},
		InvolvedObject: domain.ObjectReference{
			Kind: e.InvolvedObjectKind,
			Name: e.InvolvedObjectName,
		},
		Reason: domain.EventReason(e.Reason),
		Source: domain.EventSource{
			Component: e.SourceComponent,
		},
		Actor:   e.Actor,
		Type:    domain.EventType(e.Type),
		Message: e.Message,
		Details: details,
	}, nil
}

func EventsToDomain(events []Event, cont *string, numRemaining *int64) (domain.ResourceList[domain.Event], error) {
	eventList := make([]domain.Event, len(events))
	for i, event := range events {
		var opts []APIResourceOption
		apiResource, _ := event.ToDomain(opts...)
		eventList[i] = *apiResource
	}
	metadata := domain.Pagination{}
	if cont != nil {
		metadata.Continue = cont
		metadata.RemainingItemCount = numRemaining
	}
	return domain.NewResourceList(eventList, metadata), nil
}

func (e *Event) GetKind() string {
	return domain.EventKind
}

func (e *Event) HasNilSpec() bool {
	return true
}

func (e *Event) HasSameSpecAs(otherResource any) bool {
	return true
}

func (e *Event) GetStatusAsJson() ([]byte, error) {
	return nil, nil
}
