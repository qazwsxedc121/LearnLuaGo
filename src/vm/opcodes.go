package vm

import "api"

const (
	IABC = iota
	IABx
	IAsBx
	IAx
	IsJ
)

const (
	OP_MOVE = iota
	OP_LOADI
	OP_LOADF
	OP_LOADK
	OP_LOADKX
	OP_LOADFALSE
	OP_LFALSESKIP
	OP_LOADTRUE
	OP_LOADNIL
	OP_GETUPVAL
	OP_SETUPVAL
	OP_GETTABUP
	OP_GETTABLE
	OP_GETI
	OP_GETFIELD
	OP_SETTABUP
	OP_SETTABLE
	OP_SETI
	OP_SETFIELD
	OP_NEWTABLE
	OP_SELF
	OP_ADDI
	OP_ADDK
	OP_SUBK
	OP_MULK
	OP_MODK
	OP_POWK
	OP_DIVK
	OP_IDIVK
	OP_BANDK
	OP_BORK
	OP_BXORK
	OP_SHRI
	OP_SHLI
	OP_ADD
	OP_SUB
	OP_MUL
	OP_MOD
	OP_POW
	OP_DIV
	OP_IDIV
	OP_BAND
	OP_BOR
	OP_BXOR
	OP_SHL
	OP_SHR
	OP_MMBIN
	OP_MMBINI
	OP_MMBINK
	OP_UNM
	OP_BNOT
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_CLOSE
	OP_TBC
	OP_JMP
	OP_EQ
	OP_LT
	OP_LE
	OP_EQK
	OP_EQI
	OP_LTI
	OP_LEI
	OP_GTI
	OP_GEI
	OP_TEST
	OP_TESTSET
	OP_CALL
	OP_TAILCALL
	OP_RETURN
	OP_RETURN0
	OP_RETURN1
	OP_FORLOOP
	OP_FORPREP
	OP_TFORPREP
	OP_TFORCALL
	OP_TFORLOOP
	OP_SETLIST
	OP_CLOSURE
	OP_VARARG
	OP_VARARGPREP
	OP_EXTRAARG
)

const (
	OpArgN = iota;
	OpArgU
	OpArgR
	OpArgK
)

type opcode struct {
	testFlag byte
	setAFlag byte
	argBMode byte
	argCMode byte
	opMode byte
	name string
	action func(i Instruction, vm api.LuaVM)
}

