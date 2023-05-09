package cubery

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParseRoute(t *testing.T) {
	ok := reflect.DeepEqual(parseRoute("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parseRoute("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parseRoute("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parseRoute failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/dsxg")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.route != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if ps["name"] != "dsxg" {
		t.Fatal("name should be equal to 'dsxg'")
	}
}

func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.getRoute("GET", "/assets/file1.txt")
	ok1 := n1.route == "/assets/*filepath" && ps1["filepath"] == "file1.txt"
	if !ok1 {
		t.Fatal("route shoule be /assets/*filepath & filepath shoule be file1.txt")
	}
	n2, ps2 := r.getRoute("GET", "/assets/css/test.css")
	ok2 := n2.route == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("route shoule be /assets/*filepath & filepath shoule be css/test.css")
	}
}

func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}
	if len(nodes) != 5 {
		t.Fatal("the number of routes shoule be 4")
	}
}
