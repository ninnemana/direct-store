package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/ninnemana/direct-store/direct-store/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
