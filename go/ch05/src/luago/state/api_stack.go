package state

//基础栈操纵方法

func (self *luaState) GetTop() int {
	return self.stack.top
} //返回栈顶索引

func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
} //返回绝对索引

func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true //不考虑失败情况
} //检测容量并视情况扩容

func (self *luaState) Pop(n int) {
	//for i := 0; i < n; i++ {
	//	self.stack.pop()
	//}
	self.SetTop(-n - 1)
} //弹栈

func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
} //复制

func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
} //把索引处值压到栈顶

func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
} //将栈顶写入索引处

func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
} //将栈顶插入索引处

func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
} //移除索引处值

func (self *luaState) Rotate(idx, n int) {
	t := self.stack.top - 1
	p := self.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	self.stack.reverse(p, m)
	self.stack.reverse(m+1, t)
	self.stack.reverse(p, t)
} //旋转

func (self *luaState) SetTop(idx int) {
	newTop := self.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}
	n := self.stack.top - newTop
	if n > 0 { //减少
		for i := 0; i < n; i++ {
			self.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i > n; i-- {
			self.stack.push(nil)
		}
	}
} //设置top
