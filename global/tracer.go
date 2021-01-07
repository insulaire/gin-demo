package global

import (
	"gin-demo/pkg/tracer"
	"log"

	"github.com/opentracing/opentracing-go"
)

var (
	Tracer opentracing.Tracer
)

func init() {
	if err := setupTracer(); err != nil {
		log.Fatalln(err)
	}
}

func setupTracer() error {
	log.Println("开始注册Tracer")
	jaegerTracer, _, err := tracer.NewJaegerTracer("Test", "127.0.0.1:6831")
	if err != nil {
		return err
	}

	Tracer = jaegerTracer
	log.Println("完成注册Tracer")
	return nil
}
