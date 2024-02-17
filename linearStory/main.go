package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

func playStore(page *storyPage) {
	if page == nil {
		return
	}

	fmt.Println(page.text)
	playStore(page.nextPage)
}

func main() {

	page1 := storyPage{"It was a dark and stormy night.", nil}
	page2 := storyPage{"You are alone, and you need to find the sacred helmet before the bad guys do.", nil}
	page3 := storyPage{"You see a troll ahead.", nil}
	page4 := storyPage{"Hello my name is Rodrigo.", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3
	page3.nextPage = &page4

	playStore(&page1)
}
