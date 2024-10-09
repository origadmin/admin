package errors

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/toolkits/errors/httperr"
)

type Error = httperr.Error

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	var reply Error
	//if err == nil {
	//	return
	//}
	if !errors.As(err, &reply) {
		se := errors.FromError(err)
		reply.ID = se.Message
		reply.Code = se.Code
		reply.Detail = se.Reason

	}

	codec, _ := transhttp.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(reply)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contentType(codec.Name()))
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	//reply := NewResponse()
	//reply.Code = 200
	//reply.Data = v
	//reply.Message = "success"
	//reply.Reason = "success"
	//reply.Ts = time.Now().Format(pkgTime.MilliTimeLayout)
	reply := v

	codec, _ := transhttp.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(reply)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", contentType(codec.Name()))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return nil
}

func contentType(name string) string {
	return "application/" + name
}

func New() []transhttp.ServerOption {
	var opts []transhttp.ServerOption
	// Error decoder
	opts = append(opts, transhttp.ErrorEncoder(ErrorEncoder))
	// Returns the parameter decoder
	opts = append(opts, transhttp.ResponseEncoder(ResponseEncoder))
	return opts
}
