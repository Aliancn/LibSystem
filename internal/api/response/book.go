package response

type BookVO struct {
	BookId      int64  `json:"bookId"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
	Status      string `json:"status"`
	Location    string `json:"location"`
	BorrowTimes int    `json:"borrowTimes"`
}

type BookList struct {
	BookList []BookVO `json:"bookList"`
}
