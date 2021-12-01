package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
  "path"
)

func Check(e error) {
  if e != nil {
    panic(e);
  }
}

func ReadLines(path string) []string {
  content, err := ioutil.ReadFile(path)
  Check(err);

 return strings.Split(string(content), "\n")
}

func ReadLinesAsIntegers(p string) []int {
  lines := ReadLines(path.Join("files", p))
  ints := make([]int, len(lines))

  for i, s := range lines {
    ints[i], _ = strconv.Atoi(s)
  }

  return ints
}
