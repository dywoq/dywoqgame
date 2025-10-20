package entry

import (
	"github.com/dywoq/dywoqgame/resource"
)

// ResourceManager is a standard resource manager,
// implementing interface resource.Management.
type ResourceManager struct {
	res map[string]resource.Resource
}

// Exists checks if the resource named name exists.
func (r *ResourceManager) Exists(name string) bool {
	_, ok := r.res[name]
	return ok
}

// TableOf returns the metadata table of the resource named name.
// If there's no called resource, it returns resource.ErrNotFound.
func (r *ResourceManager) TableOf(name string) (resource.Table, error) {
	res, ok := r.res[name]
	if !ok {
		return nil, resource.ErrNotFound(name)
	}
	return res.Table(), nil
}

// Load loads the resource into the resource manager.
// Returns an error resource.ErrExists if there's already called resource.
//
// Returns an error if the metadata table of the resource fails to comply
// with contracts.
func (r *ResourceManager) Load(res resource.Resource) error {
	if r.Exists(res.Name()) {
		return resource.ErrExists(res.Name())
	}
	err := contractCheck(res)
	if err != nil {
		return err
	}
	r.res[res.Name()] = res
	return nil
}

// Unload unloads the resource from the resource manager,
// calling resource.Resource.Free method.
//
// Returns resource.ErrNotFound if there's no resource with the name.
func (r *ResourceManager) Unload(res resource.Resource) error {
	got, ok := r.res[res.Name()]
	if !ok {
		return resource.ErrNotFound(res.Name())
	}
	err := got.Free()
	if err != nil {
		return err
	}
	delete(r.res, res.Name())
	return nil
}

// KindOf returns the kind of the resource.
// Return resource.Unknown adn resource.ErrNotFound
// if there's no resource called name.
func (r *ResourceManager) KindOf(name string) (resource.Kind, error) {
	got, ok := r.res[name]
	if !ok {
		return resource.Unknown, resource.ErrNotFound(name)
	}
	return got.Kind(), nil
}
