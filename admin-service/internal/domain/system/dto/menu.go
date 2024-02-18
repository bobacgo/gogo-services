package dto

import "github.com/gogoclouds/gogo-services/admin-service/internal/model"

type MenuNode struct {
	model.Menu
	Children []*model.Menu `json:children`
}
