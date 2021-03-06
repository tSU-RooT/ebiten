// Copyright 2016 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package twenty48

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 210
	ScreenHeight = 300
	boardSize    = 4
)

type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
}

func NewGame() (*Game, error) {
	g := &Game{
		input: NewInput(),
	}
	var err error
	g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Update() error {
	if err := g.input.Update(); err != nil {
		return err
	}
	if err := g.board.Update(g.input); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) error {
	if g.boardImage == nil {
		var err error
		w, h := g.board.Size()
		g.boardImage, err = ebiten.NewImage(w, h, ebiten.FilterNearest)
		if err != nil {
			return err
		}
	}
	if err := screen.Fill(backgroundColor); err != nil {
		return err
	}
	if err := g.board.Draw(g.boardImage); err != nil {
		return err
	}
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	if err := screen.DrawImage(g.boardImage, op); err != nil {
		return err
	}
	return nil
}
