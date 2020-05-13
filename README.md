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
  -procs int
        Number of GO processes used to compute generations. By default is -1 (use as many as hardware CPUs),
        enter a positive integer to set a custom number of proceses (default -1)
```

### Random grid generator
Creates a txt file ([see its format](/doc/congolway_file_format.md)) with an (uniformly) random grid.
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
  -outputFormat string
        File format "dense" or "sparse" (default "dense")
```


## Samples

```sh
./bin/golgif -inputFilePath="./samples/grid100x100.txt" -outputFilePath="./samples/grid100x100.gif"
```

![grid100x100 gif](samples/grid100x100.gif)

See the file [samples/grid100x100.txt](samples/grid100x100.txt) and see for yourself the Congolway file format.

## TODO
* ~~Different neighborhood types.~~
* ~~Infinite grids by horizontal or vertical directions.~~
* Encode APNG.
* ~~Define a new format that is more compact (based on sparse matrix). Allow outputting in this format.~~
* Allow definition of multiple rules of spawning.
* ~~Allow cells with more states.~~ In case there is more states, allow definition of custom rules.
* Implement Grid as a sparse matrix.
* Continous integration.
* Read zipped files.
* Allow extracting size from grid.
* ~~Parallelize.~~
* Read standard formats of game-of-lifes.
* Serve in a http server.

## License
[MIT](LICENSE)
