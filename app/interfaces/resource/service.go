package resource

import (
	"context"
	resourceDto "golang-auth-app/app/interfaces/resource/dto"
)

type Service interface {
	// GetAllResources(ctx context.Context) (*resourceDto.ResourceObj, error)

	GetPublicResource(
		ctx context.Context,
		payload *resourceDto.GetPublicResourcePayload,
	) (*resourceDto.ResourcePubObj, error)
}
