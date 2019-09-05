package main

import (
	"unsafe"

	"github.com/YukiMiyatake/GOSICK/study/MyFile"
)

func main() {

	_ = MyFile.NewMyFile("hoge.txt", func(myFile MyFile.ImyFile) error {
		buf, err := myFile.Read(10)
		println(*(*string)(unsafe.Pointer(&buf)))
		return err
	})

	/*
		this is error
		mf := &MyFile.myFile{}
		buf, _ := mf.Read(10)
		println(*(*string)(unsafe.Pointer(&buf)))
	*/
}
