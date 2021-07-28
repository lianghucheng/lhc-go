package OCP2

import "fmt"

type iBook interface {
	GetAuthor() string
	GetName() string
	GetPrice() int
}

type BookBase struct {
	Author string
	Name string
	Price int
}

func (bi *BookBase) GetAuthor() string {
	return bi.Author
}

func (bi *BookBase)GetName() string {
	return bi.Name
}

func (bi *BookBase)GetPrice() int {
	return bi.Price
}


type NovelBook struct {
	BookBase
}

func NewNovelBook(info *BookBase) *NovelBook {
	return &NovelBook{BookBase: *info}
}

type OffNovelBook struct {
	NovelBook
}

func NewOffNovelBook(info *BookBase) *OffNovelBook {
	ret := &OffNovelBook{}
	ret.BookBase = *info
	return ret
}

func (ob *OffNovelBook) GetPrice() (ret int) {
	if ob.Price > 100 {
		ret = int(float64(ob.Price) * 0.8)
	} else {
		ret = int(float64(ob.Price) * 0.6)
	}
	return
}

type iComputeBook interface {
	iBook
	GetScope() string
}

type ScopeBookBase struct {
	BookBase
	Scope string
}

func (si *ScopeBookBase) GetScope() string {
	return si.Scope
}

type ComputeBook struct {
	ScopeBookBase
}

func NewComputeBook(info *ScopeBookBase) *ComputeBook {
	return &ComputeBook{ScopeBookBase: *info}
}

type BookStore struct {

}

func (bs *BookStore) ShowBooks() {
	var books = []iBook{
		NewNovelBook(&BookBase{"作者1", "书1", 100}),
		NewNovelBook(&BookBase{"作者1", "书1", 100}),
		NewNovelBook(&BookBase{"作者1", "书1", 100}),
		NewNovelBook(&BookBase{"作者1", "书1", 100}),
		NewNovelBook(&BookBase{"作者1", "书1", 100}),
		NewOffNovelBook(&BookBase{"作者1", "书1", 100}),
		NewOffNovelBook(&BookBase{"作者1", "书1", 100}),
		NewOffNovelBook(&BookBase{"作者1", "书1", 100}),
		NewComputeBook(&ScopeBookBase{BookBase:BookBase{"作者1", "书1", 100}, Scope:"aaaa"}),
		NewComputeBook(&ScopeBookBase{BookBase:BookBase{"作者1", "书1", 100}, Scope:"aaaa"}),
	}
	for _, book := range books {
		fmt.Printf("书名：%v，作者：%v，价格：%v\n", book.GetName(), book.GetAuthor(), book.GetPrice())
	}
}