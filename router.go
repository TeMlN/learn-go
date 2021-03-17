package main

import (
	"net/http"
	"strings"
) 

type router struct {
	// 키 : http메서드
	// 값 : URL 패턴별로 실행할 HandlerFunc
	handlers map[string]map[string]HandlerFunc
}

func (c *router) HandlerFunc(method, pattern string, h HandlerFunc) {
	// http 메서드로 등록된 맵이 있는지 확인.
	m, ok := r.handlers[method]
	if !ok {
		// 등록된 맵이 없으면 새 맵을 생성.
		m = make(map[string]HandlerFunc)
		r.handlers[method] = m
	}
	// http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수등록.
	m[pattern] = h
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// http 메서드에 맞는 모든 handers를 반복하여 요청 URL에 해당하는 handler를 찾음
	for pattern, handler := range r.handlers[req.Method] {
		if ok, params := match(pattern, req.URL.Path); ok {
			// Context 생성
			c := Context {
				params:			make(map[string]interface{}),
				ResponseWriter: w,
				Request:		req,
			}
			for k, v := range params {
				c.Params[k] = v
			}
			//요청 URL에 해당하는 handler수행
			handler(&c)
			return
		}
	}
	// 요청 URL에 해당하는 handler를 찾지 못하면 NotFound 에러 처리
	http.NotFound(w, req)
	return
}

func match(pattern, path string) (bool, map[string]string) {
	// ...
}