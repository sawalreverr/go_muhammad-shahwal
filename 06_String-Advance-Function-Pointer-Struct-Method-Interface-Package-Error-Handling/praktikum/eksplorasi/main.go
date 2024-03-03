package main

import (
	"fmt"
)

type Novel struct {
	Title string
	Genre string
}

type BookShelf struct {
	Stack []Novel
}

func (b *BookShelf) Push(nov Novel) {
	if !b.IsEmpty() {
		foundNovel := false
		for _, item := range b.Stack {
			if item.Title == nov.Title {
				foundNovel = true
				break
			}
		}
		if !foundNovel {
			b.Stack = append(b.Stack, nov)
			fmt.Printf("Log: Novel %v ditumpuk\n", nov.Title)
		} else {
			fmt.Printf("Err: Novel %v sudah ada ditumpukan\n", nov.Title)
		}
	} else {
		b.Stack = append(b.Stack, nov)
		fmt.Printf("Log: Novel %v ditumpuk\n", nov.Title)
	}
}

func (b *BookShelf) Pop() {
	if !b.IsEmpty() {
		lastNovel := b.Stack[len(b.Stack)-1].Title
		b.Stack = b.Stack[:len(b.Stack)-1]
		fmt.Printf("Log: %v berhasil diambil\n", lastNovel)
	} else {
		fmt.Println("Err: Bookshelf kosong")
	}
}

func (b *BookShelf) IsEmpty() bool {
	if len(b.Stack) == 0 {
		return true
	} else {
		return false
	}
}

func (b *BookShelf) Values() []Novel {
	return b.Stack
}

func main() {
	var RakBuku BookShelf

	RakBuku.Push(Novel{"Youkoso Jitsuryoku Shijou Shugi no Kyoushitsu e Vol 1", "School, Drama"})
	RakBuku.Push(Novel{"Youkoso Jitsuryoku Shijou Shugi no Kyoushitsu e Vol 2", "School, Drama"})
	RakBuku.Push(Novel{"Overlord Vol 1", "School, Drama"})
	RakBuku.Push(Novel{"Overlord Vol 2", "Fantasy, Magic"})

	RakBuku.Pop()

	isEmptyStack := RakBuku.IsEmpty()
	fmt.Println(isEmptyStack)

	stackNovel := RakBuku.Values()
	fmt.Println(stackNovel)
}
