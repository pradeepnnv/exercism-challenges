package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name          string
	matchesPlayed int
	won           int
	draw          int
	lost          int
}

func (t team) points() int {
	return t.won*3 + t.draw*1
}
func (t team) String() string {

	return fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", t.name, t.matchesPlayed, t.won, t.draw, t.lost, t.points())
}

var aa, bb, cc, dd team

func Tally(reader io.Reader, writer io.Writer) error {

	input := bufio.NewScanner(reader)
	input.Text() // for skipping the very initial new line
	isValid := true

	aa = team{
		name: "Allegoric Alaskians",
	}

	bb = team{
		name: "Blithering Badgers",
	}

	cc = team{
		name: "Courageous Californians",
	}

	dd = team{
		name: "Devastating Donkeys",
	}

	for input.Scan() {

		s := strings.ToLower(input.Text())
		if s == "" || s[0] == '#' {
			continue
		}

		if !validate(s) {
			isValid = false
			continue
		}

		split := strings.Split(s, ";")
		team1, team2, result := split[0], split[1], split[2]

		switch result {
		case "win":
			win := identifyTeam(team1)
			loss := identifyTeam(team2)
			win.won++
			win.matchesPlayed++
			loss.lost++
			loss.matchesPlayed++
		case "loss":
			win := identifyTeam(team2)
			loss := identifyTeam(team1)
			win.won++
			win.matchesPlayed++
			loss.lost++
			loss.matchesPlayed++
		case "draw":
			t1 := identifyTeam(team2)
			t2 := identifyTeam(team1)
			t1.matchesPlayed++
			t1.draw++
			t2.matchesPlayed++
			t2.draw++
		}

	}
	teams := []team{aa, bb, cc, dd}
	sort.Sort(ByPoints(teams))

	if !isValid {
		return fmt.Errorf("invalid input")
	}
	var b strings.Builder
	b.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	for _, t := range teams {
		b.WriteString(t.String())
	}
	writer.Write([]byte(b.String()))

	return nil
}

func identifyTeam(s string) *team {
	switch s {
	case "devastating donkeys":
		return &dd
	case "allegoric alaskians":
		return &aa
	case "blithering badgers":
		return &bb
	case "courageous californians":
		return &cc
	}
	return &team{}
}

func validate(s string) bool {
	split := strings.Split(s, ";")
	if len(split) < 3 || !checkName(split[0]) || !checkName(split[1]) || !checkResult(split[2]) {
		return false
	}
	return true
}

func checkName(s string) bool {
	switch s {
	case "devastating donkeys":
	case "allegoric alaskians":
	case "blithering badgers":
	case "courageous californians":
	default:
		return false
	}
	return true
}

func checkResult(s string) bool {
	switch s {
	case "win":
	case "draw":
	case "loss":
	default:
		return false
	}
	return true
}

//Sorting

type ByPoints []team

func (a ByPoints) Len() int           { return len(a) }
func (a ByPoints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPoints) Less(i, j int) bool { return a[i].points() > a[j].points() }
