package ggol

var initTestCell TestCell = TestCell{
	Alive: false,
}

var testCellIterator CellIterator = func(cell interface{}, adjacentCells []interface{}) interface{} {
	newCell := cell.(TestCell)

	var aliveNbrsCount int = 0
	for i := 0; i < len(adjacentCells); i += 1 {
		adjacentCells := adjacentCells[i].(TestCell)
		if adjacentCells.Alive {
			aliveNbrsCount += 1
		}
	}
	if newCell.Alive {
		if aliveNbrsCount != 2 && aliveNbrsCount != 3 {
			newCell.Alive = false
			return newCell
		} else {
			newCell.Alive = true
			return newCell
		}
	} else {
		if aliveNbrsCount == 3 {
			newCell.Alive = true
			return newCell
		} else {
			newCell.Alive = false
			return newCell
		}
	}
}

// Check if two AliveCellsMaps are equal.
func areAliveCellsMapsEqual(a aliveCellsMap, b aliveCellsMap) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func convertGenerationToAliveCellsMap(g *Generation) *aliveCellsMap {
	gMap := make(aliveCellsMap, 0)
	for x := 0; x < len(*g); x++ {
		gMap = append(gMap, []bool{})
		for y := 0; y < len((*g)[x]); y++ {
			gMap[x] = append(gMap[x], ((*g)[x][y]).(TestCell).Alive)
		}
	}

	return &gMap
}
