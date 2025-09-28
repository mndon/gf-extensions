package model

import "github.com/mndon/gf-extensions/kvconfigx"

type KvConfigListInput struct {
	ProjectId int    `json:"project_id"`
	K         string `json:"k"`
	Page      int    `json:"page" d:"1"`
	Size      int    `json:"size" d:"20"`
}

type KvConfigListOutput struct {
	List  []kvconfigx.KvConfigEntity `json:"list"`
	Page  int                        `json:"page" `
	Size  int                        `json:"size" `
	Total int                        `json:"total" `
}
