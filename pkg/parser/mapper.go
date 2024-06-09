package parser

import "fmt"
import "strings"
import "bufio"
// import "os"
// import "io/ioutil"


func trimTarget(s string) string {
  trim := strings.TrimSpace(s)
  unquote := trim[1:len(trim)-1]
  return unquote
}

func BuildMap(contents string) map[string][]string {
  //arg := os.Args[1]
  //fmt.Printf("Received arg %s\n", arg)
  // contents, err := ioutil.ReadFile("../../example.txt")
  // if err != nil {
  //   fmt.Println(err)
  // }
  // stdout := string(contents)
  scanner := bufio.NewScanner(strings.NewReader(contents))

	m := make(map[string][]string)

  for scanner.Scan() {
    line := scanner.Text()
    parts := strings.Split(line, "->")
    if len(parts) == 1 {
      continue
    } else {
      src := trimTarget(parts[0])
      dep := trimTarget(parts[1])
      m[src] = append(m[src], dep)
    }
  }
  fmt.Println("map:")
  for k, v := range m {
    fmt.Println("Key:", k)
    for _, i := range v {
      fmt.Println("Value:", i)
    }
  }
  return m
}
