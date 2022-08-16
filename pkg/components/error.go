package components

import (
	"github.com/lzakharov/remote-config-manager/pkg/components/metro"
)

func handleErr(err error) {
	metro.Notify(err.Error(), "alert", true)
}
