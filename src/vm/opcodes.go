package vm

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
}

var opcodes = []opcode{
	/*     T  A  B       C       mode   name */
	opcode{0, 1, OpArgR, OpArgN, IABC, "MOVE"},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "LOADI"},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "LOADF"},
	opcode{0, 1, OpArgK, OpArgN, IABx, "LOADK"},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADKX"},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADFALSE"},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADFALSESKIP"},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADTRUE"},
	opcode{0, 1, OpArgU, OpArgN, IABC, "LOADNIL"},
	opcode{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL"},
	opcode{0, 1, OpArgU, OpArgK, IABC, "GETTABUP"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "GETTABLE"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "GETI"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "GETFIELD"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SETTABUP"},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETTABLE"},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETI"},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETFIELD"},
	opcode{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "SELF"},
	opcode{0, 1, OpArgR, OpArgU, IABC, "ADDI"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "ADDK"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "SUBK"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "MULK"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "MODK"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "POWK"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "DIVK"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "IDIVK"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "BANDK"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "BORK"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "BXORK"},
	opcode{0, 1, OpArgR, OpArgU, IABC, "SHRI"},
	opcode{0, 1, OpArgR, OpArgU, IABC, "SHLI"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "ADD"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "SUB"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "MUL"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "MOD"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "POW"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "DIV"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "IDIV"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "BAND"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "BOR"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "BXOR"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "SHL"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "SHR"},
	opcode{0, 0, OpArgR, OpArgU, IABC, "MMBIN"},
	opcode{0, 0, OpArgU, OpArgU, IABC, "MMBINI"},
	opcode{0, 0, OpArgR, OpArgK, IABC, "MMBINK"},
	opcode{0, 1, OpArgR, OpArgN, IABC, "UNM"},
	opcode{0, 1, OpArgR, OpArgN, IABC, "BNOT"},
	opcode{0, 1, OpArgR, OpArgN, IABC, "NOT"},
	opcode{0, 1, OpArgR, OpArgN, IABC, "LEN"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "CONCAT"},
	opcode{0, 0, OpArgN, OpArgN, IAx, "CLOSE"},
	opcode{0, 0, OpArgN, OpArgN, IAx, "TBC"},
	opcode{0, 0, OpArgR, OpArgN, IsJ, "JMP"},
	opcode{0, 0, OpArgR, OpArgN, IABC, "EQ"},
	opcode{0, 0, OpArgR, OpArgN, IABC, "LT"},
	opcode{0, 0, OpArgR, OpArgN, IABC, "LE"},
	opcode{0, 0, OpArgK, OpArgN, IABC, "EQK"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "EQI"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "LTI"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "LEI"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "GTI"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "GEI"},
	opcode{0, 0, OpArgN, OpArgU, IABC, "TEST"},
	opcode{0, 1, OpArgR, OpArgU, IABC, "TESTSET"},
	opcode{0, 1, OpArgU, OpArgU, IABC, "CALL"},
	opcode{0, 1, OpArgU, OpArgU, IABC, "TAILCALL"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "RETURN"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "RETURN0"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "RETURN1"},
	opcode{0, 1, OpArgR, OpArgN, IABx, "FORLOOP"},
	opcode{0, 1, OpArgR, OpArgN, IABx, "FORPREP"},
	opcode{0, 1, OpArgR, OpArgN, IABx, "TFORPREP"},
	opcode{0, 0, OpArgN, OpArgU, IABC, "TFORCALL"},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP"},
	opcode{0, 0, OpArgU, OpArgU, IABC, "SETLIST"},
	opcode{0, 1, OpArgU, OpArgN, IABx, "CLOSURE"},
	opcode{0, 1, OpArgU, OpArgN, IABC, "VARARG"},
	opcode{0, 1, OpArgU, OpArgN, IABC, "VARARGPREP"},
	opcode{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG"},

}