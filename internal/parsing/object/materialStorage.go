package objectParsing

import "github.com/tmazitov/42_scop/internal/rende"

type materialStorage struct {
	values   []*rende.Material
	filePath string
}

func newMaterialStorage(filePath string) *materialStorage {
	return &materialStorage{filePath: filePath}
}

func (m *materialStorage) Add(material ...*rende.Material) {
	m.values = append(m.values, material...)
}

func (m *materialStorage) Find(name string) *rende.Material {
	for _, m := range m.values {
		if m.Name() == name {
			return m
		}
	}
	return nil
}
