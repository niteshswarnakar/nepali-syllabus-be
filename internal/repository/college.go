package repository

import "github.com/niteshswarnakar/nepali-syllabus-be/infrastructure"

type ICollegeRepository interface {
	CollegesRepository() ([]string, error)
}

type collegeRepository struct {
	datasource *infrastructure.DatabaseConnection
	env        *infrastructure.Env
}

func (c *collegeRepository) CollegesRepository() ([]string, error) {
	return []string{"College 1", "College 2"}, nil
}

func NewCollegeRepository(datasource *infrastructure.DatabaseConnection, env *infrastructure.Env) ICollegeRepository {
	return &collegeRepository{
		datasource: datasource,
		env:        env,
	}
}
