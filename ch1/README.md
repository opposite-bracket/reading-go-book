# Echo programs

I've run each program 5 times like so:

```shell
cd ch1
go build echo1
./echo1 hi there how is life with covid
```

and here's the result

| run | program | time |
| --- | ---- | --- |
| 1 | echo1 | 111.452µs |
| 2 | echo1 | 70.156µs |
| 3 | echo1 | 40.958µs |
| 4 | echo1 | 30.591µs |
| 5 | echo1 | 33.478µs |

| run | program | time |
| --- | ---- | --- |
| 1 | echo2 |  27.071µs |
| 2 | echo2 | 34.514µs |
| 3 | echo2 | 34.472µs |
| 4 | echo2 | 51.819µs |
| 5 | echo2 | 73.962µs |

| run | program | time |
| --- | ---- | --- |
| 1 | echo3 | 25.203µs |
| 2 | echo3 | 38.305µs |
| 3 | echo3 | 29.87µs |
| 4 | echo3 | 37.064µs |
| 5 | echo3 | 33.537µs |

| run | program | time |
| --- | ---- | --- |
| 1 | echo4 | 27.172µs |
| 2 | echo4 | 45.852µs |
| 3 | echo4 | 50.637µs |
| 4 | echo4 | 29.818µs |
| 5 | echo4 | 32.981µs |

| run | program | time |
| --- | ---- | --- |
| 1 | echo5 | 34.885µs |
| 2 | echo5 | 49.863µs |
| 3 | echo5 | 38.268µs |
| 4 | echo5 | 45.465µs |
| 5 | echo5 | 55.422µs |

# Dup programs

```shell
cd ch1
go run dup1.go

// "wq" to stop scanning standard input
go run dup2.go

go run dup2.go ./hello.txt ./hello2.txt

go run dup3.go ./hello.txt ./hello2.txt

go run dup4.go ./hello.txt ./hello2.txt
```
