package repository

import (
	"fmt"
	"service1/config"
	"service1/dto"
	"service1/proto"
)

func FindBookByIdAuthor(request *proto.FindBookByIdAuthorRequest) ([]*proto.FindBookByIdAuthorResponse, error) {
	response := []*proto.FindBookByIdAuthorResponse{}
	var res []dto.Book
	if err := config.DB.Table("book").Select("book.IdBook,book.Name").Where("book_author.idAuthor = ?", request.IdAuthor).Joins("JOIN book_author on book.idBook = book_author.idBook").Find(&res).Error; err != nil {
		return response, err
	}
	fmt.Print(res)
	for _, book := range res {
		bookProto := &proto.FindBookByIdAuthorResponse{}
		bookProto.IdBook = book.IdBook
		bookProto.NameBook = book.Name
		response = append(response, bookProto)
	}
	return response, nil
}
