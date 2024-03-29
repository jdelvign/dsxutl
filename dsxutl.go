/*
Copyright 2019 Jerome Delvigne

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package main provides the entry point of the 'dsxutl' commands
package main

import (
	"fmt"
	"os"

	"github.com/jdelvign/dsxutl/dsx"
)

// Function main do something with `dsxutl` Command
// Subcommands :
// 		`dsxutl grep -substr <substring> -dsxfile <dsxfile>` : find the substring inside the Job Designs
//		`dsxutl header -dsxfile <dsxfile>` : Print the DSX header
func main() {

	//start := time.Now()

	m := make(map[string]dsx.Command)
	m["grep"] = new(dsx.CommandGrep)
	m["header"] = new(dsx.CommandHeader)
	m["ljobs"] = new(dsx.CommandLJobs)

	usageMsg := `Usage: dsxutl <command> <options>
  dsxutl grep <options>
  dsxutl header <options>
  dsxutl ljobs <options>`

	if len(os.Args) < 2 {
		fmt.Println(usageMsg)
		os.Exit(1)
	}

	// Select the Subcommand
	cmd := m[os.Args[1]]

	if cmd != nil {
		cmd.Process()
	} else {
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println(usageMsg)
		os.Exit(1)
	}
}
