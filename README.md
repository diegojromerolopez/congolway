# congolway
Conway's Game of Life gif generator in Go.

I wanted to make some kind of portmanteau between Conway and Game of Life.
Hence **congolway**. 

## Usage
```sh
Usage of ./main:
  -congolwayFilePath string
        Congay file
  -delay int
        Delay between frames, in 100ths of a second (default 5)
  -generations int
        Number of generations of the cellular automaton (default 100)
  -outputFilePath string
        File path where the output gif will be saved (default "out.gif")
```

## Samples

### Some ASCII art
```sh
./main -congolwayFilePath="./samples/goya.txt" -outputFilePath="./samples/goya.gif"
```

![Goya ASCII art](samples/goya.gif)

## TODO
* Search for better examples.
* Allow extracting size from grid.
* Parallelize.
* Read standard formats of game-of-lifes.
* Serve in a http server.

## License
[MIT](LICENSE)