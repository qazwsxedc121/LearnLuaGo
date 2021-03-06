package state

func (l *luaState) PC() int {
	return l.stack.pc
}

func (l *luaState) AddPC(n int) {
	l.stack.pc += n
}

func (l *luaState) Fetch() uint32 {
	i := l.stack.closure.proto.Code[l.stack.pc]
	l.stack.pc++
	return i
}

func (l *luaState) GetConst(idx int) {
	c := l.stack.closure.proto.Constants[idx]
	l.stack.push(c)
}

func (l *luaState) GetRK(r int, k bool) {
	if k {
		l.GetConst(r)
	} else {
		l.PushValue(r + 1)
	}
}
