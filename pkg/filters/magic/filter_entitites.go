package magic

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tele "gopkg.in/telebot.v3"
)

type EntitiesMagicFilter struct {
	SliceMagicFilter[tele.Entities, tele.MessageEntity]
}

func newEntitiesFilter(getter ItemGetter[tele.Entities]) (f EntitiesMagicFilter) {
	f.getter = getter
	return
}

func (m EntitiesMagicFilter) Have(entityType tele.EntityType) tf.Filter {
	return m.AtLeastOne(func(entity tele.MessageEntity) bool {
		return entity.Type == entityType
	})
}
