package classfile

// SourceFileAttribute 可选定长属性
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FeilName() string {
	return self.cp.getUft8(self.sourceFileIndex)
}
