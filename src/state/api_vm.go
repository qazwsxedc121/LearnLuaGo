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

func (l *luaState) RegisterCount() int {
	return int(l.stack.closure.proto.MaxStackSize)
}

func (l *luaState) LoadVararg(n int) {
	if n < 0 {
		n = len(l.stack.varargs)
	}
	l.stack.check(n)
	l.stack.pushN(l.stack.varargs, n)
}

func (l *luaState) LoadProto(idx int) {
	proto := l.stack.closure.proto.Protos[idx]
	closure := newLuaClosure(proto)
	l.stack.push(closure)
}


