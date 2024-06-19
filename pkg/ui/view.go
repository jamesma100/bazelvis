package ui

import (
	"github.com/jroimartin/gocui"
	"log"
)

type view_info struct {
	target   string
	parent   *view_info
	depMap   map[string][]string
	maxLines int
	lines    []string
	cursor   int
	esp      int
	ptr      int
	gui      *gocui.Gui
}

// generate view
func (vi *view_info) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	left := maxX / 8
	if v, err := g.SetView("deps", left, 3, maxX-left, maxY-8); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "welcome to bazelvis!"

		_, vi.maxLines = v.Size()
		if vi.lines == nil {
			vi.lines = make([]string, vi.maxLines)
		}
		vi.writeDown(true, 0)

		if _, err := g.SetCurrentView("deps"); err != nil {
			return err
		}
	}
	return nil
}

// exit via ctrl-c
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func initKeybindings(g *gocui.Gui, vi *view_info) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	keyMap := map[interface{}]func(g *gocui.Gui, v *gocui.View) error{
		'k':                vi.cursorUp,
		gocui.KeyArrowUp:   vi.cursorUp,
		'j':                vi.cursorDown,
		gocui.KeyArrowDown: vi.cursorDown,
		gocui.KeyCtrlB:     vi.pageUp,
		gocui.KeyCtrlF:     vi.pageDown,
		gocui.KeyEnter:     vi.traverse,
	}

	for key, f := range keyMap {
		if err := g.SetKeybinding("deps", key, gocui.ModNone, f); err != nil {
			return err
		}
	}

	return nil
}

func StartUI(target string, depMap map[string][]string) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	vi := view_info{
		target:   target,
		parent:   nil,
		depMap:   depMap,
		maxLines: 0,
		lines:    nil,
		cursor:   0,
		esp:      0,
		ptr:      0,
		gui:      g,
	}

	g.InputEsc = true
	g.ASCII = true

	g.SetManagerFunc(vi.layout)

	if err := initKeybindings(g, &vi); err != nil {
		log.Fatal(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
