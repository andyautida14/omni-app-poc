package handler

import (
	"errors"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"
	"sync"
)

var ERR_INVALID_TMPL_DIR = errors.New("Invalid template directory")

var ERR_INVALID_CACHE_TYPE = errors.New("Invalid cache type")

func GetStaticFilesFs(embedded fs.FS, staticPathConfig string) (http.FileSystem, error) {
	if staticPathConfig != "" {
		return http.Dir(staticPathConfig), nil
	}

	staticFiles, err := fs.Sub(embedded, "static")
	if err != nil {
		return nil, err
	}

	return http.FS(staticFiles), nil
}

func GetTmplFilesFs(embedded fs.FS, tmplPathConfig string) (fs.FS, error) {
	if tmplPathConfig == "" {
		return fs.Sub(embedded, "templates")
	}

	fInfo, err := os.Stat(tmplPathConfig)
	if err != nil {
		return nil, err
	}

	if !fInfo.IsDir() {
		return nil, ERR_INVALID_TMPL_DIR
	}

	return os.DirFS(tmplPathConfig), nil
}

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandleInternalServerError(w http.ResponseWriter, _ *http.Request, err error) {
	slog.Error("an unexpected error occured:" + err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func HandleNotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func TmplMust(tmpl *HtmxTemplate, err error) *HtmxTemplate {
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}

func DSMust[T any](ds interface{}, err error) T {
	if err != nil {
		log.Fatalf("Datastore retrieval error: " + err.Error())
	}

	dsT, ok := ds.(T)
	if !ok {
		log.Fatalf("Datastore type invalid")
	}
	return dsT
}

func newCacheFunc[T interface{}](
	cacheType string,
) (func(func() (T, error)) func() (T, error), error) {
	switch cacheType {
	case "disabled":
		return func(init func() (T, error)) func() (T, error) {
			return init
		}, nil
	case "eager":
		return func(init func() (T, error)) func() (T, error) {
			val, err := init()

			return func() (T, error) {
				return val, err
			}
		}, nil
	case "lazy":
		return func(init func() (T, error)) func() (T, error) {
			var once sync.Once
			var val T
			var err error

			f := func() {
				val, err = init()
			}

			return func() (T, error) {
				once.Do(f)
				return val, err
			}
		}, nil
	}

	return nil, ERR_INVALID_CACHE_TYPE
}
