package main

func sum(a []int) (sum int) {

    sum = 0
    for _, v := range a {
        sum += v
    }

    return
}
