// Package mount provides a Datastore that has other Datastores
// mounted at various key prefixes.
package mount

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"gx/ipfs/QmdHG8MAuARdGHxx4rPQASLcvhz24fzjSQq7AJRAQEorq5/go-datastore"
	"gx/ipfs/QmdHG8MAuARdGHxx4rPQASLcvhz24fzjSQq7AJRAQEorq5/go-datastore/query"
)

var (
	ErrNoMount = errors.New("no datastore mounted for this key")
)

type Mount struct {
	Prefix    datastore.Key
	Datastore datastore.Datastore
}

func New(mounts []Mount) *Datastore {
	// make a copy so we're sure it doesn't mutate
	m := make([]Mount, len(mounts))
	for i, v := range mounts {
		m[i] = v
	}
	sort.Slice(m, func(i, j int) bool { return m[i].Prefix.String() > m[j].Prefix.String() })
	return &Datastore{mounts: m}
}

type Datastore struct {
	mounts []Mount
}

var _ datastore.Datastore = (*Datastore)(nil)

func (d *Datastore) lookup(key datastore.Key) (ds datastore.Datastore, mountpoint, rest datastore.Key) {
	for _, m := range d.mounts {
		if m.Prefix.Equal(key) || m.Prefix.IsAncestorOf(key) {
			s := strings.TrimPrefix(key.String(), m.Prefix.String())
			k := datastore.NewKey(s)
			return m.Datastore, m.Prefix, k
		}
	}
	return nil, datastore.NewKey("/"), key
}

// lookupAll returns all mounts that might contain keys that are descendant of <key>
//
// Matching: /ao/e
//
// /          B /ao/e
// /a/        not matching
// /ao/       B /e
// /ao/e/     A /
// /ao/e/uh/  A /
// /aoe/      not matching
func (d *Datastore) lookupAll(key datastore.Key) (ds []datastore.Datastore, mountpoint, rest []datastore.Key) {
	for _, m := range d.mounts {
		p := m.Prefix.String()
		if len(p) > 1 {
			p = p + "/"
		}

		if strings.HasPrefix(p, key.String()) {
			ds = append(ds, m.Datastore)
			mountpoint = append(mountpoint, m.Prefix)
			rest = append(rest, datastore.NewKey("/"))
		} else if strings.HasPrefix(key.String(), p) {
			r := strings.TrimPrefix(key.String(), m.Prefix.String())

			ds = append(ds, m.Datastore)
			mountpoint = append(mountpoint, m.Prefix)
			rest = append(rest, datastore.NewKey(r))
		}
	}
	return ds, mountpoint, rest
}

func (d *Datastore) Put(key datastore.Key, value interface{}) error {
	ds, _, k := d.lookup(key)
	if ds == nil {
		return ErrNoMount
	}
	return ds.Put(k, value)
}

func (d *Datastore) Get(key datastore.Key) (value interface{}, err error) {
	ds, _, k := d.lookup(key)
	if ds == nil {
		return nil, datastore.ErrNotFound
	}
	return ds.Get(k)
}

func (d *Datastore) Has(key datastore.Key) (exists bool, err error) {
	ds, _, k := d.lookup(key)
	if ds == nil {
		return false, nil
	}
	return ds.Has(k)
}

func (d *Datastore) Delete(key datastore.Key) error {
	ds, _, k := d.lookup(key)
	if ds == nil {
		return datastore.ErrNotFound
	}
	return ds.Delete(k)
}

func (d *Datastore) Query(q query.Query) (query.Results, error) {
	if len(q.Filters) > 0 ||
		len(q.Orders) > 0 ||
		q.Limit > 0 ||
		q.Offset > 0 {
		// TODO this is still overly simplistic, but the only callers are
		// `ipfs refs local` and ipfs-ds-convert.
		return nil, errors.New("mount only supports listing all prefixed keys in random order")
	}
	prefix := datastore.NewKey(q.Prefix)
	dses, mounts, rests := d.lookupAll(prefix)

	// current itorator state
	var res query.Results
	var mount datastore.Key
	i := 0

	return query.ResultsFromIterator(q, query.Iterator{
		Next: func() (query.Result, bool) {
			var r query.Result
			var more bool

			for try := true; try; try = len(dses) > i {
				if res == nil {
					if len(dses) <= i {
						//This should not happen normally
						return query.Result{}, false
					}

					dst := dses[i]
					mount = mounts[i]
					rest := rests[i]

					q2 := q
					q2.Prefix = rest.String()
					r, err := dst.Query(q2)
					if err != nil {
						return query.Result{Error: err}, false
					}
					res = r
				}

				r, more = res.NextSync()
				if !more {
					err := res.Close()
					if err != nil {
						return query.Result{Error: err}, false
					}
					res = nil

					i++
					more = len(dses) > i
				} else {
					break
				}
			}

			r.Key = mount.Child(datastore.RawKey(r.Key)).String()
			return r, more
		},
		Close: func() error {
			if len(mounts) > i && res != nil {
				return res.Close()
			}
			return nil
		},
	}), nil
}

func (d *Datastore) Close() error {
	for _, d := range d.mounts {
		if c, ok := d.Datastore.(io.Closer); ok {
			err := c.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type mountBatch struct {
	mounts map[string]datastore.Batch

	d *Datastore
}

func (d *Datastore) Batch() (datastore.Batch, error) {
	return &mountBatch{
		mounts: make(map[string]datastore.Batch),
		d:      d,
	}, nil
}

func (mt *mountBatch) lookupBatch(key datastore.Key) (datastore.Batch, datastore.Key, error) {
	child, loc, rest := mt.d.lookup(key)
	t, ok := mt.mounts[loc.String()]
	if !ok {
		bds, ok := child.(datastore.Batching)
		if !ok {
			return nil, datastore.NewKey(""), datastore.ErrBatchUnsupported
		}
		var err error
		t, err = bds.Batch()
		if err != nil {
			return nil, datastore.NewKey(""), err
		}
		mt.mounts[loc.String()] = t
	}
	return t, rest, nil
}

func (mt *mountBatch) Put(key datastore.Key, val interface{}) error {
	t, rest, err := mt.lookupBatch(key)
	if err != nil {
		return err
	}

	return t.Put(rest, val)
}

func (mt *mountBatch) Delete(key datastore.Key) error {
	t, rest, err := mt.lookupBatch(key)
	if err != nil {
		return err
	}

	return t.Delete(rest)
}

func (mt *mountBatch) Commit() error {
	for _, t := range mt.mounts {
		err := t.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Datastore) Check() error {
	for _, m := range d.mounts {
		if c, ok := m.Datastore.(datastore.CheckedDatastore); ok {
			if err := c.Check(); err != nil {
				return fmt.Errorf("checking datastore at %s: %s", m.Prefix.String(), err.Error())
			}
		}
	}
	return nil
}

func (d *Datastore) Scrub() error {
	for _, m := range d.mounts {
		if c, ok := m.Datastore.(datastore.ScrubbedDatastore); ok {
			if err := c.Scrub(); err != nil {
				return fmt.Errorf("scrubbing datastore at %s: %s", m.Prefix.String(), err.Error())
			}
		}
	}
	return nil
}

func (d *Datastore) CollectGarbage() error {
	for _, m := range d.mounts {
		if c, ok := m.Datastore.(datastore.GCDatastore); ok {
			if err := c.CollectGarbage(); err != nil {
				return fmt.Errorf("gc on datastore at %s: %s", m.Prefix.String(), err.Error())
			}
		}
	}
	return nil
}
