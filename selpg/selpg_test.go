package main

import (
	"os/exec"
	"testing"
)

func TestSelpg(t *testing.T) {
	stdout, err := exec.Command("bash", "-c", "./selpg -s1 -e1 t.txt").Output()
	stdout2, err2 := exec.Command("bash", "-c", "./a.out -s1 -e1 t.txt").Output()
	if err != err2 {
		t.Error(err)
	}
	//fmt.Println(stdout)
	//fmt.Println(stdout2)
	for i := 0; i < len(stdout); i++ {
		if stdout[i] != stdout2[i] {
			t.Error("not fit")
			break
		}
	}
	if err != nil {
		t.Error(err)
	}
}
