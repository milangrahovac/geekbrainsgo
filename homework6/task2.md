var a int = 5
var b = &a
var c int = 2
_____________________

fmt.Println(b) - print 0xc00001a068
fmt.Println(*b) - print 5 - разыменование b

При умножении у нас всегда два операнда, а при разыменовании указателя - один.

Если мы выполняем какие-то математические операции со значением указателей, * всегда должно стоять перед указателем.
a + *b  - return 10 -  разыменование b
a * *b  - return 25 - разыменование b и умножение var a и значения указателя b
*b - a  - return 0 - разыменование b
*b / a  - return 1 -  разыменование b

Между * и переменной не может быть пустого места. пример "*b".

fmt.Println(a**b) - 25 - разыменование b и умножение var a и значения указателя b
fmt.Println(a, *b) - return 5, 5 - разыменование b

Две переменные или переменная и понтер не могут стоять рядом друг с другом. Что-то должно оставаться между ними. Например, «+», «-», «*», «/», «,» и т. Д.
fmt.Println(ab) - undefined: ab
fmt.Println(a*b) - invalid operation: a * b

a * c - return 10 - умножение 2 var


