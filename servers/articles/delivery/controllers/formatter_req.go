package articles

import (
	helpers "go-grpc-micro/servers/articles/utils/generic"
	"log"
)

func Req_logger(method, after_id, id, content, title, user_id, order, sort string, page int32) {
	helpers.Infof("Incomming " +
		method + " Request" +
		"\n after_id = " + after_id +
		"\n id = " + id +
		"\n content = " + content +
		"\n title = " + title +
		"\n user_id = " + user_id +
		"\n order = " + order +
		"\n sort = " + sort +
		"\n page = " + string(page))
	log.Print(&helpers.Buf)
}
