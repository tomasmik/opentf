package command

import (
	"fmt"

	"github.com/placeholderplaceholderplaceholder/opentf/internal/command/arguments"
	"github.com/placeholderplaceholderplaceholder/opentf/internal/command/views"
)

type ReplayCommand struct {
	Meta
}

func (r *ReplayCommand) Run(rawArgs []string) int {
	common, rawArgs := arguments.ParseView(rawArgs)
	r.View.Configure(common)

	r.Meta.color = !common.NoColor
	r.Meta.Color = r.Meta.color

	if err := views.StreamReplay(r.View); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func (r *ReplayCommand) Help() string {
	return ""
}

func (r *ReplayCommand) Synopsis() string {
	return "Create privous logs"
}
