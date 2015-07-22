package main
	
import (
"fmt"
"time"
)

func main(){

// If Else Example

	print := fmt.Println

	tick := time.Now()

	tic := time.Now()

	he := tic.Sub(tick)

	print(he.Minutes())

	print (tick.Format(time.RFC3339))	

	print("Hello Dude")

	b := "Vamsee"

	if b == "Vamsi" {

		fmt.Println("Yes You are Vamsi")
		} else {
		 fmt.Println("Sorry you are not vamsi")
		}

	if c := "Vamsi";c == "Vamsi"{
		fmt.Println("Yes you are man you are!!!!!")
	}	else {
		fmt.Println("Sorry you are not period!!!")
	}

// For Loop Example
	k := "Hello Gopher!!!!"

	i := 10

	for m:=0;m<i;m++ {
		fmt.Println(k);
	}


fmt.Println(time.Now().Unix())


// Switch Example

switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
	fmt.Println("Its a Weekend")
default:
	fmt.Println("Its a Weekday")
}


//Arrays Example

var a[5]int

a[4] = 100
fmt.Println("set :", a)

fmt.Println("get :", a[4])

a[3] = 1000

print(a)

print(a[3])

if a[2] == 400 {
	print ("Yes it is")
} else {
	print ("Sorry its not")
}


//Slices Example

s := make([]string, 3)

print("emp:", s)

s[0] ="Hello"
s[1] ="Hey"
s[2] ="Vola"

print(s)

// Append inside the Slices
s = append(s, "New")


//Creating Two dimensional array
twoD := make([][]int, 3)

for x :=0;x <3;x++ {
	innerLen := x+1
	twoD[x] = make([]int, innerLen)
	for y :=0;y<innerLen;y++ {
		twoD[x][y] = x+y
	}
}

print(twoD)

print(s)

// Another Array example

f := [2]string{"Hey","Vamsi"}

q :=[...]string{"Hey","Dude","How","Are","You"}
print(q)

print(f)


// Map examples
//make(map[key-type]val-type).

o := make(map[int]string)

o[0] = "Hello"
o[1] = "Vamsi"

print(o)

delete(o,1)
print (o)


u := map[string]string{"Name": "Vamsi","School": "Aditya"}

print (u)


//


}