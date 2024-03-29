package repository

import (
	"clean-code/model"

	"github.com/jmoiron/sqlx"
)

type IStudentRepository interface {
	CreateOne(student model.Student) (*model.Student, error)
	GetAll() ([]model.Student, error)
	GetOneByName(name string) ([]model.Student, error)
	GetOneById(idCard string) (*model.Student, error)
	GetAllByPaging(limit int, offset int) ([]model.Student, error)
}

type StudentRepository struct {
	db *sqlx.DB
}

// Kalau pointer, bisa nil, kalau ga pointer harus struct kosong
func (s *StudentRepository) CreateOne(student model.Student) (*model.Student, error) {
	lastInsertId := 0
	err := s.db.QueryRow("INSERT INTO M_STUDENT (name, gender, age, join_date, id_card, senior) VALUES($1,$2,$3,$4,$5,$6) RETURNING id", student.Name, student.Gender, student.Age, student.JoinDate, student.IdCard, student.Senior).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}
	student.Id = lastInsertId
	return &student, nil
}

// Constructor IStudentRepository untuk connect DB
func NewStudentRepository(resource *sqlx.DB) IStudentRepository {
	return &StudentRepository{db: resource}
}

func (s *StudentRepository) GetAll() ([]model.Student, error) {
	students := []model.Student{}
	err := s.db.Select(&students, "SELECT * FROM M_STUDENT")
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentRepository) GetOneByName(name string) ([]model.Student, error) {
	students := []model.Student{}
	err := s.db.Select(&students, "SELECT * FROM M_STUDENT WHERE name like '%$1%'", name)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentRepository) GetOneById(idCard string) (*model.Student, error) {
	student := model.Student{}
	err := s.db.Get(&student, "SELECT * FROM m_student WHERE id_card=$1", idCard)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *StudentRepository) GetAllByPaging(limit int, offset int) ([]model.Student, error) {
	students := []model.Student{}
	err := s.db.Select(&students, "SELECT * FROM m_student LIMIT $1 OFFSET $2;", limit, offset)
	if err != nil {
		return nil, err
	}
	return students, nil
}
