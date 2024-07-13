package main

func main() {

}

type SubrectangleQueries struct {
	queries [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	sq := SubrectangleQueries{queries: make([][]int, len(rectangle))}
	for i := range sq.queries {
		sq.queries[i] = make([]int, len(rectangle[0]))
		copy(sq.queries[i], rectangle[i])
	}
	return sq
}

func (sq *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	buf := make([]int, col2-col1+1)
	for i := range buf {
		buf[i] = newValue
	}
	for i := row1; i <= row2; i++ {
		copy(sq.queries[i][col1:col2+1], buf)
	}
}

func (sq *SubrectangleQueries) GetValue(row int, col int) int {
	return sq.queries[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
