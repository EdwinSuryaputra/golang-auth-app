package supplier

import (
	"context"
	"fmt"

	suppliermodel "golang-auth-app/app/common/models/supplier"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"

	config "golang-auth-app/config"

	"github.com/rotisserie/eris"
)

func (i *impl) GetSupplier(
	ctx context.Context,
	supplierId string,
	supplierName string,
) ([]*suppliermodel.Supplier, error) {
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
		Id           string `json:"id"`
		Name         string `json:"name"`
		SupplierCode string `json:"supplier_code"`
	}

	payloadBody := &apiPayload{
		TableName:      "suppliers",
		Page:           1,
		Limit:          100,
		PropertyFilter: map[string]string{},
	}
	if supplierId != "" {
		payloadBody.PropertyFilter["id"] = supplierId
	}
	if supplierName != "" {
		payloadBody.PropertyFilter["name"] = fmt.Sprintf("%%%s%%", supplierName)
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
		return nil, eris.Wrap(err, "error occurred during getSuppliers")
	} else if resp.IsError() {
		return nil, eris.New("getSuppliers got error response")
	}

	result, err := sliceutil.MapWithError(successRespBody.DataList, func(dt *apiSuccessResponseDataList) (*suppliermodel.Supplier, error) {
		decodedId, err := publicfacingutil.Decode(dt.Id)
		if err != nil {
			return nil, err
		}

		return &suppliermodel.Supplier{
			Id:           decodedId,
			Name:         dt.Name,
			SupplierCode: dt.SupplierCode,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
