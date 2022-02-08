package handler

import (
	"context"

	aj "github.com/choria-io/asyncjobs"
)

func AsyncJobHandler(ctx context.Context, log aj.Logger, task *aj.Task) (interface{}, error) {
	log.Infof("Handling task %s", task.ID)

	return "success", nil
}
