package resp

import "eshop_user/model"

type LoginRespDTO struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

type RegisterRespDTO struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

type UserRespDTO struct {
	Code int         `json:"code"`
	Info string      `json:"info"`
	User *model.User `json:"user"`
}
