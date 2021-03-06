package vm

import . "api"

func closure(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1
	vm.LoadProto(bx)
	vm.Replace(a)
}

func call(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	_popResults(a, c, vm)
}

func _pushFuncAndArgs(a, b int, vm LuaVM) (nArgs int) {
	if b >= 1 {
		vm.CheckStack(b)
		for i := a; i < a+b; i++ {
			vm.PushValue(i)
		}
		return b - 1
	} else {
		_fixStack(a, vm)
		return vm.GetTop() - vm.RegisterCount() - 1
	}
}

func _popResults(a, c int, vm LuaVM) {
	if c == 1 {
	} else if c > 1 {
		for i := a + c - 2; i >= a; i-- {
			vm.Replace(i)
		}
	} else {
		vm.CheckStack(1)
		vm.PushInteger(int64(a))
	}
}

func _fixStack(a int, vm LuaVM) {
	x := int(vm.ToInteger(-1))
	vm.Pop(1)

	vm.CheckStack(x - a)
	for i := a; i < x; i++ {
		vm.PushValue(i)
	}
	vm.Rotate(vm.RegisterCount()+1, x - a)

}

func _return(i Instruction, vm LuaVM) {
	a, b, _, _ := i.ABC()
	a += 1
	if b == 1 {

	}else if b > 1 {
		vm.CheckStack(b - 1)
		for i := a; i <= a+b-2; i++ {
			vm.PushValue(i)
		}
	}else {
		_fixStack(a, vm)
	}
}

