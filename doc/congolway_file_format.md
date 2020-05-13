# Congolway file format

This is the format that is used in this project.

## Sections
The first 7 rows are general information:

### Header

Header with the name of this project.
```
CONGOLWAY
```

### Version
Version of the file (there is only 1 version at the momento):
```
version: 1
```

### Generation
In case you want to keep count of your game of life generation
this field stores it.
```
generation: 543
```

### Neighborhood type
There are two neighborhood types:

#### Moore
8 surrounding cells to our cell:
```
neighborhood_type: Moore
```

#### Von Neumann
4 cells in star position  to our cell:
```
neighborhood_type: Von Neumann
```


### Size
The size of the cell grid in rowsxcolumns format.
Note this size must match the positions in the grid.
```
size: 5x5
```

### Limits
If you want a limitless grid or be limited by grid borders.

#### Limitless
If limitless, cells passing the grid will appear at the start of it.
```
limits: no
```

#### Row-Limited grid
A cell passing through the last column will appear in the first one.
```
limits: rows
```

#### Col-limited grid
A cell passing through the last row will appear in the first one.
```
limits: cols
```

#### Row & Col-limited grid
A cell passing through the last row or col will appear in the first one.
```
limits: rows, cols
```

### Type of grid (dense or sparse)

```
grid_type: dense|sparse
```

A dense grid shows all cells as a matrix and, hence,
is an easy way to see the cells but is not efficient
space-wise.

### Grid

Cell grid, depending on the grid type, it will be shown
as a matrix:

```
grid:
     
11111
11 11
11111
     
```
Each 1 or X represents an ALIVE cell,
and each 0 or space representes a DEAD cell.

Or as a sparse matrix:
```
grid:
default: 1
0: (0,0)(1,1)(2,2)
1:
```
