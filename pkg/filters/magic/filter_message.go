package magic

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tele "gopkg.in/telebot.v3"
)

type MessageFilter struct {
	getter ItemGetter[*tele.Message]
}

func newMessageFilter(getter ItemGetter[*tele.Message]) MessageFilter {
	return MessageFilter{getter: getter}
}

func (m MessageFilter) ID() NumberFilter[int] {
	return newNumberFilter(joinGetter(m.getter, func(m *tele.Message) int {
		return m.ID
	}))
}
func (m MessageFilter) ThreadID() NumberFilter[int] {
	return newNumberFilter(joinGetter(m.getter, func(m *tele.Message) int {
		return m.ThreadID
	}))
}
func (m MessageFilter) Sender() UserMagicFilter {
	return newUserMagicFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.User {
		return m.Sender
	}))
}
func (m MessageFilter) Unixtime() NumberFilter[int64] {
	return newNumberFilter(joinGetter(m.getter, func(m *tele.Message) int64 {
		return m.Unixtime
	}))
}
func (m MessageFilter) Chat() ChatMagicFilter {
	return newChatFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.Chat {
		return m.Chat
	}))
}
func (m MessageFilter) SenderChat() ChatMagicFilter {
	return newChatFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.Chat {
		return m.SenderChat
	}))
}
func (m MessageFilter) OriginalSender() UserMagicFilter {
	return newUserMagicFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.User {
		return m.OriginalSender
	}))
}

func (m MessageFilter) OriginalChat() ChatMagicFilter {
	return newChatFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.Chat {
		return m.OriginalChat
	}))
}

func (m MessageFilter) OriginalMessageID() NumberFilter[int] {
	return newNumberFilter(joinGetter(m.getter, func(m *tele.Message) int {
		return m.OriginalMessageID
	}))
}
func (m MessageFilter) OriginalSignature() *StringPipeline {
	return newStringPipeline(joinGetter(m.getter, func(m *tele.Message) string {
		return m.OriginalSignature
	}))
}

func (m MessageFilter) OriginalSenderName() *StringPipeline {
	return newStringPipeline(joinGetter(m.getter, func(m *tele.Message) string {
		return m.OriginalSenderName
	}))
}
func (m MessageFilter) OriginalUnixtime() NumberFilter[int] {
	return newNumberFilter(joinGetter(m.getter, func(m *tele.Message) int {
		return m.OriginalUnixtime
	}))
}
func (m MessageFilter) AutomaticForward(want bool) tf.Filter {
	return newPredicate(joinGetter(m.getter, func(m *tele.Message) bool {
		return m.AutomaticForward
	}), boolFilter(want))
}

func (m MessageFilter) ReplyTo() MessageFilter {
	return newMessageFilter(func(ctx tele.Context) (*tele.Message, bool) {
		m, ok := m.getter(ctx)
		if !ok {
			return nil, false
		}
		return m.ReplyTo, m.ReplyTo != nil
	})
}
func (m MessageFilter) Via() UserMagicFilter {
	return newUserMagicFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.User {
		return m.Via
	}))
}
func (m MessageFilter) LastEdit() NumberFilter[int64] {
	return newNumberFilter(joinGetter(m.getter, func(m *tele.Message) int64 {
		return m.LastEdit
	}))
}

func (m MessageFilter) TopicMessage(want bool) tf.Filter {
	return newBoolFilter(joinGetter(m.getter, func(m *tele.Message) bool {
		return m.TopicMessage
	}), want)
}
func (m MessageFilter) Protected(want bool) tf.Filter {
	return newBoolFilter(joinGetter(m.getter, func(m *tele.Message) bool {
		return m.Protected
	}), want)
}

func (m MessageFilter) AlbumID() *StringPipeline {
	return newStringPipeline(joinGetter(m.getter, func(m *tele.Message) string {
		return m.AlbumID
	}))
}
func (m MessageFilter) Signature() *StringPipeline {
	return newStringPipeline(joinGetter(m.getter, func(m *tele.Message) string {
		return m.Signature
	}))
}
func (m MessageFilter) Text() *StringPipeline {
	return newStringPipeline(joinGetter(m.getter, func(m *tele.Message) string {
		return m.Text
	}))
}

