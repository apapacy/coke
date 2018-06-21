package grifts

import (
	"github.com/apapacy/coke/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
