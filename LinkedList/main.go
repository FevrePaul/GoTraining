package main

import (
    "fmt"
)

type LinkedList struct {
    val interface{}
    next * LinkedList
}

func ( l * LinkedList ) Insert( v interface{} ) {
    for l.next != nil {
        l = l.next
    }
    l.next = &LinkedList{ val: v, next: nil }
}

func ( l * LinkedList ) Erase( v interface{} ) *LinkedList {
    if l.val == v {
        newHead := l.next
        return newHead 
    }
    sent := l
    for sent.next != nil {
        if sent.next.val == v {
            sent.next = sent.next.next
            return l
        }
        sent = sent.next
    }
    return l
}

func ( l * LinkedList ) String() string {
    str := "LinkedList: "
    for l.next != nil {
        str += fmt.Sprint( l.val ) + " -> "
        l = l.next
    }
    str += fmt.Sprint( l.val )
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
    l = l.Erase( 5 )
    l = l.Erase( 14 )
    l = l.Erase( 12 )
    fmt.Println( l )
}
