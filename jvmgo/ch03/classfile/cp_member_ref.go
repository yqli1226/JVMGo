package classfile

// ConstantMemberRefInfo 成员符号引用
type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberRefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct{ ConstantMemberRefInfo }
type ConstantMethodRefInfo struct{ ConstantMemberRefInfo }
type ConstantInterfaceMethodRefInfo struct{ ConstantMemberRefInfo }
