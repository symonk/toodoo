package middleware

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type bodyWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func (bw bodyWriter) Write(b []byte) (int, error) {
	return bw.buf.Write(b)
}

func JsonIndenter(c *gin.Context) {
	// TODO: This does not propagate errors back for aborts.
	bw := &bodyWriter{buf: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = bw
	c.Next()
	old := bw.buf.String()
	obj := make(map[string]any)
	json.Unmarshal([]byte(old), &obj)
	new, _ := json.MarshalIndent(obj, "", "    ")
	// If the response has any json; pretty print it.
	bw.ResponseWriter.WriteString(string(new))
	bw.buf.Reset()
}
