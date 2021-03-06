package cudr

import (
	"EBook/internal/model"
	"EBook/internal/types"
	"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

func GetCountBooks(Count int, DB *gorm.DB) ([]model.Ebook, error) {
	var books []model.Ebook
	if err := DB.Limit(Count).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetPageSizeBooks(page, pagesize int, DB *gorm.DB) ([]model.Ebook, error) {
	var books []model.Ebook
	off := page * (pagesize - 1)
	if err := DB.Offset(off).Limit(pagesize).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func DownloadBookByID(name string, id uint, DB *gorm.DB) (*model.Ebook, error) {
	var book model.Ebook
	if err := DB.Where("bookname=? AND id=?", name, id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func CreateBook(req types.BookUploadReq, DB *gorm.DB) error {
	book := &model.Ebook{
		Bookname:    req.Bookname,
		Author:      req.Author,
		Uploader:    req.Uploader,
		Publisher:   req.Publisher,
		PublishDate: req.PublishDate,
		Version:     req.Version,
		Cover:       req.Cover,
		BookAllname: req.BookAllname,
		StoreAddr:   req.StoreAddr,
		UUID:        uuid.NewV4(),
	}
	if err := DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func GetTotalC(DB *gorm.DB) (int, error) {
	books := &model.Ebook{}
	var count int64
	if err := DB.Model(books).Count(&count).Error; err != nil {
		return 0, nil
	}
	return int(count), nil
}
