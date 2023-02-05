package dictionary

import (
	"atom/container"
	"atom/database/query"
	"context"
	"errors"
	"log"
	"time"
)

type Dict struct {
	dict      map[uint64]*DictInfo
	dictItems []*DictInfo
	mapAlias  map[string]uint64
	query     *query.Query
}

func init() {
	if err := container.Container.Provide(NewDictionary); err != nil {
		log.Fatal(err)
	}
}

func NewDictionary(query *query.Query) (*Dict, error) {
	dict := &Dict{
		query:     query,
		dict:      make(map[uint64]*DictInfo),
		dictItems: []*DictInfo{},
		mapAlias:  make(map[string]uint64),
	}

	if err := dict.Load(); err != nil {
		return nil, err
	}
	return dict, nil
}

func (dict *Dict) Load() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	dictTable := dict.query.SysDictionary
	items, err := dictTable.WithContext(ctx).Where(dictTable.Status.Is(true)).Find()
	if err != nil {
		return err
	}

	ids := []uint64{}
	for _, item := range items {
		ids = append(ids, item.ID)
		dict.mapAlias[item.Alias_] = item.ID
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	dictDetailTable := dict.query.SysDictionaryDetail
	dictItems, err := dictDetailTable.WithContext(ctx).
		Where(dictDetailTable.Status.Is(true)).
		Where(dictDetailTable.ID.In(ids...)).
		Order(dictDetailTable.Weight.Desc()).
		Find()
	if err != nil {
		return err
	}

	idItems := make(map[uint64][]*DictItem)
	for _, dictItem := range dictItems {
		id := uint64(dictItem.SysDictionaryID)
		idItems[id] = append(idItems[id], &DictItem{
			Label: dictItem.Label,
			Value: dictItem.Value,
		})
	}
	for _, item := range items {
		info := &DictInfo{
			ID:          item.ID,
			Name:        item.Name,
			Alias:       item.Alias_,
			Description: item.Description,
			Items:       idItems[item.ID],
		}
		dict.dictItems = append(dict.dictItems, info)
		dict.dict[item.ID] = info
	}

	return nil
}

// GetLabelByValue
func (dict *Dict) GetLabelByValue(alias, value string) (string, error) {
	items, err := dict.GetItems(alias)
	if err != nil {
		return "", err
	}

	for _, item := range items {
		if item.Value == value {
			return item.Label, nil
		}
	}

	return "", errors.New("dict item not exists")
}

// GetLabelByValue
func (dict *Dict) GetItems(alias string) ([]*DictItem, error) {
	dictID, ok := dict.mapAlias[alias]
	if !ok {
		return nil, errors.New("dict not exists")
	}

	dictItem, ok := dict.dict[dictID]
	if !ok {
		return nil, errors.New("dict not exists")
	}
	return dictItem.Items, nil
}

func (dict *Dict) All(alias string) []*DictInfo {
	return dict.dictItems
}
