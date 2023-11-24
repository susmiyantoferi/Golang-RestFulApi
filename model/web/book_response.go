package web

//response disini bisa memilih kolom mana saja yang ingin ditampilkan pada hasil response API nanti
//apabila ada kolom password tidak perlu ditampilkan ke hasil response API, karena bersifat prvasi

type BookResponse struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Descrip string `json:"descrip"`
}
