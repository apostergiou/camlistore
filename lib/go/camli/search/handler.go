/*
Copyright 2011 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package search

import (
	"fmt"
	"http"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"camli/blobref"
	"camli/blobserver"
	"camli/jsonconfig"
	"camli/httputil"
)

func init() {
	blobserver.RegisterHandlerConstructor("search", newHandlerFromConfig)
}

type searchHandler struct {
	index Index
	owner *blobref.BlobRef
}

func newHandlerFromConfig(ld blobserver.Loader, conf jsonconfig.Obj) (http.Handler, os.Error) {
	indexPrefix := conf.RequiredString("index") // TODO: add optional help tips here?
	ownerBlobStr := conf.RequiredString("owner")
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	indexHandler, err := ld.GetHandler(indexPrefix)
	if err != nil {
		return nil, fmt.Errorf("search config references unknown handler %q", indexPrefix)
	}
	indexer, ok := indexHandler.(Index)
	if !ok {
		return nil, fmt.Errorf("search config references invalid indexer %q (actually a %T)", indexPrefix, indexHandler)
	}
	ownerBlobRef := blobref.Parse(ownerBlobStr)
	if ownerBlobRef == nil {
		return nil, fmt.Errorf("search 'owner' has malformed blobref %q; expecting e.g. sha1-xxxxxxxxxxxx",
			ownerBlobStr)
	}
	return &searchHandler{
		index: indexer,
		owner: ownerBlobRef,
	}, nil
}


func jsonMap() map[string]interface{} {
	return make(map[string]interface{})
}

func jsonMapList() []map[string]interface{} {
	return make([]map[string]interface{}, 0)
}

func (sh *searchHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_ = req.Header.Get("X-PrefixHandler-PathBase")
	suffix := req.Header.Get("X-PrefixHandler-PathSuffix")

	if req.Method == "GET" {
		switch suffix {
		case "camli/search", "camli/search/recent":
			sh.serveRecentPermanodes(rw, req)
			return
		case "camli/search/describe":
			sh.serveDescribe(rw, req)
			return
		case "camli/search/claims":
			sh.serveClaims(rw, req)
			return
		case "camli/search/files":
			sh.serveFiles(rw, req)
			return
		}
	}

	// TODO: discovery for the endpoints & better error message with link to discovery info
	ret := jsonMap()
	ret["error"] = "Unsupported search path or method"
	ret["errorType"] = "input"
	httputil.ReturnJson(rw, ret)
}

func (sh *searchHandler) serveRecentPermanodes(rw http.ResponseWriter, req *http.Request) {
	ch := make(chan *Result)
	results := jsonMapList()
	errch := make(chan os.Error)
	go func() {
		log.Printf("finding recent permanodes for %s", sh.owner)
		errch <- sh.index.GetRecentPermanodes(ch, []*blobref.BlobRef{sh.owner}, 50)
	}()

	wg := new(sync.WaitGroup)
	for res := range ch {
		jm := jsonMap()
		jm["blobref"] = res.BlobRef.String()
		jm["owner"] = res.Signer.String()
		t := time.SecondsToUTC(res.LastModTime)
		jm["modtime"] = t.Format(time.RFC3339)
		results = append(results, jm)
		wg.Add(1)
		go func() {
			sh.populatePermanodeFields(jm, res.BlobRef, res.Signer, nil)
			wg.Done()
		}()
	}
	wg.Wait()

	err := <-errch

	ret := jsonMap()
	ret["results"] = results
	if err != nil {
		// TODO: return error status code
		ret["error"] = fmt.Sprintf("%v", err)
	}
	httputil.ReturnJson(rw, ret)
}

func (sh *searchHandler) serveClaims(rw http.ResponseWriter, req *http.Request) {
	ret := jsonMap()

	pn := blobref.Parse(req.FormValue("permanode"))
	if pn == nil {
		http.Error(rw, "Missing or invalid 'permanode' param", 400)
		return
	}

	// TODO: rename GetOwnerClaims to GetClaims?
	claims, err := sh.index.GetOwnerClaims(pn, sh.owner)
	if err != nil {
		log.Printf("Error getting claims of %s: %v", pn.String(), err)
	} else {
		sort.Sort(claims)
		jclaims := jsonMapList()

		for _, claim := range claims {
			jclaim := jsonMap()
			jclaim["blobref"] = claim.BlobRef.String()
			jclaim["signer"] = claim.Signer.String()
			jclaim["permanode"] = claim.Permanode.String()
			jclaim["date"] = claim.Date.Format(time.RFC3339)
			jclaim["type"] = claim.Type
			if claim.Attr != "" {
				jclaim["attr"] = claim.Attr
			}
			if claim.Value != "" {
				jclaim["value"] = claim.Value
			}

			jclaims = append(jclaims, jclaim)
		}
		ret["claims"] = jclaims
	}

	httputil.ReturnJson(rw, ret)
}

func (sh *searchHandler) serveDescribe(rw http.ResponseWriter, req *http.Request) {
	ret := jsonMap()
	defer httputil.ReturnJson(rw, ret)

	br := blobref.Parse(req.FormValue("blobref"))
	if br == nil {
		ret["error"] = "Missing or invalid 'blobref' param"
		ret["errorType"] = "input"
		return
	}

	dmap := func(b *blobref.BlobRef) map[string]interface{} {
		bs := b.String()
		if m, ok := ret[bs]; ok {
			return m.(map[string]interface{})
		}
		m := jsonMap()
		ret[bs] = m
		return m
	}

	mime, size, err := sh.index.GetBlobMimeType(br)
	if err != os.ENOENT {
		if err != nil {
			ret["errorText"] = err.String()
		} else {
			m := dmap(br)
			setMimeType(m, mime)
			m["size"] = size

			if mime == "application/json; camliType=permanode" {
				pm := jsonMap()
				m["permanode"] = pm
				sh.populatePermanodeFields(pm, br, sh.owner, dmap)
			}
		}
	}
}

func (sh *searchHandler) serveFiles(rw http.ResponseWriter, req *http.Request) {
	ret := jsonMap()
	defer httputil.ReturnJson(rw, ret)

	br := blobref.Parse(req.FormValue("bytesref"))
	if br == nil {
		// TODO: formalize how errors are returned And make
		// ReturnJson set the HTTP status to 400 automatically
		// in some cases, if errorType is "input"?  Document
		// this somewhere.  Are there existing JSON
		// conventions to use?
		ret["error"] = "Missing or invalid 'bytesref' param"
		ret["errorType"] = "input"
		return
	}

	files, err := sh.index.ExistingFileSchemas(br)
	if err != nil {
		ret["error"] = err.String()
		ret["errorType"] = "server"
		return
	}

	strList := []string{}
	for _, br := range files {
		strList = append(strList, br.String())
	}
	ret["files"] = strList
	return
}

// dmap may be nil, returns the jsonMap to populate into
func (sh *searchHandler) populatePermanodeFields(jm map[string]interface{}, pn, signer *blobref.BlobRef, dmap func(b *blobref.BlobRef) map[string]interface{}) {
	jm["content"] = ""
	attr := jsonMap()
	jm["attr"] = attr

	claims, err := sh.index.GetOwnerClaims(pn, signer)
	if err != nil {
		log.Printf("Error getting claims of %s: %v", pn.String(), err)
	} else {
		sort.Sort(claims)
	claimLoop:
		for _, cl := range claims {
			switch cl.Type {
			case "del-attribute":
				if cl.Value == "" {
					attr[cl.Attr] = nil, false
				} else {
					sl, ok := attr[cl.Attr].([]string)
					if ok {
						filtered := make([]string, 0, len(sl))
						for _, val := range sl {
							if val != cl.Value {
								filtered = append(filtered, val)
							}
						}
						attr[cl.Attr] = filtered
					}
				}
			case "set-attribute":
				attr[cl.Attr] = nil, false
				fallthrough
			case "add-attribute":
				if cl.Value == "" {
					continue
				}
				sl, ok := attr[cl.Attr].([]string)
				if ok {
					for _, exist := range sl {
						if exist == cl.Value {
							continue claimLoop
						}
					}
				} else {
					sl = make([]string, 0, 1)
					attr[cl.Attr] = sl
				}
				attr[cl.Attr] = append(sl, cl.Value)
			}
		}
	}

	// If the content permanode is now known, look up its type
	if content, ok := jm["content"].(string); ok && content != "" {
		cbr := blobref.Parse(content)

		dm := jm
		if dmap != nil {
			dm = dmap(cbr)
		}

		mime, size, err := sh.index.GetBlobMimeType(cbr)
		if err == nil {
			setMimeType(dm, mime)
			dm["size"] = size
		}
	}
}

const camliTypePrefix = "application/json; camliType="

func setMimeType(m map[string]interface{}, mime string) {
	m["type"] = mime
	if strings.HasPrefix(mime, camliTypePrefix) {
		m["camliType"] = mime[len(camliTypePrefix):]
	}
}
