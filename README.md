# Ethereum address generator

cli tool to generate ethereum addresses


## How it works

Run the program in the terminal and/or add generator options.

## options



```
  -n int
    	Number of addresses generated (default 1)
  -p string
    	Address prefix (default none)
  -q string
    	Address substring (default none)
  -s string
    	Address suffix (default none)
```
## example

generate 2 addresses starting with '0xee' and ending with 'c'
```
go run . -p ee -s c -n 2
```
## Authors

- [@reshkovv](https://github.com/reshkovv)

