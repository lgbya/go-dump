package dump

import (
	"io/ioutil"
	"reflect"
	"testing"
)


type Demo struct {
	One
	two Two
}

type One struct {
	Number int
	List 	[][3]int
}

type Two struct {
	Number int
}

func TestVarDump(t *testing.T) {

	data := map[int8]map[string]map[string]Demo{
		1: {
			"a": {
				"a":  Demo{
					One:One{1, [][3]int{ {1, 2, 3}, {4, 5, 6} } },
					two:Two{1},
				},
				"a2":  Demo{
					One:One{1, [][3]int{ {1, 2, 3}, {4, 5, 6} }},
					two:Two{1},
				},
			},
		},
		2: {
			"b": {
				"b":   Demo{
					One:One{1, [][3]int{ {1, 2, 3}, {4, 5, 6} }},
					two:Two{1},
				},
				"b2":  Demo{
					One:One{1, [][3]int{ {1, 2, 3}, {4, 5, 6} }},
					two:Two{1},
				},
			},
		},
	}

	got := Format(data)
	wantByte, _ := ioutil.ReadFile("data.txt")
	want := string(wantByte)

	if !reflect.DeepEqual(want, got){
		t.Errorf("excepted:%v, got:%v", want, got)
	}


}

