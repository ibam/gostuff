# grouplines
Take a given input file, group the input every given lines using the given separator, write it back to the given output file.

Options:
```
-i Path to input file (mandatory)
-o Path to output file (mandatory)
-s Separator character (optional, defaults to ',')
-l Number of lines to be grouped (optional, defaults to 20)
```

Example usage:

```
go run grouplines.go -i input.txt -o output.txt -s \| -l 10
```