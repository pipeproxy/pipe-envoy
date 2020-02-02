package clean

import (
	"encoding/json"
	"sort"
)

func Clean(config []byte) ([]byte, error) {
	var f struct {
		Pipe       json.RawMessage
		Init       json.RawMessage
		Components []json.RawMessage
	}
	err := json.Unmarshal(config, &f)
	if err != nil {
		return nil, err
	}

	save := map[string]struct{}{}

	if len(f.Pipe) != 0 {
		err = getNeed(save, f.Pipe)
		if err != nil {
			return nil, err
		}
	}

	if len(f.Init) != 0 {
		err = getNeed(save, f.Init)
		if err != nil {
			return nil, err
		}
	}

	cn := map[string]json.RawMessage{}
	components := f.Components
	swap := []json.RawMessage{}
	for {
		continued := false
		for _, component := range components {
			var f struct {
				Name string `json:"@Name"`
			}
			err := json.Unmarshal(component, &f)
			if err != nil {
				return nil, err
			}

			_, ok := save[f.Name]
			if ok && f.Name != "" && cn[f.Name] == nil {
				cn[f.Name] = component
				err = getNeed(save, component)
				if err != nil {
					return nil, err
				}
				continued = true
			} else {
				swap = append(swap, component)
			}
		}

		components = swap
		swap = swap[:0]
		if !continued {
			break
		}
	}

	f.Components = f.Components[:0]
	for _, c := range cn {
		f.Components = append(f.Components, c)
	}

	sort.Slice(f.Components, func(i, j int) bool {
		return lessBytes(f.Components[i], f.Components[j])
	})
	return json.Marshal(f)
}

func lessBytes(i, j []byte) bool {
	for k := 0; k != len(i) && k != len(j); k++ {
		if i[k] != j[k] {
			return i[k] < j[k]
		}
	}
	return len(i) < len(j)
}

func getNeed(save map[string]struct{}, config []byte) error {
	d := newDependencies()
	err := d.Decode(config)
	if err != nil {
		return err
	}
	for _, node := range d.Nodes() {
		for _, ref := range node.Refs {
			save[ref] = struct{}{}
		}
	}
	return nil
}
