Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
    //close(ch)
	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1                                                         
...                                                        
9                                                         
fatal error: all goroutines are asleep - deadlock! 
тк из цикла for n := range ch произойдет выход если канал закроется, иначе программа будет ждать новых данных в канале
если раскомментить строку с закрытием канала после цикла, то дедлока не будет
```
