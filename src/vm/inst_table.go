package vm

import . "api"

func newTable(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.CreateTable(Fb2int(b), Fb2int(c))
	vm.Replace(a)
}

func getTable(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.PushValue(c)
	vm.GetTable(b+1)
	vm.Replace(a)
}

func setTable(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.PushValue(b+1)
	vm.PushValue(c+1)
	vm.SetTable(a)
}

func getField(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.GetConst(c)
	vm.GetTable(b+1)
	vm.Replace(a)
}

func getI(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.PushInteger(int64(c))
	vm.GetTable(b+1)
	vm.Replace(a)
}

func setField(i Instruction, vm LuaVM) {
	a, b, c, k := i.ABC()
	a += 1
	if k {
		vm.GetConst(b)
		vm.GetConst(c)
	}
	vm.SetTable(a)

}

func setI(i Instruction, vm LuaVM) {
	a, b, c, k := i.ABC()
	a += 1
	vm.PushInteger(int64(b))
	if k {
		vm.GetConst(c)
	}else{
		vm.PushValue(c+1)
	}
	vm.SetTable(a)
}

func setList(i Instruction, vm LuaVM) {
	a, b, c, k := i.ABC()
	a += 1
	n := b
	last := c

	if n == 0 {
		n = int(vm.ToInteger(-1)) - a - 1
		vm.Pop(1)
	}

	last += n
	if k {
		last += Instruction(vm.Fetch()).Ax()
	}
	idx := int64(c * LFIELDS_PER_FLUSH)
	for j := 1; j <= b; j++ {
		idx++
		vm.PushValue(a + j)
		vm.SetI(a, idx)
	}
}

const LFIELDS_PER_FLUSH = 50



