package variable

import "errors"

type Str struct {
	Name    string
	Content string
}

type Integer struct {
	Name    string
	Content int
}

type Float1 struct {
	Name    string
	Content float32
}

type Float2 struct {
	Name    string
	Content float64
}

type Boolean struct {
	Name    string
	Content bool
}

func FindStrByName(arr []Str, targetName string) (Str, error) {
	for _, obj := range arr {
		if obj.Name == targetName {
			return obj, nil
		}
	}
	return Str{}, errors.New("variableNotFound")
}

func FindIntegerByName(arr []Integer, targetName string) (Integer, error) {
	for _, obj := range arr {
		if obj.Name == targetName {
			return obj, nil
		}
	}
	return Integer{}, errors.New("variableNotFound")
}

func FindFloat1ByName(arr []Float1, targetName string) (Float1, error) {
	for _, obj := range arr {
		if obj.Name == targetName {
			return obj, nil
		}
	}
	return Float1{}, errors.New("variableNotFound")
}

func FindFloat2ByName(arr []Float2, targetName string) (Float2, error) {
	for _, obj := range arr {
		if obj.Name == targetName {
			return obj, nil
		}
	}
	return Float2{}, errors.New("variableNotFound")
}

func FindBooleanByName(arr []Boolean, targetName string) (Boolean, error) {
	for _, obj := range arr {
		if obj.Name == targetName {
			return obj, nil
		}
	}
	return Boolean{}, errors.New("variableNotFound")
}