func (m MessageFilter) Payload() *StringPipeline {
	return newStringPipeline(joinGetter(m.getter, func(m *tele.Message) string {
		return m.Payload
	}))
}

func (m MessageFilter) Entities() EntitiesMagicFilter {
	return newEntitiesFilter(joinGetter(m.getter, func(m *tele.Message) tele.Entities {
		return m.Entities
	}))
}
func (m MessageFilter) Caption() *StringPipeline {
	return newStringPipeline(joinGetter(m.getter, func(m *tele.Message) string {
		return m.Caption
	}))
}

func (m MessageFilter) CaptionEntities() EntitiesMagicFilter {
	return newEntitiesFilter(joinGetter(m.getter, func(m *tele.Message) tele.Entities {
		return m.CaptionEntities
	}))
}
func (m MessageFilter) Audio(f ItemFilter[*tele.Audio]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Audio {
		return m.Audio
	}), f)
}
func (m MessageFilter) Document(f ItemFilter[*tele.Document]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Document {
		return m.Document
	}), f)
}
func (m MessageFilter) Photo(f ItemFilter[*tele.Photo]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Photo {
		return m.Photo
	}), f)
}
func (m MessageFilter) Sticker(f ItemFilter[*tele.Sticker]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Sticker {
		return m.Sticker
	}), f)
}
func (m MessageFilter) Voice(f ItemFilter[*tele.Voice]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Voice {
		return m.Voice
	}), f)
}
func (m MessageFilter) VideoNote(f ItemFilter[*tele.VideoNote]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.VideoNote {
		return m.VideoNote
	}), f)
}
func (m MessageFilter) Video(f ItemFilter[*tele.Video]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Video {
		return m.Video
	}), f)
}
func (m MessageFilter) Animation(f ItemFilter[*tele.Animation]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Animation {
		return m.Animation
	}), f)
}
func (m MessageFilter) Contact(f ItemFilter[*tele.Contact]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Contact {
		return m.Contact
	}), f)
}
func (m MessageFilter) Location(f ItemFilter[*tele.Location]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Location {
		return m.Location
	}), f)
}
func (m MessageFilter) Venue(f ItemFilter[*tele.Venue]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Venue {
		return m.Venue
	}), f)
}
func (m MessageFilter) Poll(f ItemFilter[*tele.Poll]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Poll {
		return m.Poll
	}), f)
}
func (m MessageFilter) Game(f ItemFilter[*tele.Game]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Game {
		return m.Game
	}), f)
}
func (m MessageFilter) Dice(f ItemFilter[*tele.Dice]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Dice {
		return m.Dice
	}), f)
}

func (m MessageFilter) UserJoined() UserMagicFilter {
	return newUserMagicFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.User {
		return m.UserJoined
	}))
}

func (m MessageFilter) UserLeft() UserMagicFilter {
	return newUserMagicFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.User {
		return m.UserLeft
	}))
}

func (m MessageFilter) NewGroupTitle() *StringPipeline {
	return newStringPipeline(joinGetter(m.getter, func(m *tele.Message) string {
		return m.NewGroupTitle
	}))
}

func (m MessageFilter) NewGroupPhoto(f ItemFilter[*tele.Photo]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Photo {
		return m.NewGroupPhoto
	}), f)
}

func (m MessageFilter) UsersJoined() SliceMagicFilter[[]tele.User, tele.User] {
	return newArrayFilter(joinGetter(m.getter, func(m *tele.Message) []tele.User {
		return m.UsersJoined
	}))
}

func (m MessageFilter) GroupPhotoDeleted(want bool) tf.Filter {
	return newBoolFilter(joinGetter(m.getter, func(m *tele.Message) bool {
		return m.GroupPhotoDeleted
	}), want)
}

func (m MessageFilter) GroupCreated(want bool) tf.Filter {
	return newBoolFilter(joinGetter(m.getter, func(m *tele.Message) bool {
		return m.GroupCreated
	}), want)
}

func (m MessageFilter) SuperGroupCreated(want bool) tf.Filter {
	return newBoolFilter(joinGetter(m.getter, func(m *tele.Message) bool {
		return m.SuperGroupCreated
	}), want)
}

