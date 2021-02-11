package vm

import . "api"

func loadNil(i Instruction, vm LuaVM) {
	a, b, _, _ := i.ABC()
	a += 1
	vm.PushNil()
	for i := a; i <= a + b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}

func loadTrue(i Instruction, vm LuaVM) {
	a, _, _, _ := i.ABC()
	a += 1
	vm.PushBoolean(true)
	vm.Replace(a)
}

func loadFalse(i Instruction, vm LuaVM) {
	a, _, _, _ := i.ABC()
	a += 1
	vm.PushBoolean(false)
	vm.Replace(a)
}

func loadFalseSkip(i Instruction, vm LuaVM) {
	loadFalse(i, vm)
	vm.AddPC(1)
}

func loadK(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1
	vm.GetConst(bx)
	vm.Replace(a)
}

func loadKx(i Instruction, vm LuaVM) {
	a, _ := i.ABx()
	a += 1
	ax := Instruction(vm.Fetch()).Ax()
	vm.GetConst(ax)
	vm.Replace(a)
}

func loadI(i Instruction, vm LuaVM) {
	a, b := i.AsBx()
	a += 1
	vm.PushInteger(int64(b))
	vm.Replace(a)
}

func loadF(i Instruction, vm LuaVM) {
	a, b := i.AsBx()
	a += 1
	vm.PushNumber(float64(b))
	vm.Replace(a)
}

