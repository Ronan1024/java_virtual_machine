// Package main
// @author: Ronan
// @since: 2023/8/4
// @desc:
package main

import (
	"fmt"
	"java_virtual_machine/classpath"
	"strings"
)

func main() {
	cmd := &CMD{
		helpFlag:    false,
		versionFlag: false,
		cpOption:    "",
		XjreOption:  "/Library/Java/JavaVirtualMachines/jdk1.8.0_271.jdk/Contents/Home/jre",
		class:       "java.lang.Object",
		args:        nil,
	}
	fmt.Println(cmd)
	if cmd.versionFlag {
		fmt.Println("version 1.0.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}

}

func startJVM(cmd *CMD) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
