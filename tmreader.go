// Copyright 2011 The GoGL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.mkd file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	tmEmptyOrCommentRE = regexp.MustCompile("^[ \\t]*(#.*)?$")
	tmTypePairRE       = regexp.MustCompile("^([_A-Za-z0-9]+),\\*,\\*,[\\t ]*([A-Za-z0-9\\*_ ]+),\\*,\\*")
)

func ReadTypeMapFromFile(name string) (TypeMap, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadTypeMap(file)
}

func ReadTypeMap(r io.Reader) (TypeMap, error) {
	tm := make(TypeMap)
	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, "\n")
		//fmt.Println(line)

		if tmEmptyOrCommentRE.MatchString(line) {
			// skip
		} else if typePair := tmTypePairRE.FindStringSubmatch(line); typePair != nil {
			tm[typePair[1]] = typePair[2]
		} else {
			fmt.Fprintf(os.Stderr, "WARNING: Unable to parse line: %v\n", line)
		}
	}

	return tm, nil
}
