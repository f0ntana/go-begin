package comment

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	DB *gorm.DB
}

type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments, slug).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}

	return comments, nil
}
