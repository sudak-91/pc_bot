package command

import (
	"context"

	"github.com/sudak-91/pc_bot/pkg/repository"
)

//TODO:
type MarkAsAnswer struct {
	Question repository.Questions
	ctx      context.Context
}
