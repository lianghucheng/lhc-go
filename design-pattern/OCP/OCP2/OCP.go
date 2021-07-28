package OCP1

import "fmt"

type iBook interface {
	GetAuthor() string
	GetName() string
	GetPrice() int
}

type NovelBook struct {
	Author string
	Name string
	Price int
}

type BookStore struct {

}

func NewNovelBook(author, name string, price int) *NovelBook {
	return &NovelBook{
		Author:author,
		Name:name,
		Price:price,
	}
}

func (nb *NovelBook)GetAuthor() string {
	return nb.Author
}

func (nb *NovelBook)GetName() string {
	return nb.Name
}

func (nb *NovelBook)GetPrice() int {
	return nb.Price
}

func (bs *BookStore) ShowBooks() {
	var books = []iBook{
		NewNovelBook("作者1", "书1", 100),
		NewNovelBook("作者1", "书1", 100),
		NewNovelBook("作者1", "书1", 100),
		NewNovelBook("作者1", "书1", 100),
		NewNovelBook("作者1", "书1", 100),
		NewNovelBook("作者1", "书1", 100),
		NewNovelBook("作者1", "书1", 100),
		NewNovelBook("作者1", "书1", 100),
	}
	for _, book := range books {
		fmt.Printf("书名：%v，作者：%v，价格：%v\n", book.GetName(), book.GetAuthor(), book.GetPrice())
	}
}

type OffNovelBook struct {
	NovelBook
}

func (ob *OffNovelBook) GetPrice() (ret int) {
	if ob.Price > 100 {
		ret = int(float64(ob.Price) * 0.8)
	} else {
		ret = int(float64(ob.Price) * 0.6)
	}
	return
}