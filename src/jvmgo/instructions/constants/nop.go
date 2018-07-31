package constants

import "jvmgo/rtda"
import "jvmgo/instructions/base"

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
}
