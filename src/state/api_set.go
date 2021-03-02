package state

func (l *luaState) SetTable(idx int) {
	t := l.stack.get(idx)
	v := l.stack.pop()
	k := l.stack.pop()
	l.setTable(t, k, v)
}

func (l *luaState) setTable(t, k, v luaValue) {
	if tbl, ok := t.(*luaTable); ok {
		tbl.put(k, v)
		return
	}
	panic("not a table!")
}

func (l *luaState) SetField(idx int, k string) {
	t := l.stack.get(idx)
	v := l.stack.pop()
	l.setTable(t, k, v)
}

func (l *luaState) SetI(idx int, i int64) {
	t := l.stack.get(idx)
	v := l.stack.pop()
	l.setTable(t, i, v)
}

