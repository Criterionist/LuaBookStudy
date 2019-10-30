package api

type LuaType = int

type LuaState interface {
	/* 栈的基本操作 */
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx int, toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx, n int)
	SetTop(idx int)
	/* Lua栈->Go语言的接口函数 */
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)
	/* Go语言->Lua栈的接口函数 */
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)
} //LuaState接口
