package classfile

import "fmt"

// go 语言的访问控制 大写的类型、结构体、字段、变量、函数、方法均为公开，小写即为私有
// 下面声明的结构体中为Java虚拟机规范规定的Class文件格式
type ClassFile struct {
	// 魔数
	//magic	uint32
	// 次版本号
	minorVersion uint16
	// 主版本号
	majorVersion uint16
	// 常量池
	constantPool ConstantPool
	// 类访问标志
	accessFlags uint16
	// 类
	thisClass uint16
	// 超类
	superClass uint16
	// 接口索引表
	interfaces []uint16
	// 字段
	fields []*MemberInfo
	// 方法表
	methods []*MemberInfo
	// 属性表
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	//TODO NO.001 go 没有异常处理机制 只有panic-recover机制 这里需要填坑
	// 填坑信息请参看README 填坑部分
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	// 创建一个ClassReader实例并传入参数classData
	cr := &ClassReader{classData}
	// 创建一个ClassFile实例
	cf = &ClassFile{}
	// 调用read方法读取classData
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// 很多文件格式都会规定满足该格式的文件必须以某几个固定字节开头，这几个字节起表示作用，叫做魔数
// class文件的魔数是"0xCAFEBABE" 咖啡宝贝
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		// JVM虚拟机规范需要在不符合class文件格式时抛出此异常 现在这个toyJVM还没有能力抛出 故使用panic终止程序
		panic("java.lang.ClassFormatError： magic!")
	}
}

// Java8的主版本号为52 SE8支持的版本号为45.0-52.0的class文件 不在此支持范围内，虚拟机就会抛出Java.lang.UnsupportedClassVersionError
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 次版本在前 主版本在后
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		// 1.2之前的主版本号均为45，且1.2之前有使用次版本号所以当主版本号为45时不做判断
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// getter
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

// getter
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

// getter
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

// getter
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

// getter
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

// getter
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

// 从常量池查找className
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// 从常量池查找类名superClassName
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有java.lang.Object没有超类
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
