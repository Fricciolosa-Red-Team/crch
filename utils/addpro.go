package main

/*

cat subs-tmp.txt | removepro | addpro > subs-final.txt

*/

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	input := ScanTargets()

	f, err := os.Create("temp")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	for _, elem := range removeDuplicateValues(input) {
		_, err := f.WriteString("https://" + elem + "\n")
		if err != nil {
			panic(err)
		}
	}

}

//ScanTargets return the array of elements
//taken as input on stdin.
func ScanTargets() []string {

	var result []string
	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}
	return result
}

//removeDuplicateValues
func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
