package main

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"encoding/json"
	"bytes"
	"fmt"
)

func main() {
	router := routing.New()
	router.Get("/", func (c *routing.Context) error {
		data := ProcessRequest(c)

		res, err := json.Marshal(data)
		if err != nil {
			c.Error(err.Error(), 500)
			return err
		}

		// JSON beautification
		// Can be removed if only programmatic access suggested
		var out bytes.Buffer
		json.Indent(&out, res, "", "\t")

		c.SetBody(out.Bytes())
		return nil
	})
	fmt.Println(">>> Server is on 8080")
	panic(fasthttp.ListenAndServe("0.0.0.0:8080", router.HandleRequest))
}

// This can be replaced to separate file
// Also this can be replaced by multiple middlewares
func ProcessRequest(c *routing.Context) map[string]interface{} {
	data := make(map[string]interface {})

	if c.QueryArgs().Has("ip") {
		data["ip"] = c.RemoteIP().String()
	}

	if c.QueryArgs().Has("referrer") {
		ref := c.Request.Header.Peek("Referer")
		data["referrer"] = fmt.Sprintf("%s", ref)
	}

	if c.QueryArgs().Has("user-agent") {
		data["agent"] = fmt.Sprintf("%s", c.UserAgent())
	}

	return data
} 