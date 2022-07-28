package concurency

// channels always pass values unless you want to copy the address
// if your data size is same as address of that data
// so passing by pointer will not reduce cost of garbage collector
// order of goroutines working on channel is fifo
// its reason is goroutines and go schedular
// because all of goroutines are being hold by P part of go schedular
// and P is a Q of running goroutines so its fifo and make channels fifo
// channel is a struct made of three main fields circular buffer, read index, write index, lock
// whenever you read/write it will changes the index of reading/writing with a lock
// and you pay cost of copying two times(read and write) whenever sending data
// that is nessecery for memory safety
// initializing channel allocates some memory on heap and returns a pointer to it so pointer to channel is useless

// goroutines g
// goroutines are user-space threads light weight and managed by go runtime(go schedular)
// every g has a stack in memory and gs will never write/read from other gs stack but it has 1 exception
// and it is when writing a writer to an empty channel that has a paused reader
// file:./threads_goroutins.png

// go schedular
// go schedular M:N model manage the threads and goroutines
// it has 3 main parts
// M : osThread
// G : goroutine
// P : it holds queue for goroutines that are runnable
// so how it works lets start with simple example think about a situation that channel is full
// and we want to write element T1 to it. it will pause and send gopark command to go scher but how ?
// first of all go schedular will change the state of goroutine to waiting and detatching it from M
// poping another g from P and attaching to M(running os thread)
// the main advantage of this model is here it wont block the os thread it will run another attached G
// what is happening to channel
// channel has 2 Qs sendQ and recQ that both of them are same thing they have a pointer to reader/writer G
// and T1 (a pointer to the element that we will read from or write into it) 
// G will copy its element to T1 when writing to channel (in reading its sh different)
// so paused Gs will be here in theses 2 Qs and whenever reader comes in
// reader will remove writer G from writeQ and copy its element on circular buffer
// and so when writer resumes it will not mess with the channel and its an optimization
// then reader call goready command to go schedular and it will make it runnable
// and schedualr puts it on P
// when reader reading from an empty channel it will create sudog and put it on recQ of chan and calls gopark
// now its paused when writer comes across it wont write it to Circular buffer because now
// we have pointer of element of reading channel in recQ so we will write it directly to that element
// and thats NICE !


// const
// all consts should have an exact value at compile time
// e.g : const a = getA() has compile error cause getA() will be known at runtime
// there is no implicit conversion in golang except using untyped const
// consts default :
// 123        //Default hidden type is int
// "circle"   //Default hidden type is string
// 5.6.       //Default hidden type is float64
// true       //Default hidden type is bool
// 'a'        //Default hidden type is int32 or rune
// 3+5i       //Default hidden type is complex128



// defer
// when you defer a function go puts it into a stack 
// go just put any defer he sees into that stack and after when closing each stack function 
// it will pop 1 defer and call it and so on ...
// so below code will result in : 3, 2, 1
//     i := 0
//     i = 1
//     defer fmt.Println(i)
//     i = 2
//     defer fmt.Println(i)
//     i = 3
//     defer fmt.Println(i)

// defer statements are evaluated in the line we call it
// and it doesnt care about any changes between defer line and closing of stack of the function


// pointer arithmetic is not possible in golang and you cannot add to refrence and 
// see the next byte of memory which is able in C



// structs 
// you can add meta data to your struct fields
// e.g when you add 'json: "n"' to a field its json encoding will be different
// and that key should be n not the name of the field  
// 2 structs are equal if and only if all of their fields have same value
// if the struct has slice, map or function its not comparable and will result in compile error