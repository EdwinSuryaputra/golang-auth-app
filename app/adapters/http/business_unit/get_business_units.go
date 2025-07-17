package businessunit

import (
	"context"
	"fmt"
	"strings"

	businessunitmodel "golang-auth-app/app/common/models/business_unit"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"

	config "golang-auth-app/config"

	"github.com/rotisserie/eris"
)

func (i *impl) GetBusinessUnits(
	ctx context.Context,
	businessUnitLevel string,
	businessUnitId string,
) ([]*businessunitmodel.BusinessUnit, error) {
	type apiPayload struct {
		TableName      string            `json:"tableName"`
		Page           int32             `json:"page"`
		Limit          int32             `json:"limit"`
		PropertyFilter map[string]string `json:"propertyFilter"`
	}

	type apiSuccessResponse[T any] struct {
		DataList []*T `json:"dataList"`
	}

	type apiSuccessResponseDataList struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	tableNameMappings := map[string]string{
		"AREA":      "areas",
		"REGION":    "regions",
		"WITEL":     "witels",
		"WAREHOUSE": "warehouses",
	}

	payloadBody := &apiPayload{
		TableName:      tableNameMappings[strings.ToUpper(businessUnitLevel)],
		Page:           1,
		Limit:          100,
		PropertyFilter: map[string]string{},
	}
	if businessUnitId != "" {
		payloadBody.PropertyFilter["id"] = businessUnitId
	}

	var successRespBody *apiSuccessResponse[apiSuccessResponseDataList]
	cfg := config.InternalService.Medio

	resp, err := i.restyClient.R().
		SetContext(ctx).
		SetHeaders(map[string]string{
			"Authorization": ctx.Value("authToken").(string),
		}).
		SetResult(&successRespBody).
		SetBody(payloadBody).
		Post(fmt.Sprintf("%s%s", cfg.Host, cfg.Routes.GetList))
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetBusinessUnits")
	} else if resp.IsError() {
		return nil, eris.Wrap(err, "GetBusinessUnits got error response")
	}

	result, err := sliceutil.MapWithError(successRespBody.DataList, func(dt *apiSuccessResponseDataList) (*businessunitmodel.BusinessUnit, error) {
		decodedId, err := publicfacingutil.Decode(dt.Id)
		if err != nil {
			return nil, err
		}

		return &businessunitmodel.BusinessUnit{
			Id:   decodedId,
			Name: dt.Name,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
