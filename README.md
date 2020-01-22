# dirTreeBuild
The program prints out the file (if -f options is passed) and directory tree.


```go
go run main.go . -f // with files
go run main.go . // without files

where . is a path relative to current program directory
```

## Example

```go
├───main.go (1881b)
├───main_test.go (1318b)
└───testdata
	├───project
	│	├───file.txt (19b)
	│	└───gopher.png (70372b)
	├───static
	│	├───css
	│	│	└───body.css (28b)
	│	├───html
	│	│	└───index.html (57b)
	│	└───js
	│		└───site.js (10b)
	├───zline
	│	└───empty.txt (empty)
	└───zzfile.txt (empty)
```
