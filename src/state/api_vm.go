package state

func (l *luaState) PC() int {
	return l.pc
}

func (l *luaState) AddPC(n int) {
	l.pc += n
}

func (l *luaState) Fetch() uint32 {
	i := l.proto.Code[l.pc]
	l.pc++
	return i
}

func (l *luaState) GetConst(idx int) {
	c := l.proto.Constants[idx]
	l.stack.push(c)
}

func (l *luaState) GetRK(r int, k bool) {
	if k {
		l.GetConst(r)
	} else {
		l.PushValue(r + 1)
	}
}
