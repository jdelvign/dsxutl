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

// Package dsx contains code for the dsxutl command
package dsx

import (
	"bufio"
	"fmt"
	"os"
)

// CommandHeader ...
type CommandHeader struct{}

// Process ...
func (t *CommandHeader) Process() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: dsxutl header <DSXFILE>\n")
		os.Exit(1)
	} 

	dsxFileName := os.Args[len(os.Args)-1]

	f, r := openFile(dsxFileName)
	defer f.Close()

	scanner := bufio.NewScanner(r)
	buffer := make([]byte, bufferSize)
	scanner.Buffer(buffer, bufferSize)

	display := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == beginHeader {
			display = true
		}

		if display {
			fmt.Println(line) // Println will add back the final '\n'
		}

		if line == endHeader {
			display = false
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading dsx file: %e", err)
	}
}
