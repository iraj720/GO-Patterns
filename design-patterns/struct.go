package designpatterns

// when passing struct as an interface be careful that
// its your option to just use methods that have value receiver and pass it by value
// or you have pointer receiver so you should use its refrence
// so its some how different from other languages that methods are always dedicated to a single object
// here objects are dedicated to methods and it can be by ref or value some how same as python
// that you can pass self as a param or not
// but others are not dynamic in this case
//
// so what is the trade-off ?? is golang and pyhton are better in this case ??
// actually not. because in python and golang you are passing refrence to the objects
// or passing value of object to methods and it will increase cost of memory
// but in others like c# cpp objects has refrence of methods and call them
// so in golang you pay memory and get dynamic design but in others none of them
// so memory/design (dont forget that most of times its a cheap cost that its worth to pay ;) )
