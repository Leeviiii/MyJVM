package main

import "fmt"
import "strings"
import "jvmgo/classpath"
import "jvmgo/classfile"

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
