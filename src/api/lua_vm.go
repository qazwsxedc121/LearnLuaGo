package api

type LuaVM interface {
	LuaState
	PC() int
	AddPC(n int)
	Fetch() uint32
	GetConst(idx int)
	GetRK(r int, k bool)

	RegisterCount() int
	LoadVararg(n int)
	LoadProto(idx int)
}


