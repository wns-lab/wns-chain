package types

import (
	ens "github.com/wealdtech/go-ens"
)

func (name Name) IsValid() bool {
	err := ens.Normalize(name.Name)
	if err != nil {
		return false
	}
	return true
}
