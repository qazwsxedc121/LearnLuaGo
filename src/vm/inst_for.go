package vm

import . "api"

func forPrep(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1
	if _forPrep(a, vm){
		vm.AddPC(bx+1)
	}
}

func _forPrep(a int, vm LuaVM) bool {
	init := vm.ToInteger(a)
	limit := vm.ToInteger(a+1)
	step := vm.ToInteger(a+2)
	vm.PushInteger(init)
	vm.Replace(a+3)

	if init > limit{
		return true
	}

	count := limit - init
	if step != 1 {
		count = count / step
	}
	vm.PushInteger(count)
	vm.Replace(a+1)
	return false
}

func forLoop(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	count := vm.ToInteger(a+2)
	if count > 0 {
		step := vm.ToInteger(a+3)
		idx := vm.ToInteger(a+1)

		vm.PushInteger(count-1)
		vm.Replace(a+2)

		idx = idx + step
		vm.PushInteger(idx)
		vm.Replace(a+1)

		vm.PushInteger(idx)
		vm.Replace(a+4)

		vm.AddPC(-bx)

	}
}
