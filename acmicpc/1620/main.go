package main

import (
	"bufio"
	"os"
	"strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
	sc.Scan()
	for _, c := range sc.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wt.Flush()
	pokemons, problems := nextInt(), nextInt()
	pokeMonBook := make(map[string]string)

	ReadPokeMons(&pokeMonBook, pokemons)
	SolvePokeMons(&pokeMonBook, problems)
}

func ReadPokeMons(pokeMonBook *map[string]string, pokemons int) {
	for idx := 1; idx <= pokemons; idx++ {
		sc.Scan()
		pokemon := sc.Text()
		i := strconv.Itoa(idx)
		(*pokeMonBook)[pokemon] = i
		(*pokeMonBook)[i] = pokemon
	}
}

func SolvePokeMons(pokeMonBook *map[string]string, problems int) {
	for ; problems > 0; problems-- {
		sc.Scan()
		wt.WriteString((*pokeMonBook)[sc.Text()])
		wt.WriteByte('\n')
	}
}
