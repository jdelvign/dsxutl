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
	"strings"
)

// CommandLJobs ...
type CommandLJobs struct{}

// Process ...
func (t CommandLJobs) Process() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: dsxutl ljobs <DSXFILE>\n")
		os.Exit(1)
	}

	dsxFileName := os.Args[len(os.Args)-1]

	f, r := openFile(dsxFileName)
	defer f.Close()

	scanner := bufio.NewScanner(r)
	buffer := make([]byte, bufferSize)
	scanner.Buffer(buffer, bufferSize)

	dsjob := false
	dsroutines := false
	dsrecord := false
	dsProject := "<not available>"
	dsJobName := "<not available>"
	dsRoutineName := "<not available>"
	dsCategory := "<not available>"

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, toolInstanceID) {
			dsProject = strings.Split(line, "\"")[1]
		}

		if line == beginDSJOB {
			dsjob = true
		}

		if line == beginDSROUTINES {
			dsroutines = true
		}

		if dsjob {
			if strings.HasPrefix(line, dsjobIDENTIFIER) {
				dsJobName = strings.Split(line, "\"")[1]
			}
			if strings.HasPrefix(line, dsjobCATEGORY) {
				dsCategory = strings.Split(line, "\"")[1]
			}
		}

		if dsroutines {
			if line == beginDSRECORD {
				dsrecord = true
			}
		}

		if dsroutines && dsrecord {
			if strings.HasPrefix(line, dsroutineIDENTIFIER) {
				dsRoutineName = strings.Split(line, "\"")[1]
			}
			if strings.HasPrefix(line, dsroutineCATEGORY) {
				dsCategory = strings.Split(line, "\"")[1]
			}
			if line == endDSRECORD {
				// Print routine info now !
				fmt.Printf("%s\t%s\t%s\n", dsProject, dsRoutineName, dsCategory)

				dsrecord = false
				dsRoutineName = "<not available>"
			}
			if line == endDSROUTINES {
				dsroutines = false
			}
		}

		if line == endDSJOB {
			// Print job info now !
			fmt.Printf("%s\t%s\t%s\n", dsProject, dsJobName, dsCategory)

			dsjob = false
			dsJobName = "<not available>"
			dsCategory = "<not available>"
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading dsx file: %e", err)
	}

}
