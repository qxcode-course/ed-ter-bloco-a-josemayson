package main

import (
	"fmt"
)

type Node struct {
	info rune
	next *Node
	prev *Node
}

type Editor struct {
	head   *Node
	cursor *Node
}

func NewEditor() *Editor {
	editor := &Node{}
	editor.next = editor
	editor.prev = editor
	return &Editor{head: editor, cursor: editor}
}

func (editor *Editor) Insert(b *Node, value rune) {
	a := b.prev
	c := &Node{
		info: value,
		next: b,
		prev: a,
	}
	a.next = c
	b.prev = c
	editor.cursor = b
}

func (editor *Editor) Remove(b *Node) {
	a := b.prev
	c := b.next
	a.next = c
	c.prev = a
	b.prev = nil
	b.next = nil
	editor.cursor = c
}

func (editor *Editor) String() string {
	saida := ""

	for it := editor.head.next; it != editor.head; it = it.next {
		if editor.cursor == it {
			saida += "|"
		}
		saida += fmt.Sprintf("%c", it.info)
	}
	if editor.cursor == editor.head {
		saida += "|"
	}
	return saida
}

func (editor *Editor) rigth() {
	if editor.cursor != editor.head {
		editor.cursor = editor.cursor.next
	}
}

func (editor *Editor) left() {
	if editor.cursor.prev != editor.head {
		editor.cursor = editor.cursor.prev
	}
}

func main() {
	var palavra string
	fmt.Scan(&palavra)

	editor := NewEditor()

	for _, c := range palavra {
		switch c {
		case 'R':
			editor.Insert(editor.cursor, '\n')
		case 'D':
			editor.Remove(editor.cursor)
		case 'B':
			editor.Remove(editor.cursor.prev)
		case '>':
			editor.rigth()
		case '<':
			editor.left()
		default:
			editor.Insert(editor.cursor, c)
		}
	}

	fmt.Println(editor)
}
