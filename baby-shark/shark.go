package baby_shark

import "strings"

var r = strings.Repeat

func Shark() (l string) {
	for _, m := range []string{"Baby shark", "Mommy shark", "Daddy shark", "Grandma shark", "Grandpa shark", "Let's go hunt"} {
		l += r(m+","+r(" doo", 6)+"\n", 3) + m + "!\n"
	}
	l += "Run away,…\n"
	return
}
