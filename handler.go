package handler

import (
	"context"
	"log"

	aj "github.com/choria-io/asyncjobs"
)

func AsyncJobHandler(ctx context.Context, task *aj.Task) (interface{}, error) {
	log.Printf("Handling task %s", task.ID)

	return "success", nil
}