var opcodes = []opcode{
	/*     T  A  B       C       mode   name */
	opcode{0, 1, OpArgR, OpArgN, IABC, "MOVE", move},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "LOADI", loadI},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "LOADF", loadF},
	opcode{0, 1, OpArgK, OpArgN, IABx, "LOADK", loadK},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADKX", loadKx},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADFALSE", loadFalse},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADFALSESKIP", loadFalseSkip},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADTRUE", loadTrue},
	opcode{0, 1, OpArgU, OpArgN, IABC, "LOADNIL", loadNil},
	opcode{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL", nil},
	opcode{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL", nil},
	opcode{0, 1, OpArgU, OpArgK, IABC, "GETTABUP", getTabup},
	opcode{0, 1, OpArgR, OpArgK, IABC, "GETTABLE", getTable},
	opcode{0, 1, OpArgR, OpArgR, IABC, "GETI", getI},
	opcode{0, 1, OpArgR, OpArgK, IABC, "GETFIELD", getField},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SETTABUP", setTabup},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETTABLE", setTable},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETI", setI},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETFIELD", setField},
	opcode{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE", newTable},
	opcode{0, 1, OpArgR, OpArgK, IABC, "SELF", nil},
	opcode{0, 1, OpArgR, OpArgU, IABC, "ADDI", addi},
	opcode{0, 1, OpArgR, OpArgK, IABC, "ADDK", addk},
	opcode{0, 1, OpArgR, OpArgK, IABC, "SUBK", subk},
	opcode{0, 1, OpArgR, OpArgK, IABC, "MULK", mulk},
	opcode{0, 1, OpArgR, OpArgK, IABC, "MODK", modk},
	opcode{0, 1, OpArgR, OpArgK, IABC, "POWK", powk},
	opcode{0, 1, OpArgR, OpArgK, IABC, "DIVK", divk},
	opcode{0, 1, OpArgR, OpArgK, IABC, "IDIVK", idivk},
	opcode{0, 1, OpArgR, OpArgK, IABC, "BANDK", bandk},
	opcode{0, 1, OpArgR, OpArgK, IABC, "BORK", bork},
	opcode{0, 1, OpArgR, OpArgK, IABC, "BXORK", bxork},
	opcode{0, 1, OpArgR, OpArgU, IABC, "SHRI", shri},
	opcode{0, 1, OpArgR, OpArgU, IABC, "SHLI", shli},
	opcode{0, 1, OpArgR, OpArgR, IABC, "ADD", add},
	opcode{0, 1, OpArgR, OpArgR, IABC, "SUB", sub},
	opcode{0, 1, OpArgR, OpArgR, IABC, "MUL", mul},
	opcode{0, 1, OpArgR, OpArgR, IABC, "MOD", mod},
	opcode{0, 1, OpArgR, OpArgR, IABC, "POW", pow},
	opcode{0, 1, OpArgR, OpArgR, IABC, "DIV", div},
	opcode{0, 1, OpArgR, OpArgR, IABC, "IDIV", idiv},
	opcode{0, 1, OpArgR, OpArgR, IABC, "BAND", band},
	opcode{0, 1, OpArgR, OpArgR, IABC, "BOR", bor},
	opcode{0, 1, OpArgR, OpArgR, IABC, "BXOR", bxor},
	opcode{0, 1, OpArgR, OpArgR, IABC, "SHL", shl},
	opcode{0, 1, OpArgR, OpArgR, IABC, "SHR", shr},
	opcode{0, 0, OpArgR, OpArgU, IABC, "MMBIN", mmbin},
	opcode{0, 0, OpArgU, OpArgU, IABC, "MMBINI", mmbini},
	opcode{0, 0, OpArgR, OpArgK, IABC, "MMBINK", mmbink},
	opcode{0, 1, OpArgR, OpArgN, IABC, "UNM", nil},
	opcode{0, 1, OpArgR, OpArgN, IABC, "BNOT", bnot},
	opcode{0, 1, OpArgR, OpArgN, IABC, "NOT", not},
	opcode{0, 1, OpArgR, OpArgN, IABC, "LEN", len},
	opcode{0, 1, OpArgR, OpArgR, IABC, "CONCAT", concat},
	opcode{0, 0, OpArgN, OpArgN, IAx, "CLOSE", nil},
	opcode{0, 0, OpArgN, OpArgN, IAx, "TBC", nil},
	opcode{0, 0, OpArgR, OpArgN, IsJ, "JMP", jmp},
	opcode{0, 0, OpArgR, OpArgN, IABC, "EQ", eq},
	opcode{0, 0, OpArgR, OpArgN, IABC, "LT", lt},
	opcode{0, 0, OpArgR, OpArgN, IABC, "LE", le},
	opcode{0, 0, OpArgK, OpArgN, IABC, "EQK", eqk},
	opcode{0, 0, OpArgU, OpArgN, IABC, "EQI", eqi},
	opcode{0, 0, OpArgU, OpArgN, IABC, "LTI", lti},
	opcode{0, 0, OpArgU, OpArgN, IABC, "LEI", lei},
	opcode{0, 0, OpArgU, OpArgN, IABC, "GTI", gti},
	opcode{0, 0, OpArgU, OpArgN, IABC, "GEI", gei},
	opcode{0, 0, OpArgN, OpArgU, IABC, "TEST", test},
	opcode{0, 1, OpArgR, OpArgU, IABC, "TESTSET", testSet},
	opcode{0, 1, OpArgU, OpArgU, IABC, "CALL", call},
	opcode{0, 1, OpArgU, OpArgU, IABC, "TAILCALL", nil},
	opcode{0, 0, OpArgU, OpArgN, IABC, "RETURN", nil},
	opcode{0, 0, OpArgU, OpArgN, IABC, "RETURN0", nil},
	opcode{0, 0, OpArgU, OpArgN, IABC, "RETURN1", nil},
	opcode{0, 1, OpArgR, OpArgN, IABx, "FORLOOP", forLoop},
	opcode{0, 1, OpArgR, OpArgN, IABx, "FORPREP", forPrep},
	opcode{0, 1, OpArgR, OpArgN, IABx, "TFORPREP", nil},
	opcode{0, 0, OpArgN, OpArgU, IABC, "TFORCALL", nil},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP", nil},
	opcode{0, 0, OpArgU, OpArgU, IABC, "SETLIST", setList},
	opcode{0, 1, OpArgU, OpArgN, IABx, "CLOSURE", nil},
	opcode{0, 1, OpArgU, OpArgN, IABC, "VARARG", nil},
	opcode{0, 1, OpArgU, OpArgN, IABC, "VARARGPREP", varargprep},
	opcode{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG", extraarg},

}