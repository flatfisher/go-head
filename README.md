# go-head
head command in Go lang.

## 使い方(go runで実行)
- test.txtの中身を表示(デフォルト10)

```
$ go run head.go test.txt

row1
row2
row3
row4
row5
row6
row7
row8
row9
row10
```

- 複数ファイルも対応

```
$ go run head.go test.txt test.txt

==> test.txt <==
row1
row2
row3
row4
row5
row6
row7
row8
row9
row10

==> test.txt <==
row1
row2
row3
row4
row5
row6
row7
row8
row9
row10
```

- test.txtの中身を15(n)行名まで表示

```
$ go run head.go -n 15 test.txt

row1
row2
row3
row4
row5
row6
row7
row8
row9
row10
row11
row12
row13
row14
row15
```

- 複数ファイルも対応

```
$ go run head.go -n 3 test.txt test.txt

==> test.txt <==
row1
row2
row3

==> test.txt <==
row1
row2
row3
```

- test.txtの行数を表示(中身は表示されない)

```
$ go run head.go -s test.txt

test.txt has 20 lines
```