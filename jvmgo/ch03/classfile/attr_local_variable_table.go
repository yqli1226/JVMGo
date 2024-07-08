package classfile

// LocalVariableTableAttribute
// 存放方法的局部变量信息
type LocalVariableTableAttribute struct {
	LocalVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc       uint16
	localVariable uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.LocalVariableTable = make([]*LocalVariableTableEntry, lineNumberTableLength)
	for i := range self.LocalVariableTable {
		self.LocalVariableTable[i] = &LocalVariableTableEntry{
			startPc:       reader.readUint16(),
			localVariable: reader.readUint16(),
		}
	}
}
