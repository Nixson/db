package db

import (
	"github.com/Nixson/db/postgres"
	"github.com/Nixson/environment"
	"gorm.io/gorm"
	"math"
)

var gormInstance *gorm.DB

func InitDb() {
	env := environment.GetEnv()
	switch env.GetString("db.driver") {
	case "postgres":
		postgres.InitDb()
		gormInstance = postgres.Get()
	}
}

func Get() *gorm.DB {
	if gormInstance == nil {
		InitDb()
	}
	return gormInstance
}

type Pagination struct {
	Page       int
	Size       int
	TotalRows  int64
	TotalPages int64
}

type PaginationResponse struct {
	Content  []interface{} `json:"content"`
	Pageable struct {
		Sort struct {
			Sorted   bool `json:"sorted"`
			Unsorted bool `json:"unsorted"`
			Empty    bool `json:"empty"`
		} `json:"sort"`
		PageNumber int  `json:"pageNumber"`
		PageSize   int  `json:"pageSize"`
		Offset     int  `json:"offset"`
		Paged      bool `json:"paged"`
		Unpaged    bool `json:"unpaged"`
	} `json:"pageable"`
	TotalPages    int64 `json:"totalPages"`
	TotalElements int64 `json:"totalElements"`
	Last          bool  `json:"last"`
	Sort          struct {
		Sorted   bool `json:"sorted"`
		Unsorted bool `json:"unsorted"`
		Empty    bool `json:"empty"`
	} `json:"sort"`
	NumberOfElements int  `json:"numberOfElements"`
	First            bool `json:"first"`
	Size             int  `json:"size"`
	Number           int  `json:"number"`
	Empty            bool `json:"empty"`
}

func Paginate(page, size int, p *Pagination, t interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}

		offset := (page - 1) * size

		Get().Model(t).Count(&p.TotalRows)
		p.TotalPages = int64(math.Ceil(float64(p.TotalRows) / float64(size)))
		p.Page = page
		p.Size = size

		return db.Offset(offset).Limit(size)
	}
}
