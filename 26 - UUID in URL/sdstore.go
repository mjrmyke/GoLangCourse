package mem

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func storeDstore(m model, req *http.Request) error {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Photos", m.ID, 0, nil)

	_, err := datastore.Put(ctx, key, &m)
	if err != nil {
		log.Errorf(ctx, "error storing datastore due to datastore.put: %s", err)
		return err
	}
	return nil
}

func retrieveDstore(id string, req *http.Request) (model, error) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Photos", id, 0, nil)

	var m model
	err := datastore.Get(ctx, key, &m)
	if err != nil {
		log.Errorf(ctx, "error retrieving data due to datastore.Get: %s", err)
		return m, err
	}
	return m, nil
}
