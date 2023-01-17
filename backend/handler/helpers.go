// Package handler - Helper functions
package handler

func checkRedis() bool {
	// TODO: Make key unique to evade collisions between instances
	// TODO: Make value unique to evade false positives
	key := "healthz"
	value := "foo"

	_, err := CRUD.Update(key, value)
	if err != nil {
		return false
	}
	resultUpdate, err := CRUD.Read(key)
	if err != nil {
		return false
	}

	resultDelete, err := CRUD.Delete(key)
	if err != nil {
		return false
	}

	return resultUpdate == value && resultDelete
}

func all(values []bool) bool {
	for _, v := range values {
		if v != true {
			return false
		}
	}
	return true
}
