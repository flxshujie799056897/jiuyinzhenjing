//+build osinfo

package main
import (
	"os"
	"fmt"
	"net/http"
	"../html"
)
func Get_plugin_name()(pluginname string){
	pluginname = "/osinfo"
	return
}
func Func_plugin(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"<html><body>")
	html.Title(w,"os info")
	fmt.Fprintf(w,"<h1>info</h1>")
	fmt.Fprintf(w,"<h2>pwd:%s<h2>",os.Getenv("PWD"))
	fmt.Fprintf(w,"</body></html>")
	return
}