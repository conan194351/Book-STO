package dto

import "book-sto/errs"

type Message struct {
	Message string `json:"message"`
}

func MessageAddSuccess(obj string) *Message {

	return &Message{

		Message: "Add " + obj + " Success!!!",
	}
}

func MessageCreateSuccess(obj string) *Message {

	return &Message{

		Message: "Create " + obj + " Success!!!",
	}
}

func CheckID(id int) *errs.AppError {

	if id == 0 {

		return errs.BadRequestError("Lost information!!!")
	}

	return nil
}

func LoginSuccess(obj string) *Message {
	return &Message{

		Message: "Login " + obj + " Success!!!",
	}
}

func LoginFalse() *Message {
	return &Message{

		Message: "Login failed",
	}
}

func NotPermissions() *Message {
	return &Message{

		Message: "You have no permissions",
	}
}
