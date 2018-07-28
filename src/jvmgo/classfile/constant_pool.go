package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //long 跟 double占常量池的两个位置
		}
	}
	return cp
}
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constatnt pool index")
}
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	utInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(utInfo.nameIndex)
	_type := self.getUtf8(utInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classinfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classinfo.nameIndex)
}
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
