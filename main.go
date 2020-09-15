// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", test)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func test(w http.ResponseWriter, r *http.Request) {

	data := make(map[interface{}]interface{})

	var content map[string]interface{}

	data["Content"] = content

	execTpl(w, data, testTpl)
}

// beego: https://github.com/astaxie/beego
func execTpl(rw http.ResponseWriter, data map[interface{}]interface{}, tpls ...string) {
	tmpl := template.Must(template.New("main").Parse(mainTpl))
	for _, tpl := range tpls {
		tmpl = template.Must(tmpl.Parse(tpl))
	}
	tmpl.Execute(rw, data)
}

var mainTpl = `
<html>

<body>

{{template "content" .}}

</body>
</html>
`

var testTpl = `
{{define "content"}}

<form>

测试输入：
<input type="text" name="input1">

</form>

{{end}}
`
