//
// description.go
// Copyright(c)2014-2015 Google, Inc.
//
// This file is part of skicka.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"os"
)

func description(args []string) int {
	if len(args) < 2 {
		fmt.Printf("Usage: skicka desc drive_path description_text\n")
		fmt.Printf("Run \"skicka help\" for more detailed help text.\n")
		return 1
	}

	errs := 0
	fn := args[0]
        text := args[1]
	file, err := gd.GetFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "skicka: %s: %v\n", fn, err)
		errs++
	}
	if err := gd.UpdateDescription(file, text); err != nil {
		fmt.Fprintf(os.Stderr, "skicka: %s: %v\n", fn, err)
		errs++
	}
	return errs
}
