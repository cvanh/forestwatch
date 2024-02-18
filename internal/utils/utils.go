package utils

import "log"

// implements php's array_diff()
// checks what array1 misses
// see https://www.php.net/manual/en/function.array-diff.php
// https://github.com/php/php-src/blob/ec9b68cb6a318f3a55a4d9377076abea340fe814/ext/standard/array.c#L5460
// https://github.com/php/php-src/blob/ec9b68cb6a318f3a55a4d9377076abea340fe814/ext/standard/tests/array/array_diff_1.phpt#L4
// https://www.php2golang.com/method/function.array-diff.html
func ArrayDiff(array1, array2 interface{}) []any {
	arr := []any{}

	for i := 0; i < len(array1); i++ {
		log.Println("asd")

	}
	return arr

}
