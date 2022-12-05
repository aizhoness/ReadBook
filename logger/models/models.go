package models

// Book info
type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"authorid"`
}

// BooksResp is the response for the GET /books endpoint
type BooksResp struct {
	Books []Book `json:"books"`
}

//Author info
type Author struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

//Reader info
type Reader struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

//Mylibrary as the books that are read
type MyLibrary struct {
	ReaderID  int    `json:"readerid"`
	ReadBooks []Book `json:"books"`
}
