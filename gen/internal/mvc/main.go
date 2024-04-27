package main

import (
	"log"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gogoclouds/gogo-services/framework/pkg"
)

func main() {
	var module = struct {
		RootPath   string // ../
		Project    string // github.com/gogoclouds/gogo-service
		Service    string // admin-service
		Module     string // system
		Struct     string // User
		Domain     string // internal/domain/system
		ApiVersion string // v1
		File       string // user
		Model      string // sql table
	}{
		RootPath: "../",
		Project:  "github.com/gogoclouds/gogo-services",
		Service:  "admin-service",
		Module:   "system",
		Struct:   "Menu",
		Domain:   "internal/domain",
	}
	if module.ApiVersion == "" {
		module.ApiVersion = "v1"
	}
	if module.File == "" {
		module.File = strings.ToLower(module.Struct)
	}
	if module.Model == "" {
		module.Model = module.Struct
	}
	// github.com/gogoclouds/gogo-services/admin-service
	module.Project = filepath.Join(module.Project, module.Service)
	filepaths := []string{
		"./internal/mvc/tmpl/router.tmpl",
		"./internal/mvc/tmpl/api.tmpl",
		"./internal/mvc/tmpl/handler.tmpl",
		"./internal/mvc/tmpl/service.tmpl",
		"./internal/mvc/tmpl/repo.tmpl",
	}
	for _, fp := range filepaths {
		tp, err := template.ParseFiles(fp)
		if err != nil {
			log.Println(err)
			return
		}

		var (
			// admin-service/internal/domain/system
			dir = filepath.Join(module.Service, module.Domain, module.Module)
			// user.go
			filename = module.File + ".go"
			// router | api | handler | service | repo
			fType = strings.ToLower(strings.TrimSuffix(filepath.Base(fp), filepath.Ext(fp)))
		)

		switch fType {
		case "router":
			filename = "router.go"
		case "api":
			// ../admin-service/api/system/v1
			dir = filepath.Join(module.Service, "api", module.Module, module.ApiVersion)
		default:
			// ../admin-service/internal/domain/system
			dir = filepath.Join(dir, fType)
		}
		f, err := pkg.MustOpen(filename, filepath.Join(module.RootPath, dir))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err = tp.Execute(f, module); err != nil {
			log.Println(err)
		}
	}
}
