package main

import "testing"

func TestFirst(t *testing.T) {
	say:=First()
	if say!="Hello,World!github-acion-demo"{
		t.Errorf("method:First() return error,result:%s",say)
	}else{
		t.Log("success!")
	}
}