func (m MessageFilter) ChannelCreated(want bool) tf.Filter {
	return newBoolFilter(joinGetter(m.getter, func(m *tele.Message) bool {
		return m.ChannelCreated
	}), want)
}

func (m MessageFilter) MigrateTo() NumberFilter[int64] {
	return newNumberFilter(joinGetter(m.getter, func(m *tele.Message) int64 {
		return m.MigrateTo
	}))
}

func (m MessageFilter) MigrateFrom() NumberFilter[int64] {
	return newNumberFilter(joinGetter(m.getter, func(m *tele.Message) int64 {
		return m.MigrateFrom
	}))
}

func (m MessageFilter) PinnedMessage() MessageFilter {
	return newMessageFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.Message {
		return m.PinnedMessage
	}))
}
func (m MessageFilter) Invoice(f ItemFilter[*tele.Invoice]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Invoice {
		return m.Invoice
	}), f)
}
func (m MessageFilter) Payment(f ItemFilter[*tele.Payment]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.Payment {
		return m.Payment
	}), f)
}
func (m MessageFilter) UserShared(f ItemFilter[*tele.RecipientShared]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.RecipientShared {
		return m.UserShared
	}), f)
}
func (m MessageFilter) ChatShared(f ItemFilter[*tele.RecipientShared]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.RecipientShared {
		return m.ChatShared
	}), f)
}
func (m MessageFilter) ConnectedWebsite() *StringPipeline {
	return newStringPipeline(joinGetter(m.getter, func(m *tele.Message) string {
		return m.ConnectedWebsite
	}))
}
func (m MessageFilter) VideoChatStarted(f ItemFilter[*tele.VideoChatStarted]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.VideoChatStarted {
		return m.VideoChatStarted
	}), f)
}
func (m MessageFilter) VideoChatEnded(f ItemFilter[*tele.VideoChatEnded]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.VideoChatEnded {
		return m.VideoChatEnded
	}), f)
}
func (m MessageFilter) VideoChatParticipants(f ItemFilter[*tele.VideoChatParticipants]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.VideoChatParticipants {
		return m.VideoChatParticipants
	}), f)
}
func (m MessageFilter) VideoChatScheduled(f ItemFilter[*tele.VideoChatScheduled]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.VideoChatScheduled {
		return m.VideoChatScheduled
	}), f)
}
func (m MessageFilter) WebAppData(f ItemFilter[*tele.WebAppData]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.WebAppData {
		return m.WebAppData
	}), f)
}

func (m MessageFilter) ProximityAlert(f ItemFilter[*tele.ProximityAlert]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.ProximityAlert {
		return m.ProximityAlert
	}), f)
}

func (m MessageFilter) AutoDeleteTimer(f ItemFilter[*tele.AutoDeleteTimer]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.AutoDeleteTimer {
		return m.AutoDeleteTimer
	}), f)
}

func (m MessageFilter) TopicCreated() TopicFilter {
	return newTopicFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.Topic {
		return m.TopicCreated
	}))
}

func (m MessageFilter) TopicReopened() TopicFilter {
	return newTopicFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.Topic {
		return m.TopicReopened
	}))
}
func (m MessageFilter) TopicEdited() TopicFilter {
	return newTopicFilter(joinNotNil(m.getter, func(m *tele.Message) *tele.Topic {
		return m.TopicEdited
	}))
}

func (m MessageFilter) WriteAccessAllowed(f ItemFilter[*tele.WriteAccessAllowed]) tf.Filter {
	return newPredicate(joinNotNil(m.getter, func(m *tele.Message) *tele.WriteAccessAllowed {
		return m.WriteAccessAllowed
	}), f)
}

func (m MessageFilter) HasMediaSpoiler(want bool) tf.Filter {
	return newBoolFilter(joinGetter(m.getter, func(m *tele.Message) bool {
		return m.HasMediaSpoiler
	}), want)
}

func (m MessageFilter) On(f ItemFilter[*tele.Message]) tf.Filter {
	return newPredicate(m.getter, f)
}

func (m MessageFilter) All(factories ...FilterFactory[MessageFilter]) tf.Filter {
	return logicBranch(m, And, factories)
}

func (m MessageFilter) Any(factories ...FilterFactory[MessageFilter]) tf.Filter {
	return logicBranch(m, Or, factories)
}
