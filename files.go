package main

import "fmt"
import "os"

func main () {


	f,err := os.Open("example.txt")
	if (err!=nil) {
		panic((err))
	}

	f1,err := f.Stat()
	if (err!=nil) {
		panic(err)
	}
	fmt.Println("filename", f1.Name())
	fmt.Printf("file or folder", f1.IsDir())
	fmt.Println("file size",f1.Size())

	//close the file 
	defer f.Close()

	buf := make([]byte,3)
	d, err := f.Read(buf)
	if err!=nil {
		panic(err)
	}

	for i :=0;i<len(buf);i++ {
		println("data",d,string(buf[i]))
	}

	//not recommended for large files 
	files,err := os.ReadFile("example.txt")
	if err!= nil {
		panic(err)
	}
	fmt.Println(string(files))

	//code to see folders
	dir,err := os.Open("..")
	if err!= nil {
		panic(err)
	}
	info,err := dir.ReadDir(-1)
	if err!= nil {
		panic(err)
	}
	for _,fi := range info {
		fmt.Println(fi.Name(),fi.IsDir())
	}

	//create files and write
	write,err := os.Create("example2.txt ")
	if err!= nil {
		panic(err)
	}
	defer write.Close()
	write.WriteString("bye")

}
