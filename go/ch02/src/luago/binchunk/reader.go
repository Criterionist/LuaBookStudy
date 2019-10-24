package binchunk

type reader struct {
	data []byte //要被解析的chunk数据
}

func (self *reader) readByte() byte {
	b := self.data[0]
	self.data = self.data[1:]
	return b
} //读取一个字节
