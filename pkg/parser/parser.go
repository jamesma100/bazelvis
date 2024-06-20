package parser

import "strings"
import "bufio"
import "log"

func trimTarget(s string) string {
	trim := strings.TrimSpace(s)
	unquote := trim[1 : len(trim)-1]
	return unquote
}

func PrintMap(m map[string][]string) {
	log.Println("[DEBUG] output map:")
	for k, v := range m {
		log.Println("[DEBUG] key:", k)
		for _, i := range v {
			log.Println("[DEBUG] value:", i)
		}
	}
}

func BuildMap(contents string) map[string][]string {
	scanner := bufio.NewScanner(strings.NewReader(contents))

	m := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "->")
		if len(parts) == 1 {
			continue
		} else {
			src := trimTarget(parts[0])
			deps := strings.Split(trimTarget(parts[1]), "\\n")
			for _, dep := range deps {
				m[src] = append(m[src], dep)
			}
		}
	}
	return m
}
