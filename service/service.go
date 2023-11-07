package service

import (
	"membuattasktodo/helpers"
	"membuattasktodo/model"
	"membuattasktodo/repository"
	"time"
)

type Service interface {
	GetUsers() ([]model.User, error)
	GetAlltask(id int) ([]model.TaskRes, error)
	GetTaskById(id int, taskId int) (model.TaskRes, error)
	CreateTask(req model.TaskReq, Id int, ImageURL string) (model.TaskRes, error)
	UpdateTask(req model.TaskReq, ImageURL string, Id int) (model.TaskRes, error)
	DeleteTask(Id int) error
	BulkDeleteTask(taskIds []int, Id int) error
	Login(email string, password string) (model.UserLogRespon, error)
	Regis(email string, password string) (model.UserRegisRespon, error)
	SaveToken(token string, userId int) error

	//KATEGORI
	GetAllKategori() ([]model.Kategori, error)
	CreateKategori(req model.KategoriReq) (model.Kategori, error)
	DeleteKategori(Id int) error
	EditKategori(Id int, kategori model.KategoriReq) (model.Kategori, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{repository}
}

func (s *service) GetUsers() ([]model.User, error) {
	data, err := s.repository.GetUsers()
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (s *service) GetAlltask(id int) ([]model.TaskRes, error) {
	data, err := s.repository.GetAlltask(id)
	if err != nil {
		return []model.TaskRes{}, err
	}
	return data, nil
}

func (s *service) GetTaskById(id int, taskId int) (model.TaskRes, error) {
	data, err := s.repository.GetTaskById(id, taskId)
	if err != nil {
		return model.TaskRes{}, err
	}

	return data, nil
}

func (s *service) DeleteTask(Id int) error {
	err := s.repository.DeleteTask(Id)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) BulkDeleteTask(taskIds []int, Id int) error {
	err := s.repository.BulkDeleteTask(taskIds, Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateTask(req model.TaskReq, Id int, ImageURL string) (model.TaskRes, error) {
	layout := "2006-01-02 15:04"
	parsedDate, err := time.Parse(layout, req.Date)
	if err != nil {
		return model.TaskRes{}, err
	}

	data, err := s.repository.CreateTask(req, parsedDate, ImageURL, Id)
	if err != nil {
		return model.TaskRes{}, err
	}

	return data, nil
}

func (s *service) UpdateTask(req model.TaskReq, ImageURL string, Id int) (model.TaskRes, error) {
	layout := "2006-01-02 15:04"
	parsedDate, err := time.Parse(layout, req.Date)
	if err != nil {
		return model.TaskRes{}, err
	}

	data, err := s.repository.UpdateTasks(req, parsedDate, ImageURL, Id)
	if err != nil {
		return model.TaskRes{}, err
	}

	return data, nil
}

//AUTH
func (s *service) Login(email string, password string) (model.UserLogRespon, error) {
	data, err := s.repository.Login(email)
	if err != nil {
		return model.UserLogRespon{}, err
	}

	match, err := helpers.ComparePassword(data.Password, password)
	if err != nil {
		return model.UserLogRespon{}, err
	}
	if !match {
		return model.UserLogRespon{}, err
	}
	return data, nil
}

func (s *service) Regis(email string, password string) (model.UserRegisRespon, error) {
	data, err := s.repository.Regis(email, password)
	if err != nil {
		return model.UserRegisRespon{}, err
	}

	password, err = helpers.HashPassword(password)
		if err != nil {
			return model.UserRegisRespon{}, err
		}

	return data, nil
}

func (s *service) SaveToken(token string, userId int) error {
	err := s.repository.SaveToken(token, userId)
	if err != nil {
		return err
	}

	return nil
}



// KATEGORI
func (s *service) GetAllKategori() ([]model.Kategori, error) {
	data, err := s.repository.GetAllKategori()
	if err != nil {
		return []model.Kategori{}, err
	}
	return data, nil
}
func (s *service) CreateKategori(req model.KategoriReq) (model.Kategori, error) {

	kategori, err := s.repository.CreateKategori(req)
	if err != nil {
		return model.Kategori{}, err
	}

	return kategori, nil
}
func (s *service) EditKategori(Id int, req model.KategoriReq) (model.Kategori, error) {

	kategori, err := s.repository.EditKategori(Id, req)
	if err != nil {
		return model.Kategori{}, err
	}

	return kategori, nil
}
func (s *service) DeleteKategori(Id int) error {
	err := s.repository.DeleteKategori(Id)
	if err != nil {
		return err
	}

	return nil
}
