package classfile

// ConstantValueAttribute
// constantValue是定长属性，只存在在field_info结构中,用于表达常量表达式的值
//
//	constantValue 结构定义为
//
//	ConstantValue_attribute {
//		u2 	attribute_name_index;
//		u4	attribute_length;  // 此值必须为2
//		u2	constantvalue_index;  //常量池索引
//	}
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
