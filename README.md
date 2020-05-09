# congolway
Conway's Game of Life gif generator in Go.

I wanted to make some kind of portmanteau between Conway and Game of Life.
Hence **congolway**.

## Construction
Use makefile to create executables in bin directory:

```sh
make build
```

## Usage

### GIF generator
Creates a GIF animation of the Game of Life.
```sh
Usage of ./bin/golgif:
  -inputFilePath string
        Congay file
  -delay int
        Delay between frames, in 100ths of a second (default 5)
  -generations int
        Number of generations of the cellular automaton (default 100)
  -outputFilePath string
        File path where the output gif will be saved (default "out.gif")
```

### Random grid generator
```sh
Usage of ./bin/randomgol:
  -columns int
        Number of columns of the grid (default 100)
  -outputFilePath string
        File path where the random grid will be saved (default "out.txt")
  -randomSeed int
        Rnadom 
  -rows int
        Number of rows of the grid (default 100)
```


## Samples

```sh
./bin/golgif -inputFilePath="./samples/grid100x100.txt" -outputFilePath="./samples/grid100x100.gif"
```

![grid100x100 gif](samples/grid100x100.gif)

## TODO
* Allow infinite grids by horizontal or vertical directions.
* Allow extracting size from grid.
* Parallelize.
* Read standard formats of game-of-lifes.
* Serve in a http server.

## License
[MIT](LICENSE)