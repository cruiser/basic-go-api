package models

// basically making an interface that makes all models accessible from ths function without having to import them individually
func GetModels() []interface{} {
	return []interface{}{
		&Users{},
		&Client{},
		&Choreography{},
	}
}
