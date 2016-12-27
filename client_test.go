// by liudan
package httpclient

import (
	"io/ioutil"
	"net/http"
	"strings"

	"testing"
)

func TestDo(t *testing.T) {
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := HttpClientClusterDo(req)
	if err != nil {
		t.Fatal(err)
	}
	defer rsp.Body.Close()
	data, _ := ioutil.ReadAll(rsp.Body)
	// t.Logf("rsp:%s", data)

	req, err = http.NewRequest("POST", "http://www.baidu.com", strings.NewReader("baidu=fxck"))
	if err != nil {
		t.Fatal(err)
	}
	rsp, err = HttpClientClusterDo(req)
	if err != nil {
		t.Fatal(err)
	}
	defer rsp.Body.Close()
	data, _ = ioutil.ReadAll(rsp.Body)
	t.Logf("rsp:%s", data)
}
