package main

import (
    "fmt"
)

type LinkedList struct {
    val interface{}
    next * LinkedList
}

func ( l * LinkedList ) Insert( v interface{} ) {
    sent := l
    for sent.next != nil {
        sent = sent.next
    }
    sent.next = &LinkedList{ val: v, next: nil }
}

func ( l * LinkedList ) Erase( v interface{} ) bool {
    sent := l
    for sent.next != nil {
        sent = sent.next
    }
}

func ( l * LinkedList ) String() string {
    str := "LinkedList: "
    sent := l
    for sent.next != nil {
        str += fmt.Sprint( sent.val ) + " -> "
        sent = sent.next
    }
    str += fmt.Sprint( sent.val )
    return str
}

func main() {
    l := &LinkedList{ val: 5 }
    l.Insert( 81 )
    l.Insert( 1 )
    l.Insert( 1002 )
    l.Insert( 12 )
    l.Insert( 8 )
    l.Insert( 15 )
    l.Insert( 6 )
    l.Insert( 56 )
    fmt.Println( l )
}
