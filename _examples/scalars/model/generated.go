// Code generated by github.com/niko0xdev/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/niko0xdev/gqlgen/_examples/scalars/external"
)

type Address struct {
	ID       external.ObjectID `json:"id"`
	Location *Point            `json:"location,omitempty"`
}

type Query struct {
}
