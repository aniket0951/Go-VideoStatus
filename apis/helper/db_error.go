package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

func HandleDBErr(err error) (err_ error) {
	switch e := err.(type) {
	case *pq.Error:
		switch e.Code {
		case "23502":
			// not-null constraint violation
			fmt.Println("Some required data was left out:\n\n", e.Message)
		case "23505":
			// unique constraint violation
			err_ = errors.New(e.Detail)
			return
		case "23514":
			// check constraint violation
			if strings.Contains(e.Message, "contact") {
				err_ = errors.New("contact should not be empty")
				return
			} else if strings.Contains(e.Message, "email") {
				err_ = errors.New("email should not be empty")
				return
			}
		case "23503":
			err_ = errors.New("invalid id has been provided,please try with valid id's")
			return
		default:
			msg := e.Message
			if d := e.Detail; d != "" {
				msg += "\n\n" + d
			}
			if h := e.Hint; h != "" {
				msg += "\n\n" + h
			}
			fmt.Println("Message from default : ", e.Code)
			err_ = errors.New(msg)
			return
		}
	default:
		err_ = nil
		return
	}

	return
}
