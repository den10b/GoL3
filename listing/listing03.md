Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
В пустых интерфейсах нет ссылки на метаданные(в т.ч список методов)
Вывод:
nil
false
тк в err лежит ссылка на nil-ошибку, а не просто nil

```
