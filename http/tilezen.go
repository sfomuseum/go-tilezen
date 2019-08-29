package http

import (
	gohttp "net/http"
	"github.com/sfomuseum/go-tilezen"
	"github.com/whosonfirst/go-whosonfirst-cache"	
	"io"
)

type TilezenProxyOptions struct {
	Cache cache.Cache
}

func TilezenProxyHandler(proxy_opts *TilezenProxyOptions) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		path := req.URL.Path

		tile, err := tilezen.ParseURI(path)
		
		if err != nil {
			gohttp.Error(rsp, "Invalid path", gohttp.StatusBadRequest)
			return
		}

		q := req.URL.Query()

		api_key := q.Get("api_key")

		if api_key == "" {
			gohttp.Error(rsp, "Missing API key", gohttp.StatusBadRequest)
			return
		}
		
		tilezen_opts := &tilezen.Options{
			ApiKey: api_key,
		}
		
		t_rsp, err := tilezen.FetchTileWithCache(proxy_opts.Cache, tile, tilezen_opts)
		
		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)						
			// gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		_, err = io.Copy(rsp, t_rsp)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		return
	}

	return gohttp.HandlerFunc(fn), nil	
}
