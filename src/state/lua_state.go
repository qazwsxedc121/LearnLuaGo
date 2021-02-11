package state

import "binchunk"

type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype
	pc int
}

