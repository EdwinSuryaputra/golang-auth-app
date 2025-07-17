package activitylog

import (
	"context"
	"fmt"
	config "golang-auth-app/config"

	"github.com/rotisserie/eris"
)

func (i *impl) Insert(
	ctx context.Context,
	activityLogId string,
	message string,
	status string,
) error {
	type apiPayload struct {
		ActivityLogId string `json:"activityLogId"`
		Message       string `json:"message"`
	}

	payloadBody := &apiPayload{
		ActivityLogId: activityLogId,
		Message:       message,
	}

	cfg := config.InternalService.Medio

	resp, err := i.restyClient.R().
		SetContext(ctx).
		SetHeaders(map[string]string{
			"Authorization": ctx.Value("authToken").(string),
		}).
		SetBody(payloadBody).
		Post(fmt.Sprintf("%s%s", cfg.Host, cfg.Routes.GetActivityLog))
	if err != nil {
		return eris.Wrap(err, "error occurred during insertActivityLog")
	} else if resp.IsError() {
		return eris.New("insertActivityLog got error response")
	}

	return nil
}
