package main

import "github.com/01-edu/z01"

func main() {
	QuadA(1,5)
}

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		space := -2
		printWidth(x, &space, true)
		if y >= 2 {
			if x > 1 {
				for z := 1; z <= y-2; z++ {
					z01.PrintRune('*')
					for j := 1; j <= space; j++ {
						z01.PrintRune(' ')
					}
					z01.PrintRune('*')
					z01.PrintRune('\n')
				}
				printWidth(x, &space, false)
			} else {
				for z := 1; z <= y-2; z++ {
					z01.PrintRune('*')
					z01.PrintRune('\n')
				}
				printWidth(x, &space, false)
			}
		}
	}
}

func printWidth(x int, space *int, firstLine bool) {
	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			if firstLine == true{
				if i == 1 {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('\\')
				}
			} else {
				if i == 1 {
					z01.PrintRune('\\')
				} else {
					z01.PrintRune('/')
				}
			}
		} else {
			z01.PrintRune('*')
		}
		*space++
	}
	z01.PrintRune('\n')
}
