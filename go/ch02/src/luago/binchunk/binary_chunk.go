package binchunk
import "fmt"

type binaryChunk struct {
	header
	sizeUpvalues byte
	mainFunc *Prototype
}

type header struct {
	signature [4]byte
	version byte
	format byte
	luacData [6]byte
	instructionSize byte
	luaIntegerSize byte
	luaNumberSize byte
	luacInt int64
	luacNum float64
}

const (
	LUA_SIGNATURE = "\x1bLua"
	LUAC_VERSION = 0x54
	LUAC_FORMAT = 0
	LUAC_DATA = "\x19\x93\r\n\x1a\n"
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE = 8
	LUAC_INT = 0x5678
	LUAC_NUM = 370.5
)

type Prototype struct {
	Source string
	LineDefined uint32
	LastLineDefined uint32
	NumParams byte
	IsVararg byte
	MaxStackSize byte
	Code []uint32
	Constants []interface{}
	Upvalues []Upvalue
	Protos []*Prototype
	LineInfo []uint32
	AbsLineInfos []AbsLineInfo
	LocVars []LocVar
	UpvalueNames []string
}

const (
	TAG_NIL = 0x00
	TAG_FALSE = 0x01
	TAG_TRUE = 0x11
	TAG_NUMBER = 0x03
	TAG_INTEGER = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR = 0x14
)

type Upvalue struct {
	Instack byte
	Idx byte
	Kind byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC uint32
}

type AbsLineInfo struct {
	PC uint32
	Line uint32
}

func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte()
	return reader.readProto("")
}

func readVarIntS(source string, v uint64) {
	data := []byte(source)
	rd := &reader{data}
	a := rd.readVarInt()
	var t byte = 2
	fmt.Printf("%b %v %v %v\n", a == v, a, data, t << 7)
	return 
}

func Test() {
	readVarIntS("\x02\xac", 300)
	readVarIntS("\x81", 1)
}