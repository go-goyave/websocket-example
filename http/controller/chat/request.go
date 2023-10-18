package chat

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

// JoinRequest the query validation rules joining the chat room.
func JoinRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: "name", Rules: v.List{v.Required(), v.String(), v.Between(2, 50)}},
	}
}
