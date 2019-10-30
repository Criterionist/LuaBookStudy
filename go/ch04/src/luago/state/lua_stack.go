package state

type luaStack struct {
	slots []luaValue
	top   int
} //Lua栈

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
} //创建栈

func (self *luaStack) check(n int) {
	free := len(self.slots) - self.top
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil)
	}
} //检测栈的剩余空间是否能容纳n个值，如不能则扩容

func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots) {
		panic("stack overflow!")
	}
	self.slots[self.top] = val
	self.top++
} //压栈

func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("stack underflow!")
	}
	self.top--
	val := self.slots[self.top]
	self.slots[self.top] = nil
	return val
} //弹栈

func (self *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return idx + self.top + 1
} //将索引转成绝对索引

func (self *luaStack) isValid(idx int) bool {
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx <= self.top
} //判断索引合法性

func (self *luaStack) get(idx int) luaValue {
	absIdx := self.absIndex(idx)
	if self.isValid(absIdx) {
		return self.slots[absIdx-1]
	}
	return nil
} //根据索引取值

func (self *luaStack) set(idx int, val luaValue) {
	absIdx := self.absIndex(idx)
	if self.isValid(absIdx) {
		self.slots[absIdx-1] = val
		return
	}
	panic("invalid index!")
} //根据索引写值

func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
} //反转
