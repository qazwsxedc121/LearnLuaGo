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
	if !vm.IsNil(a+4) {
		vm.PushInteger(int64(a+4))
		vm.Replace(a+2)
		vm.AddPC(-bx)
	}
}
