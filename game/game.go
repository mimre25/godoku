package game

import (
	"errors"
	"fmt"
	"math/bits"
	"os"
	"sync"
)

type square struct {
	num int
	// byte i (from left) are 1 if the field can assume number i
	options int
}

type Block interface {
	[9]*square
}

func isValid[T Block](b T) bool {
	var x int = 0
	for _, i := range b {
		x |= 1 << i.num
	}
	return x == 0b1111111110
}

type Game struct {
	_data   [81]square
	rows    [9][9]*square
	columns [9][9]*square
	blocks  [9][9]*square
}

func NewGame() *Game {
	g := new(Game)
	for i := range 9 {
		block_i := i / 3
		inner_block_i := i % 3
		for j := range 9 {
			flat_idx := i*9 + j
			g._data[flat_idx].options = 0b1111111110
			addr := &g._data[flat_idx]
			g.rows[i][j] = addr
			g.columns[j][i] = addr

			block_j := j / 3
			inner_block_j := j % 3
			g.blocks[block_i*3+block_j][inner_block_i*3+inner_block_j] = addr

		}
	}

	return g

}

func (game *Game) PrintGame(printOptions ...bool) {
	for i := range 9 {
		fmt.Print(" ")
		for j := range 9 {
			if len(printOptions) > 0 {
				fmt.Printf("%010b ", game.rows[i][j].options)
			} else {
				fmt.Print(game.rows[i][j].num, " ")
			}
			if (j+1)%3 == 0 && j != 8 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if (i+1)%3 == 0 && i != 8 {
			fmt.Println("-----------------------")
		}
	}
}

func GameFromFile(filename string) *Game {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var numbers = make([]int, 0, 81)
	for _, d := range data {
		if d != '\n' {
			numbers = append(numbers, (int)(d-'0'))
		}
	}

	g := NewGame()
	for i, n := range numbers {
		g._data[i].num = n
	}
	return g

}

func (game *Game) IsFinished() bool {

	for i := range 9 {
		if !isValid(game.rows[i]) {
			return false
		}
		if !isValid(game.columns[i]) {
			return false
		}
		if !isValid(game.blocks[i]) {
			return false
		}
	}
	return true

}

func evaluateOptions[T Block](block T) {
	var options = 0
	for _, b := range block {
		if b.num != 0 {
			options |= 1 << b.num
			b.options = 0
		}
	}
	options = ^options
	for _, b := range block {
		b.options &= options
	}
}

func evaluateNumbers[T Block](block T) (bool, error) {
	var changed = false
	for _, b := range block {
		if b.options == 0 && b.num == 0 {
			return false, errors.New("Invalid Position")
		}
		if b.num != 0 {
			continue
		}
		switch b.options {
		case 1 << 1:
			b.num = 1
			changed = true
		case 1 << 2:
			b.num = 2
			changed = true
		case 1 << 3:
			b.num = 3
			changed = true
		case 1 << 4:
			b.num = 4
			changed = true
		case 1 << 5:
			b.num = 5
			changed = true
		case 1 << 6:
			b.num = 6
			changed = true
		case 1 << 7:
			b.num = 7
			changed = true
		case 1 << 8:
			b.num = 8
			changed = true
		case 1 << 9:
			b.num = 9
			changed = true
		}
		if changed {
			break
		}
	}
	return changed, nil
}

func (game *Game) evaluateOptions() {
	for i := range 9 {
		evaluateOptions(game.rows[i])
		evaluateOptions(game.columns[i])
		evaluateOptions(game.blocks[i])
	}
}

func (game *Game) evaluateNumbers() (bool, error) {
	var ret = false
	var changed bool
	var err error
	for i := range 9 {
		changed, err = evaluateNumbers(game.rows[i])
		if err != nil {
			return false, err
		}

		if changed {
			return true, nil
		}
		ret = ret || changed

		changed, err = evaluateNumbers(game.columns[i])
		if err != nil {
			return false, err
		}

		if changed {
			return true, nil
		}
		ret = ret || changed

		changed, err = evaluateNumbers(game.blocks[i])
		if err != nil {
			return false, err
		}

		if changed {
			return true, nil
		}
		ret = ret || changed
	}
	return ret, nil
}

func (game *Game) minOptions() int {
	var _min = 9
	var idx = -1
	for i, s := range game._data {
		if s.num == 0 {
			return i
		}
		if s.num != 0 {
			continue
		}
		ones := bits.OnesCount(uint(s.options))
		if ones < _min {
			_min = ones
			idx = i
		}
	}
	return idx

}

func (game *Game) deepCopy() *Game {
	newGame := NewGame()
	for i := range 81 {
		newGame._data[i].num = game._data[i].num
		newGame._data[i].options = game._data[i].options
	}
	return newGame
}

func (game *Game) Solve() error {
	for !game.IsFinished() {
		game.evaluateOptions()
		changed, err := game.evaluateNumbers()
		if err != nil {
			return err
		}
		if game.IsFinished() {
			break
		}
		if !changed {

			c1 := make(chan int, 81)
			var wg sync.WaitGroup
			idx := game.minOptions()

			if idx == -1 {
				return errors.New("No options left")
			}

			for _i := range 9 {
				i := _i + 1
				test := 1 << i
				if test&game._data[idx].options != 0 {
					wg.Add(1)

					go func(g *Game, c chan<- int) {
						g._data[idx].num = i
						err := g.Solve()
						if err == nil {
							c <- i
						} else {
							c <- -1
						}
						wg.Done()
					}(game.deepCopy(), c1)
				}
			}

			wg.Wait()
			close(c1)
			var val int
			for val = range c1 {
				if val != -1 {
					break
				}
			}
			if val == -1 {
				return errors.New("No solution found")
			}
			game._data[idx].num = val
		}
	}
	return nil

}
