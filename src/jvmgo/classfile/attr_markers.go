package classfile

type DeprecatedAttribute struct{ MakerAttribute }
type SyntheicAttribute struct{ MakerAttribute }
type MakerAttribute struct {
}

func (self *MakerAttribute) readInfo(reader *ClassReader) {

}
