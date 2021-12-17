package main

import (
	"fmt"
	"math"
	"os"
)
import "bufio"

/**
 * Grab the pellets as fast as you can!
 **/

const (
	MINE     int = 1
	OPPONENT int = 0
)

type Position struct {
	x int // x: position in the grid
	y int // y: position in the grid
}

type PacID int
type PacMap map[PacID]Pac

type Pac struct {
	Position               // position in the grid
	pacId           PacID  // pacId: pac number (unique within a team)
	typeId          string // typeId: unused in wood leagues
	speedTurnsLeft  int    // speedTurnsLeft: unused in wood leagues
	abilityCooldown int    // abilityCooldown: unused in wood leagues
}

type Teams [2]Team

type Team struct {
	score  int
	pacMap PacMap
}

type Pellets struct {
	data          map[int]map[int]int
	maxValue      int
	maxPelletsPos map[Position]Position
}

func NewTeams() Teams {
	var teams Teams
	teams[MINE].pacMap = make(PacMap)
	teams[OPPONENT].pacMap = make(PacMap)
	return teams
}

func NewPellets(width, height int) Pellets {
	var pellets Pellets
	pellets.data = make(map[int]map[int]int)
	for i := 0; i <= width; i++ {
		pellets.data[i] = make(map[int]int)
	}
	return pellets
}

func (p *Pellets) insert(pos Position, value int) {
	p.data[pos.x][pos.y] = value
	if p.maxValue < value {
		p.maxValue = value
		p.maxPelletsPos = map[Position]Position{pos: pos}
	} else if p.maxValue == value {
		p.maxPelletsPos[pos] = pos
	}
	return
}

func (p Pellets) closestBigPellet(pacPosition Position) (closestPos Position) {
	minDistance := float64(math.MaxInt32)
	for _, pos := range p.maxPelletsPos {
		distance := math.Abs(float64(pacPosition.x-pos.x)) + math.Abs(float64(pacPosition.y-pos.y))
		if minDistance > distance {
			minDistance = distance
			closestPos = pos
		}
	}
	delete(p.maxPelletsPos, closestPos) // for next mine Pac
	return
}

func readLine(scanner *bufio.Scanner, args ...interface{}) {
	scanner.Scan()
	fmt.Sscan(scanner.Text(), args...)
}

func readPellets(scanner *bufio.Scanner, width, height, visiblePelletCount int) Pellets {
	pellets := NewPellets(width, height)
	for i := 0; i < visiblePelletCount; i++ {
		// value: amount of points this pellet is worth
		var value int
		var pos Position
		readLine(scanner, &pos.x, &pos.y, &value)
		pellets.insert(pos, value)
	}
	return pellets
}

func readTeams(scanner *bufio.Scanner) Teams {
	teams := NewTeams()
	readLine(scanner, &teams[MINE].score, &teams[OPPONENT].score)

	// visiblePacCount: all your pacs and enemy pacs in sight
	var visiblePacCount int
	readLine(scanner, &visiblePacCount)

	for i := 0; i < visiblePacCount; i++ {
		var pac Pac
		var mine int
		readLine(scanner, &pac.pacId, &mine, &pac.x, &pac.y, &pac.typeId, &pac.speedTurnsLeft, &pac.abilityCooldown)
		teams[mine].pacMap[pac.pacId] = pac
	}
	// visiblePelletCount: all pellets in sigh

	return teams
}

func orderToMove(teams Teams, pellets Pellets) {
	for id, pac := range teams[MINE].pacMap {
		pellet := pellets.closestBigPellet(pac.Position)
		fmt.Printf("MOVE %v %d %d\n", id, pellet.x, pellet.y) // MOVE <pacId> <x> <y>
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// width: size of the grid
	// height: top left corner is (x=0, y=0)
	var width, height int
	readLine(scanner, &width, &height)

	for i := 0; i < height; i++ {
		scanner.Scan()
		row := scanner.Text()
		_ = row // to avoid unused error // one line of the grid: space " " is floor, pound "#" is wall
	}

	for {
		teams := readTeams(scanner)
		// visiblePelletCount: all pellets in sigh
		var visiblePelletCount int
		readLine(scanner, &visiblePelletCount)

		pellets := readPellets(scanner, width, height, visiblePelletCount)

		// _ = PacMap
		// _ = pellets
		// fmt.Printf("MOVE 0 10 15\n") // MOVE <pacId> <x> <y>

		orderToMove(teams, pellets)
	}
}
