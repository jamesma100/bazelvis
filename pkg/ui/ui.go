package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func (vi *view_info) cursorUp(g *gocui.Gui, v *gocui.View) error {
	if vi.cursor > 0 {
		vi.cursor--
		start_idx := vi.ptr - vi.maxLines + 1
		if start_idx < 0 {
			start_idx = 0
		}
		vi.writeDown(true, start_idx)
	} else {
		if vi.ptr-vi.maxLines+1 > 0 {
			vi.writeDown(true, vi.ptr-vi.maxLines)
		}
	}
	return nil
}

func (vi *view_info) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if vi.cursor < vi.maxLines-1 {
		if vi.cursor < vi.esp {
			vi.cursor++
			vi.writeDown(true, 0)
		}
	} else {
		if vi.ptr < len(vi.depMap[vi.target])-1 {
			vi.writeUp(true, vi.ptr+1)
		}
	}
	return nil
}

func (vi *view_info) pageUp(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func (vi *view_info) pageDown(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func (vi *view_info) traverse(g *gocui.Gui, v *gocui.View) error {
	newTarget := vi.lines[vi.cursor]
	if newTarget == "../" {
		newTarget = vi.parent.target
	}
	if newTarget == vi.target {
		return nil
	}
	if vi.parent != nil && newTarget == vi.parent.target {
		vi.backTrack()
		return nil
	}
	vi.forwardTrack()
	return nil
}

func (vi *view_info) forwardTrack() {
	p := view_info{
		target:   vi.target,
		parent:   vi.parent,
		depMap:   vi.depMap,
		maxLines: vi.maxLines,
		lines:    vi.lines,
		cursor:   vi.cursor,
		esp:      vi.esp,
		ptr:      vi.ptr,
		gui:      vi.gui,
	}
	vi.target = vi.lines[vi.cursor]
	vi.cursor = 0
	vi.parent = &p
	vi.writeDown(true, 0)
}

func (vi *view_info) backTrack() {
	vi.target = vi.parent.target
	vi.cursor = vi.parent.cursor
	vi.parent = vi.parent.parent
	vi.writeDown(true, 0)
}

func (vi *view_info) writeDown(refreshView bool, start int) {
	vi.lines = make([]string, vi.maxLines)
	i := 0
	if vi.parent != nil && vi.parent.target != "" {
		if start == 0 {
			i++
			vi.lines[0] = "../"
		}
	}
	j := start
	deps := vi.depMap[vi.target]
	for ; i < vi.maxLines; i++ {
		if j > len(deps)-1 {
			break
		}
		vi.lines[i] = deps[j]
		j++
	}
	vi.esp = i - 1
	vi.ptr = j - 1

	if refreshView {
		vi.refreshView()
	}
}

func (vi *view_info) writeUp(refreshView bool, end int) {
	start := 0
	vi.lines = make([]string, vi.maxLines)
	deps := vi.depMap[vi.target]
	j := end
	for i := vi.maxLines - 1; i >= start; i-- {
		vi.lines[i] = deps[j]
		j--
	}
	vi.ptr = end

	if refreshView {
		vi.refreshView()
	}
}

func (vi *view_info) refreshView() {
	v, _ := vi.gui.View("deps")
	v.Clear()
	for idx, line := range vi.lines {
		if idx == vi.cursor {
			fmt.Fprintf(v, "\x1b[0;30;47m")
			fmt.Fprintf(v, line)
			fmt.Fprintf(v, "\x1b[m")
		} else {
			fmt.Fprintf(v, line)
		}
		fmt.Fprintf(v, "\n")
	}
}
