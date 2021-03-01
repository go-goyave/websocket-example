package chat

import "goyave.dev/goyave/v3/validation"

var (
	// JoinRequest the validation rules joining the chat room.
	JoinRequest = validation.RuleSet{
		"name": {"required", "string", "between:3,50"},
	}
)
