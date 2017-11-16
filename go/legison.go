package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	/* i lik' to move it mov,it
	 * i lik' to move it mov,it
	 * ya like to
	 * MOVEIT!
	 */

	/*  argument from the sea~ */
	arg := os.Args[1:]

	/*  make dir output/	*/
	err := os.MkdirAll("output", os.FileMode(0777))
	if err != nil {
		panic(err)
	}

	t0 := time.Now()

	/*  foreach arg[], load file, to .json and fireouttt!!	*/
	for i := range arg {
		fmt.Println(arg[i] + "  (_'')")
		legis, err := ioutil.ReadFile(arg[i])
		if err != nil {
			panic(err)
		}

		jsonout, name := legisToJSON(legis)
		if name == " uid" {
			name = "unidentify" + time.Since(t0).String()
		}

		err = ioutil.WriteFile("output/"+name+".json", jsonout, 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println(name + "   exported.")
		fmt.Println()
	}

	elapsed := time.Since(t0)
	fmt.Println("-----------------\n(''_) ", elapsed)
}

/*  input: []byte legis
    output: []byte .json
    thi, man catch 'em
*/
func legisToJSON(file []byte) ([]byte, string) {
	legis := *catchit(file)
	output, err := json.Marshal(legis)
	if err != nil {
		panic(err)
	}
	a := "uid"
	if len(legis.PassDate)-4 >= 0 {
		a = legis.PassDate[len(legis.PassDate)-4:]
	}
	return output, legis.Name + " " + a
}
