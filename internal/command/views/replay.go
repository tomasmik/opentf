package views

import "github.com/placeholderplaceholderplaceholder/opentf/internal/replay"

func StreamReplay(view *View) error {
	return replay.ReplayApply(view.streams)
}
