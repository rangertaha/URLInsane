package none

import (
	"github.com/rangertaha/urlinsane"
	"github.com/rangertaha/urlinsane/plugins/algorithms"
)

type WrongTLD struct {
	// Code() string
	// Name() string
	// Description() string
	// Fields() []string
	// Exec() func(Result) []Result
}

func (n *WrongTLD) Code() string {
	return "wtld"
}

func (n *WrongTLD) Name() string {
	return "Wrong TLD"
}

func (n *WrongTLD) Description() string {
	return "Wrong Top Level Domain"
}

func (n *WrongTLD) Fields() []string {
	return []string{}
}

func (n *WrongTLD) Headers() []string {
	return []string{}
}

func (n *WrongTLD) Exec(urlinsane.Typo) (results []urlinsane.Typo) {
	return
}

// Register the plugin
func init() {
	algorithms.Add("wtld", func() urlinsane.Algorithm {
		return &WrongTLD{}
	})
}
