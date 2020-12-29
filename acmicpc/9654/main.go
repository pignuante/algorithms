package main

import (
    "bufio"
    "os"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    wt.WriteString("SHIP NAME      CLASS          DEPLOYMENT IN SERVICE")
    wt.WriteByte('\n')
    wt.WriteString("N2 Bomber      Heavy Fighter  Limited    21        ")
    wt.WriteByte('\n')
    wt.WriteString("J-Type 327     Light Combat   Unlimited  1         ")
    wt.WriteByte('\n')
    wt.WriteString("NX Cruiser     Medium Fighter Limited    18        ")
    wt.WriteByte('\n')
    wt.WriteString("N1 Starfighter Medium Fighter Unlimited  25        ")
    wt.WriteByte('\n')
    wt.WriteString("Royal Cruiser  Light Combat   Limited    4         ")
}
