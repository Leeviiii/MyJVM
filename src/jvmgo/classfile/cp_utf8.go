package classfile

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// 简化版 字符串中不可以包含null以及补充字符
// 否则需要参考java.io.DataInputStream.readUTF() 方法
// 因为java用的是MUTF8编码
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
