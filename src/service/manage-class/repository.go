package manage_class

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"shub_go/src/enums"
	"shub_go/src/models"
	"strings"
	"time"
)

type IRepository interface {
	GetAllSubjects() ([]models.Subject, error)
	GetAllGrades() ([]models.Grade, error)
	CreateClass(ctx context.Context, class *models.Class) error
	FindByID(ctx context.Context, id int) (*models.Class, error)
	FindByCode(ctx context.Context, code string) (*models.Class, error)
	QueryByUserId(ctx context.Context, userId int, params QueryClass) ([]models.Class, error)
	InsertUserClass(ctx context.Context, userId int, classId int, role enums.RoleClass) error
	FindByStudentAndClass(ctx context.Context, studentId int, classId int) (*models.UserClass, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindByStudentAndClass(ctx context.Context, studentId int, classId int) (*models.UserClass, error) {
	var output models.UserClass
	err := r.db.Raw("SELECT * FROM user_class WHERE user_id = ? AND class_id = ?", studentId, classId).Scan(&output).Error

	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (r *repository) InsertUserClass(ctx context.Context, userId int, classId int, role enums.RoleClass) error {

	err := r.db.Exec("INSERT INTO user_class (user_id,class_id,role,created_at,updated_at) VALUES (?,?,?,?,?)", userId, classId, role, time.Now(), time.Now()).Error

	return err
}

func (r *repository) GetAllSubjects() ([]models.Subject, error) {
	var result []models.Subject

	err := r.db.Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) GetAllGrades() ([]models.Grade, error) {
	var result []models.Grade

	err := r.db.Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) CreateClass(ctx context.Context, class *models.Class) error {

	err := r.db.Create(&class).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByID(ctx context.Context, id int) (*models.Class, error) {
	var class models.Class

	err := r.db.Joins("Grade").Joins("Subject").First(&class, id).Error

	if err != nil {
		return nil, err
	}

	return &class, nil

}

func (r *repository) FindByCode(ctx context.Context, code string) (*models.Class, error) {
	var class models.Class

	err := r.db.Where("code = ?", code).First(&class).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *repository) QueryByUserId(ctx context.Context, userId int, params QueryClass) ([]models.Class, error) {

	query := r.db.Model(&models.Class{}).Where("user_id = ?", userId)

	if params.name != nil {
		queryLike := fmt.Sprintf("%s%s%s", "%", *params.name, "%")
		query = query.Where("name LIKE ? ", queryLike)
	}

	if params.orderBy != nil {
		order := fmt.Sprintf("%s %s", strings.ToLower(*params.orderBy), strings.ToLower(string(params.direction)))
		query = query.Order(order)
	}

	var result []models.Class

	err := query.Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil

}
