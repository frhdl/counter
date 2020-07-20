// Modification of: https://github.com/adonovan/gopl.io/blob/master/ch1/dup3/main.go

// Counter prints ordered count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// Command .
type Command struct {
	Name  string
	Count int
}

// Input .
type Input struct {
	commands []Command
}

func (c Command) String() string {
	return fmt.Sprintf("%d : %s", c.Count, c.Name)
}

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "counter: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	var input Input

	for line, n := range counts {
		if n > 5 {
			input.commands = append(input.commands, Command{
				Name:  line,
				Count: n,
			})
		}
	}

	sort.Slice(input.commands, func(i, j int) bool {
		return input.commands[i].Count > input.commands[j].Count
	})

	for _, command := range input.commands {
		fmt.Println(command.String())
	}
}
