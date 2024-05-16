package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"gte=17,lte=56"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Tittle      string `json:"tittle" binding:"min=2,max=22"`
	Description string `json:"description" binding:"min=22,max=222"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
