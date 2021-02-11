package vm

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
	return int(self >> 7)
}

const MAXARG_Bx = 1<<17 - 1
const MAXARG_sBx = MAXARG_Bx

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