package classfile

// DeprecatedAttribute 标记属性
type DeprecatedAttribute struct {
	MakerAttribute
}

// SyntheticAttribute 标记属性
type SyntheticAttribute struct {
	MakerAttribute
}

type MakerAttribute struct{}

// Deprecated 和 Synthetic 都是只起标记作用 不包含任何数据 所以读取为空
func (self *MakerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
