package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6678")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopyss(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopyss(w io.Writer,r io.Reader)  {
	_,err:=io.Copy(w,r)
	if err!= nil{
		log.Fatal(err)
	}
}

func aaa()  {

}