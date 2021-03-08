package vm

import . "api"

func move(i Instruction, vm LuaVM) {
	a, b, _, _ := i.ABC()
	a += 1
	b += 1
	vm.Copy(b, a)
}

func jmp(i Instruction, vm LuaVM) {
	a := i.SJ()
	vm.AddPC(a)
}

func varargprep(i Instruction, vm LuaVM) {

}

func extraarg(i Instruction, vm LuaVM) {

}

func setTabup(i Instruction, vm LuaVM) {

}

func getTabup(i Instruction, vm LuaVM) {

}

