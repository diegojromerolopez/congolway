# congolway
[Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) gif and apng generator in Go.

I wanted to make some kind of portmanteau between Conway and Game of Life.
Hence **congolway**.

## Features
* Parallel next generation implementation.
* Storing instances of Game of Life in text files.
* Generation of GIF and APNG animations for your game of life instances.
* Tested and developed following the advice of Go community.
* Support for Plaintext files, i.e. [.cells files](https://www.conwaylife.com/wiki/Plaintext).


## Construction
Use makefile to create executables in bin directory:

```sh
make build
```

## Usage

### APNG generator
Creates a [APNG](https://en.wikipedia.org/wiki/APNG) animation of the Game of Life.
```sh
Usage of ./bin/golapng:
  -generations int
        Number of generations of the cellular automaton (default 100)
  -inputFilePath string
        File path of the Congolway (.txt) or cells (.cells) file
  -outputFilePath string
        File path where the output apng will be saved (default "out.apng")
  -procs int
        Number of GO processes used to compute generations. By default is -1 (use as many as hardware CPUs), enter a positive integer to set a custom number of proceses (default -1)
```

### GIF generator
Creates a [GIF](https://en.wikipedia.org/wiki/GIF) animation of the Game of Life.
```sh
Usage of ./bin/golgif:
  -delay int
        Delay between frames, in 100ths of a second (default 5)
  -generations int
        Number of generations of the cellular automaton (default 100)
  -inputFilePath string
        File path of the Congolway (.txt) or cells (.cells) file
  -outputFilePath string
        File path where the output gif will be saved (default "out.gif")
  -procs int
        Number of GO processes used to compute generations. By default is -1 (use as many as hardware CPUs), enter a positive integer to set a custom number of proceses (default -1)
```

### SVG generator
Creates a [SVG](https://en.wikipedia.org/wiki/Scalable_Vector_Graphics) animation of the Game of Life.

**NOTE: this tool is in a highly-unstable state and generates extremly heavyweight SVGs.
I'm in process of optimizing it. Use it at your own risk.**

```sh
Usage of ./bin/golsvg:
  -delay int
        Delay between frames, in 100ths of a second (default 1)
  -generations int
        Number of generations of the cellular automaton (default 100)
  -inputFilePath string
        File path of the Congolway (.txt) or cells (.cells) file
  -outputFilePath string
        File path where the output gif will be saved (default "out.svg")
  -procs int
        Number of GO processes used to compute generations. By default is -1 (use as many as hardware CPUs), enter a positive integer to set a custom number of proceses (default -1)
```

### Random grid generator
Creates a txt file ([see its format](/doc/congolway_file_format.md)) with an (uniformly) random grid.
```sh
Usage of ./bin/randomgol:
  -columns int
        Number of columns of the grid (default 100)
  -description string
        Description of the game of life instance that will be created
  -name string
        Name of the game of life instance that will be created (default "Random Gol")
  -outputFilePath string
        File path where the random grid will be saved (only .txt and .cells extensions are allowed) (default "out.txt")
  -outputFormat string
        Only used for congolway files (.txt files). File format "dense" or "sparse"
  -randomSeed int
        Random seed
  -rows int
        Number of rows of the grid (default 100)
```

## Samples

Using the file [samples/grid100x100.txt](samples/grid100x100.txt):

### 100x100 gif animation
```sh
./bin/golgif -inputFilePath="./samples/grid100x100.txt" -outputFilePath="./samples/grid100x100.gif"
```

![grid100x100 gif](samples/grid100x100.gif)

### 100x100 apng animation
```sh
./bin/golgif -inputFilePath="./samples/grid100x100.txt" -outputFilePath="./samples/grid100x100.apng"
```

![grid100x100 apng](samples/grid100x100.apng)




## TODO
* ~~Different neighborhood types.~~
* ~~Infinite grids by horizontal or vertical directions.~~
* ~~Encode APNG.~~
* ~~Define a new format that is more compact (based on sparse matrix). Allow outputting in this format.~~
* Read the following formats:
  * ~~[Plaintext](https://www.conwaylife.com/wiki/Plaintext)~~
  * [Life 1.05](https://www.conwaylife.com/wiki/Life_1.05), i.e. .lif or .life files.
  * [Other formats](http://psoup.math.wisc.edu/mcell/ca_files_formats.html)
* Allow definition of multiple rules of spawning.
* ~~Allow cells with more states.~~ In case there is more states, allow definition of custom rules.
* Implement Grid as a sparse matrix.
* Continous integration.
* Read zipped files.
* Allow extracting size from grid.
* ~~Parallelize.~~
* Serve in a http server.

## License
[MIT](LICENSE)
