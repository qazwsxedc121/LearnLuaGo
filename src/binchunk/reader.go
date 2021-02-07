package binchunk

import (
	"encoding/binary"
	"math"
)

type reader struct {
	data []byte
}

func (self *reader) readByte() byte {
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

func (self *reader) readUint32() uint32 {
	i := binary.LittleEndian.Uint32(self.data)
	self.data = self.data[4:]
	return i
}

func (self *reader) readUint64() uint64 {
	i := binary.LittleEndian.Uint64(self.data)
	self.data = self.data[8:]
	return i
}

func (self *reader) readLuaInteger() int64 {
	return int64(self.readUint64())
}

func (self *reader) readLuaNumber() float64 {
	return math.Float64frombits(self.readUint64())
}

func (self *reader) readString() string {
	size := uint(self.readVarInt())
	bytes := self.readBytes(size - 1)
	ret := string(bytes)
	return ret
}

func (self *reader) readBytes(n uint) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

func (self *reader) checkHeader() {
	if string(self.readBytes(4)) != LUA_SIGNATURE {
		panic("not a precompiled chunk!")
	} else if self.readByte() != LUAC_VERSION {
		panic("version mismatch!")
	} else if self.readByte() != LUAC_FORMAT {
		panic("format mismatch")
	} else if string(self.readBytes(6)) != LUAC_DATA {
		panic("corrupted!")
		// } else if self.readByte() != CINT_SIZE {
		// 	panic("int size mismatch!")
		// } else if self.readByte() != CSIZET_SIZE {
		// 	panic("size_t size mismatch!")
	} else if self.readByte() != INSTRUCTION_SIZE {
		panic("instruction size mismatch!")
	} else if self.readByte() != LUA_INTEGER_SIZE {
		panic("lua_Integer size mismatch!")
	} else if self.readByte() != LUA_NUMBER_SIZE {
		panic("lua_Number size mismatch!")
	} else if self.readLuaInteger() != LUAC_INT {
		panic("endianness mismatch!")
	} else if self.readLuaNumber() != LUAC_NUM {
		panic("float format mismatch!")
	}
}

func (self *reader) readProto(parentSource string) *Prototype {
	source := self.readString()
	if source == "" {
		source = parentSource
	}
	v1 := uint32(self.readVarInt())
	v2 := uint32(self.readVarInt())
	v3 := self.readByte()
	v4 := self.readByte()
	v5 := self.readByte()
	v6 := self.readCode()
	v7 := self.readConstants()
	v8 := self.readUpvalues()
	return &Prototype{
		Source:          source,
		LineDefined:     v1,
		LastLineDefined: v2,
		NumParams:       v3,
		IsVararg:        v4,
		MaxStackSize:    v5,
		Code:            v6,
		Constants:       v7,
		Upvalues:        v8,
		Protos:          self.readProtos(source),
		LineInfo:        self.readLineInfo(),
		AbsLineInfos:    self.readAbsLineInfo(),
		LocVars:         self.readLocVars(),
		UpvalueNames:    self.readUpvalueNames(),
	}
}

func (self *reader) readCode() []uint32 {
	code := make([]uint32, self.readVarInt())
	for i := range code {
		code[i] = self.readUint32()
	}
	return code
}

func (self *reader) readVarInt() uint64 {
	b := self.readByte()
	var mask byte = 1 << 7
	count := 0
	d := make([]byte, 8)
	df := make([]byte, 8)
	for b&mask == 0 {
		d[count] = b
		count++
		b = self.readByte()
	}
	mask = 0x7F
	d[count] = b & mask
	for i := 0; i < count; i++ {
		df[7-i] = (d[count-i] >> i) | (d[count-1-i] << (7 - i))
	}
	df[7-count] = d[0] >> count
	ret := binary.BigEndian.Uint64(df)
	return ret
}

func (self *reader) readConstants() []interface{} {
	constants := make([]interface{}, self.readVarInt())
	for i := range constants {
		constants[i] = self.readConstant()
	}
	return constants
}

func (self *reader) readConstant() interface{} {
	switch self.readByte() {
	case TAG_NIL:
		return nil
	case TAG_TRUE:
		return true
	case TAG_FALSE:
		return false
	case TAG_INTEGER:
		return self.readLuaInteger()
	case TAG_NUMBER:
		return self.readLuaNumber()
	case TAG_SHORT_STR:
		return self.readString()
	case TAG_LONG_STR:
		return self.readString()
	default:
		panic("corrupted!")
	}
}

func (self *reader) readUpvalues() []Upvalue {
	upvalues := make([]Upvalue, self.readVarInt())
	for i := range upvalues {
		upvalues[i] = Upvalue{
			Instack: self.readByte(),
			Idx:     self.readByte(),
			Kind:    self.readByte(),
		}
	}
	return upvalues
}

func (self *reader) readProtos(parentSource string) []*Prototype {
	protos := make([]*Prototype, self.readVarInt())
	for i := range protos {
		protos[i] = self.readProto(parentSource)
	}
	return protos
}

func (self *reader) readLineInfo() []uint32 {
	sizeLine := self.readVarInt()
	lineInfo := make([]uint32, sizeLine)
	for i := range lineInfo {
		lineInfo[i] = uint32(self.readByte())
	}
	return lineInfo
}

func (self *reader) readAbsLineInfo() []AbsLineInfo {
	lineInfo := make([]AbsLineInfo, self.readVarInt())
	for i := range lineInfo {
		lineInfo[i] = AbsLineInfo{
			PC:   uint32(self.readVarInt()),
			Line: uint32(self.readVarInt()),
		}
	}
	return lineInfo
}

func (self *reader) readLocVars() []LocVar {
	locVars := make([]LocVar, self.readVarInt())
	for i := range locVars {
		locVars[i] = LocVar{
			VarName: self.readString(),
			StartPC: uint32(self.readVarInt()),
			EndPC:   uint32(self.readVarInt()),
		}
	}
	return locVars
}

func (self *reader) readUpvalueNames() []string {
	names := make([]string, self.readVarInt())
	for i := range names {
		names[i] = self.readString()
	}
	return names
}
