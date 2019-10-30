package binchunk

type binarychunk struct {
	header                  //头部
	sizeUpvalues byte       //主函数upvalue变量
	mainFunc     *Prototype //主函数原型
}

type header struct {
	signature       [4]byte //签名
	version         byte    //版本号
	format          byte    //格式号
	luacData        [6]byte //LUAC_DATA
	cintSize        byte    //cint长度
	sizetSize       byte    //size_t长度
	instructionSize byte    //Lua虚拟机指令宽度
	luaIntegerSize  byte    //Lua整数长度
	luaNumberSize   byte    //Lua浮点数长度
	luacInt         int64   //LUAC_INT
	luacNum         float64 //LUAC_NUM
}

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSIZET_SIZE      = 4
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

type Prototype struct {
	Source          string        //源文件名
	LineDefined     uint32        //起始行号
	LastLineDefined uint32        //终止行号
	NumParams       byte          //固定参数个数
	IsVararg        byte          //是否是变参函数
	MaxStackSize    byte          //寄存器数量
	Code            []uint32      //指令表
	Constants       []interface{} //常量表
	Upvalues        []Upvalue     //Upvalue表
	Protos          []*Prototype  //子函数原型表
	LineInfo        []uint32      //行号表
	LocVars         []LocVar      //局部变量表
	UpvalueNames    []string      //Upvalue名列表
} //函数原型

const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

type Upvalue struct {
	Instack byte
	Idx     byte
}

type LocVar struct {
	VarName string //变量名
	StartPC uint32 //起始指令索引
	EndPC   uint32 //终止指令索引
} //局部变量

func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()        //校验头部
	reader.readByte()           //跳过Upvalue变量
	return reader.readProto("") //读取函数原型
} //解析二进制chunk
