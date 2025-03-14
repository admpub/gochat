/**
 * Created by lock
 * Date: 2019-08-12
 * Time: 11:36
 */
package site

import (
	"fmt"
	"gochat/config"
	"net/http"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type Site struct {
}

func New() *Site {
	return &Site{}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	// Here you can send your custom 404 back.
	data, _ := os.ReadFile("./site/index.html")
	w.Write(data)
}

func server(fs http.FileSystem) http.Handler {
	fileServer := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := path.Clean("./site" + r.URL.Path)
		_, err := os.Stat(filePath)
		if err != nil {
			notFound(w, r)
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}

func (s *Site) Run() {
	siteConfig := config.Conf.Site
	port := siteConfig.SiteBase.ListenPort
	addr := fmt.Sprintf(":%d", port)
	logrus.Fatal(http.ListenAndServe(addr, server(http.Dir("./site"))))
}
