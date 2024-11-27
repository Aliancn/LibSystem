package response

type InfoResponse struct {
	BookNum     int            `json:"book_num"`
	UserNum     int            `json:"user_num"`
	PaperNum    int            `json:"paper_num"`
	BorrowNum   map[string]int `json:"borrow_num"`
	DownloadNum map[string]int `json:"download_num"`
}
