package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	args := append([]string{"get"}, os.Args[1:]...)
	args = append(args, "-o=json")
	cmd := exec.Command("kubectl", args...)
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		fmt.Printf("%v %v\n", cmd.Stderr, cmd.Stdout)
		panic(err)
	}
	data, _ := ioutil.ReadAll(stdout)
	cmd.Wait()

	var obj interface{}
	json.Unmarshal(data, &obj)
	items := obj.(map[string]interface{})["items"].([]interface{})
	names := make([]string, len(items))
	for ix := range items {
		item := items[ix]
		name := item.(map[string]interface{})["metadata"].(map[string]interface{})["name"]
		names[ix] = name.(string)
		os.Stderr.Write([]byte(fmt.Sprintf("%d] %s\n", (ix + 1), name)))
	}

	input := bufio.NewReader(os.Stdin)
	os.Stderr.Write([]byte("Select an item: "))
	selection, _ := input.ReadString('\n')
	ix, _ := strconv.Atoi(strings.TrimSpace(selection))
	ix = ix - 1
	if ix < 0 || ix >= len(items) {
		panic("invalid index")
	}
	fmt.Printf("%s\n", names[ix])	
}
