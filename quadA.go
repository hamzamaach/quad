package piscine

import "github.com/01-edu/z01"

func QuadA(x int, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= x; i++ {
			if i == 1 || i == x {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')
		for j := 1; j <= y-2; j++ {
			z01.PrintRune('|')
			if x > 1 {

				for k := 1; k <= x-2; k++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune('|')

			}
			z01.PrintRune('\n')
		}
		if y > 1 {
			for i := 1; i <= x; i++ {
				if i == 1 || i == x {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
