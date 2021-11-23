package chat

import "goyave.dev/goyave/v4/validation"

var (
	// JoinRequest the validation rules joining the chat room.
	JoinRequest = validation.RuleSet{
		"name": validation.List{"required", "string", "between:3,50"},
	}
)
