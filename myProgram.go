package main
import (
"ethos/syscall"
"ethos/ethos"
ethosLog "ethos/log"
"ethos/efmt" 
"fmt"
"math"
)
func main () {
me := syscall.GetUser()
path := "/user/" + me + "/myDir/"
status := ethosLog.RedirectToLog("myProgram")
if status != syscall.StatusOk {
efmt.Fprintf(syscall.Stderr, "Error opening %v: %v\n", path, 
status)
syscall.Exit(syscall.StatusOk)
}

fmt.Println(" ")
fmt.Println("============================================")
fmt.Println(" ")
fmt.Println("[Initiate Box]------------------------------")
data := Box{}
fmt.Println(data)
fmt.Println("[change points valuse]----------------------")
data.pointA.x=1
data.pointA.y=1
data.pointB.x=10
data.pointB.y=5
fmt.Println(data)
fmt.Println("[shift the box]-----------------------------")
data.pointA.x = 5*data.pointA.x
data.pointA.y = 5*data.pointA.y
data.pointB.x = 5*data.pointB.x
data.pointB.y = 5*data.pointB.y
fmt.Println(data)
fmt.Println("[Hight & width]-----------------------------")
fmt.Println("The hight of the box is = ", data.pointB.y - data.pointA.y)
fmt.Println("The width of the box is = ", data.pointB.x - data.pointA.x)
fmt.Println("The square root of point A is = ", math.Sqrt(data.pointA.x), math.Sqrt(data.pointA.y))
fmt.Println("The square root of point B is = ", math.Sqrt(data.pointB.x), math.Sqrt(data.pointB.y))
fmt.Println("--------------------------------------------")

fd, status := ethos.OpenDirectoryPath(path)
data.Write(fd)
data.WriteVar(path +"foobar")
efmt.Fprint(syscall.Stderr, "/n *****this will print in the log***** /n")
}
