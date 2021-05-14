package ssc

// Context used as temporary store for data

var context = map[Page]map[string]interface{}{}

func SetContext(p Page, key string, value interface{}) {
	space, ok := context[p]
	if !ok {
		space = map[string]interface{}{}
	}
	space[key] = value
	context[p] = space
}

func GetContext(p Page, key string) interface{} {
	return context[p][key]
}
