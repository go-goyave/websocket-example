package chat

import "github.com/System-Glitch/goyave/v3/validation"

var (
	// JoinRequest the validation rules joining the chat room.
	JoinRequest = validation.RuleSet{
		"name": {"required", "string", "between:3,50"},
	}
)
