package main

type cell struct {
	col int
	row int
}

type neighbors struct {
	topLeft     cell
	top         cell
	topRight    cell
	right       cell
	bottomRight cell
	bottom      cell
	bottomLeft  cell
	left        cell
}

func (c cell) val() int {
	return u.state[c.row][c.col]
}

func (c cell) numOfNeighbors() int {
	n := c.getNeighbors()

	return n.topLeft.val() + n.top.val() + n.topRight.val() + n.right.val() + n.bottomRight.val() + n.bottom.val() + n.bottomLeft.val() + n.left.val()
}

func (c cell) getNeighbors() neighbors {
	size := u.size

	n := neighbors{
		topLeft: cell{c.col-1, c.row-1},
		top: cell{c.col, c.row-1},
		topRight: cell{c.col+1, c.row-1},
		right: cell{c.col+1, c.row},
		bottomRight: cell{c.col+1, c.row+1},
		bottom: cell{c.col, c.row+1},
		bottomLeft: cell{c.col-1,c.row+1},
		left: cell{c.col-1, c.row},
	}

	// When a cell reaches the edge of the universe, its neighbors wrap around to the other side
	if c.col == 0 {
		n.topLeft.col = size - 1
		n.left.col = size - 1
		n.bottomLeft.col = size - 1
	}

	if c.col == size - 1 {
		n.topRight.col = 0
		n.right.col = 0
		n.bottomRight.col = 0
	}

	if c.row == 0 {
		n.topLeft.row = size - 1
		n.top.row = size - 1
		n.topRight.row = size - 1
	}

	if c.row == size - 1 {
		n.bottomLeft.row = 0
		n.bottom.row = 0
		n.bottomRight.row = 0
	}

	return n
}
