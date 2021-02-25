package vm

import "api"

type Instruction uint32

func (self Instruction) Opcode() int {
	return int(self & 0x7F)
}

func (self Instruction) ABC() (a, b, c int, k bool) {
	a = int(self >> 7 & 0xFF)
	b = int(self >> 16 & 0xFF)
	c = int(self >> 24 & 0xFF)
	k = self >> 15 & 0x1 == 1
	return
}

func (self Instruction) ABx() (a, bx int) {
	a = int(self >> 7 & 0xFF)
	bx = int(self >> 15)
	return
}

func (self Instruction) AsBx() (a, abx int) {
	a, bx := self.ABx()
	return a, bx - MAXARG_sBx
}

func (self Instruction) Ax() int {
	return int(self >> 7)
}

func (self Instruction) SJ() int {
	return int(self >> 7) - OFFSET_sJ
}

const SIZE_C = 8
const SIZE_B = 8
const SIZE_Bx = (SIZE_C + SIZE_B + 1)
const SIZE_A = 8
const SIZE_Ax = (SIZE_Bx + SIZE_A)
const SIZE_sJ = SIZE_Bx + SIZE_A

const MAXARG_Bx = 1 << SIZE_Bx - 1
const MAXARG_sBx = MAXARG_Bx>>1
const MAXARG_sJ = 1 << SIZE_sJ - 1
const OFFSET_sJ = MAXARG_sJ >> 1

func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].opMode
}

func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].argBMode
}

func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].argCMode
}

func (self Instruction) Execute(vm api.LuaVM) {
	action := opcodes[self.Opcode()].action
	if action != nil {
		action(self, vm)
	}else{
		panic(self.OpName())
	}
}