package vm

import . "api"

func newTable(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.CreateTable(Fb2int(b), Fb2int(c))
	vm.Replace(a)
}

func getTable(i Instruction, vm LuaVM) {
	a, b, c, k := i.ABC()
	a += 1
	b += 1
	vm.GetRK(c, k)
	vm.GetTable(b)
	vm.Replace(a)
}

func setTable(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.PushValue(b+1)
	vm.PushValue(c+1)
	vm.SetTable(a)
}

func setList(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	if c > 0 {
		c = c - 1
	} else {
		c = Instruction(vm.Fetch()).Ax()
	}
	idx := int64(c * LFIELDS_PER_FLUSH)
	for j := 1; j <= b; j++ {
		idx++
		vm.PushValue(a + j)
		vm.SetI(a, idx)
	}
}

const LFIELDS_PER_FLUSH = 50



