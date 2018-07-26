package main

import "fmt"
import "strings"
import "jvmgo/classpath"

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
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not found or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
