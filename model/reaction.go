// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type Reaction struct {
	UserId    string `json:"user_id"`
	PostId    string `json:"post_id"`
	EmojiName string `json:"emoji_name"`
	CreateAt  int64  `json:"create_at"`
}

func (o *Reaction) ToJson() string {
	if b, err := json.Marshal(o); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ReactionFromJson(data io.Reader) *Reaction {
	var o Reaction

	if err := json.NewDecoder(data).Decode(&o); err != nil {
		return nil
	} else {
		return &o
	}
}

func ReactionsToJson(o []*Reaction) string {
	if b, err := json.Marshal(o); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ReactionsFromJson(data io.Reader) []*Reaction {
	var o []*Reaction

	if err := json.NewDecoder(data).Decode(&o); err != nil {
		return nil
	} else {
		return o
	}
}

func (o *Reaction) IsValid() *AppError {
	if len(o.UserId) != 26 {
		return NewLocAppError("Reaction.IsValid", "model.reaction.is_valid.user_id.app_error", nil, "user_id="+o.UserId)
	}

	if len(o.PostId) != 26 {
		return NewLocAppError("Reaction.IsValid", "model.reaction.is_valid.post_id.app_error", nil, "post_id="+o.PostId)
	}

	if len(o.EmojiName) == 0 || len(o.EmojiName) > 64 {
		return NewLocAppError("Reaction.IsValid", "model.reaction.is_valid.emoji_name.app_error", nil, "emoji_name="+o.EmojiName)
	}

	if o.CreateAt == 0 {
		return NewLocAppError("Reaction.IsValid", "model.reaction.is_valid.create_at.app_error", nil, "")
	}

	return nil
}

func (o *Reaction) PreSave() {
	if o.CreateAt == 0 {
		o.CreateAt = GetMillis()
	}
}
