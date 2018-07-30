package main

import "fmt"
import "strings"
import "jvmgo/classpath"
import "jvmgo/classfile"
import "jvmgo/rtda"

// ./jvmgo java.lang.Object
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.01")
	} else if cmd.helpFlag {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	cf := loadClass(className, cp)
	printClsasInfo(cf)
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}
func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	fmt.Println(vars.GetInt(0))
	fmt.Println(vars.GetInt(1))
	fmt.Println(vars.GetLong(2))
	fmt.Println(vars.GetLong(4))
	fmt.Println(vars.GetFloat(6))
	fmt.Println(vars.GetDouble(7))
	fmt.Println(vars.GetRef(9))
}
func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	fmt.Println(ops.PopRef())
	fmt.Println(ops.PopDouble())
	fmt.Println(ops.PopFloat())
	fmt.Println(ops.PopLong())
	fmt.Println(ops.PopLong())
	fmt.Println(ops.PopInt())
	fmt.Println(ops.PopInt())
}
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}
func printClsasInfo(cf *classfile.ClassFile) {
	fmt.Printf("version :%v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("accessFlags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("field count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("method count: %v\n", len(cf.Methods()))
	for _, f := range cf.Methods() {
		fmt.Printf(" %s\n", f.Name())
	}
}
