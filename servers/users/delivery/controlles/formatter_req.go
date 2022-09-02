package users

import (
	helpers "go-grpc-micro/servers/users/utils/generic"
	"log"
)

func Req_logger(method, after_id, id, image_url, name, order, sort string, page int32) {
	helpers.Infof("Incomming " +
		method + " Request" +
		"\n after_id = " + after_id +
		"\n id = " + id +
		"\n image_url = " + image_url +
		"\n name = " + name +
		"\n order = " + order +
		"\n sort = " + sort +
		"\n page = " + string(page))
	log.Print(&helpers.Buf)
}
