package classfile

// 这里的utf8我一直很迷，java什么时候有这么号类型了，其实这是在常量池的一个常量项 存储的是字符串的字面量
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// TODO 填坑 NO.002 此处为简易版 只要字符串不包含null或者补充字符即可正常工作
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
