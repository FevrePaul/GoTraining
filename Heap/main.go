package main

import (
    "fmt"
)

type MaxHeap struct {
    arr []int 
}

func ( h * MaxHeap ) Insert( key int ) {
    h.arr = append( h.arr, key )
    h.heapifyUp( len( h.arr ) - 1 )
}

func ( h * MaxHeap ) ExtractMax() int {
    extracted := h.arr[ 0 ]
    l := len( h.arr ) - 1
    if l < 0 {
        fmt.Println("Cannot extract, heap is empty.")
        return -1
    }
    h.arr[ 0 ] = h.arr[ l ]
    h.arr = h.arr[:l]
    h.heapifyDown()
    return extracted
}

func ( h * MaxHeap ) heapifyDown() {
    idx := 0
    l, r := left( idx ), right( idx )
    childIdx := 0
    for l < len( h.arr ) - 1 {
        if l == len( h.arr ) - 1 {
            childIdx = l
        } else if h.arr[ r ] < h.arr[ l ] { 
            childIdx = l
        } else {
            childIdx = r
        }

        if h.arr[ idx ] < h.arr[ childIdx ] {
            h.arr[ idx ], h.arr[ childIdx ] = h.arr[ childIdx ], h.arr[ idx ]
        }

        idx = childIdx
        l, r = left( idx ), right( idx )
    }
}

func ( h * MaxHeap ) heapifyUp( idx int ) {
    for p := getParentIdx( idx ); h.arr[ p ] < h.arr[ idx ]; p = getParentIdx( idx ) {
        h.arr[ p ], h.arr[ idx ] = h.arr[ idx ], h.arr[ p ]
        idx = p
    }
}

func getParentIdx( idx int ) int {
    return ( idx - 1 ) / 2
}

func left( idx int ) int {
    return ( idx * 2 ) + 1
}

func right( idx int ) int {
    return ( idx * 2 ) + 2
}


func ( h * MaxHeap ) check( i int ) bool {
    if i >= ( len( h.arr ) - 1 ) / 2 {
        return true
    }
    if h.arr[ i ] >= h.arr[ 2 * i + 1 ] &&
         h.arr[ i ] >= h.arr[ 2 * i + 2 ] &&
         h.check( 2 * i + 1 ) && h.check( 2 * i + 2 ) {
        return true
    }

    return false
}


func main() {
    h := &MaxHeap{}
    h.Insert( 7 )
    h.Insert( 1 )
    h.Insert( 2 )
    h.Insert( 9 )
    h.Insert( 18 )
    h.Insert( 12 )
    h.ExtractMax()
    fmt.Println(h.check( 0 ))
    h.ExtractMax()
    fmt.Println(h.check( 0 ))
    h.ExtractMax()
    fmt.Println(h.check( 0 ))
    h.ExtractMax()
    fmt.Println(h.check( 0 ))
    fmt.Println( h )
}
