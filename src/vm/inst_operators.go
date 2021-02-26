package vm

import . "api"

func _binaryArith(i Instruction, vm LuaVM, op ArithOp){
	a, b, c, _ := i.ABC()
	a += 1
	vm.PushValue(b+1)
	vm.PushValue(c+1)
	vm.Arith(op)
	vm.Replace(a)
}

func _binaryArithK(i Instruction, vm LuaVM, op ArithOp) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.PushValue(b+1)
	vm.GetConst(c)
	vm.Arith(op)
	vm.Replace(a)
}

func _unaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, _, _ := i.ABC()
	a += 1
	vm.PushValue(b+1)
	vm.Arith(op)
	vm.Replace(a)
}

func addi(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.PushValue(b+1)
	vm.PushInteger(int64(c))
	vm.Arith(LUA_OPADD)
	vm.Replace(a)
}

func shri(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.PushValue(b+1)
	vm.PushInteger(int64(c))
	vm.Arith(LUA_OPSHR)
	vm.Replace(a)
}

func shli(i Instruction, vm LuaVM) {
	a, b, c, _ := i.ABC()
	a += 1
	vm.PushInteger(int64(c))
	vm.PushValue(b+1)
	vm.Arith(LUA_OPSHL)
	vm.Replace(a)
}

func add(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPADD)}
func sub(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPSUB)}
func mul(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPMUL)}
func mod(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPMOD)}
func pow(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPPOW)}
func div(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPDIV)}
func idiv(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPIDIV)}
func band(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPBAND)}
func bor(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPBOR)}
func bxor(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPBXOR)}

func shl(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPSHL)}
func shr(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPSHR)}

func unm(i Instruction, vm LuaVM) { _unaryArith(i, vm, LUA_OPUNM)}
func bnot(i Instruction, vm LuaVM) { _unaryArith(i, vm, LUA_OPBNOT)}

func addk(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPADD)}
func subk(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPSUB)}
func mulk(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPMUL)}
func modk(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPMOD)}
func powk(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPPOW)}
func divk(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPDIV)}
func idivk(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPIDIV)}
func bandk(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPBAND)}
func bork(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPBOR)}
func bxork(i Instruction, vm LuaVM) { _binaryArithK(i, vm, LUA_OPBXOR)}


func len(i Instruction, vm LuaVM) {
	a, b, _, _ := i.ABC()
	a += 1
	b += 1
	vm.Len(b)
	vm.Replace(a)
}

func concat(i Instruction, vm LuaVM) {
	a, b, _, _ := i.ABC()
	a += 1
	vm.CheckStack(b)
	for i := a; i < a + b; i++ {
		vm.PushValue(i)
	}
	vm.Concat(b)
	vm.Replace(a)
}

func _compare(i Instruction, vm LuaVM, op CompareOp) {
	a, b, _, k := i.ABC()
	a += 1
	b += 1
	vm.PushValue(a)
	vm.PushValue(b)
	if vm.Compare(-2, -1, op) != k {
		vm.AddPC(1)
	}
	vm.Pop(2)
}

func eq(i Instruction, vm LuaVM) { _compare(i, vm, LUA_OPEQ) }
func lt(i Instruction, vm LuaVM) { _compare(i, vm, LUA_OPLT) }
func le(i Instruction, vm LuaVM) { _compare(i, vm, LUA_OPLE) }

func eqk(i Instruction, vm LuaVM) {
	a, b, _, k := i.ABC()
	a += 1
	vm.PushValue(a)
	vm.GetConst(b)
	if vm.Compare(-2, -1, LUA_OPEQ) != k {
		vm.AddPC(1)
	}
	vm.Pop(2)
}

func _compareI(i Instruction, vm LuaVM, op CompareOp) {
	a, b, _, k := i.ABC()
	a += 1
	sb := b - OFFSET_sC
	vm.PushValue(a)
	vm.PushInteger(int64(sb))
	if vm.Compare(-2, -1, op) != k {
		vm.AddPC(1)
	}
	vm.Pop(2)
}

func eqi(i Instruction, vm LuaVM) { _compareI(i, vm, LUA_OPEQ) }
func lti(i Instruction, vm LuaVM) { _compareI(i, vm, LUA_OPLT) }
func lei(i Instruction, vm LuaVM) { _compareI(i, vm, LUA_OPLE) }
func gti(i Instruction, vm LuaVM) { _compareI(i, vm, LUA_OPGT) }
func gei(i Instruction, vm LuaVM) { _compareI(i, vm, LUA_OPGE) }

func not(i Instruction, vm LuaVM) {
	a, b, _, _ := i.ABC()
	a += 1
	b += 1
	vm.PushBoolean(!vm.ToBoolean(b))
	vm.Replace(a)
}

func testSet(i Instruction, vm LuaVM) {
	a, b, _, k := i.ABC()
	a += 1
	b += 1
	if vm.ToBoolean(b) != k {
		vm.AddPC(1)
	} else {
		vm.Copy(b, a)
	}
}

func test(i Instruction, vm LuaVM) {
	a, _, _, k := i.ABC()
	a += 1
	if vm.ToBoolean(a) != k {
		vm.AddPC(1)
	}
}

func mmbin(i Instruction, vm LuaVM) {

}

func mmbini(i Instruction, vm LuaVM) {

}

func mmbink(i Instruction, vm LuaVM) {

}