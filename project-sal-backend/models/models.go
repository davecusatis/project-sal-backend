package models

import jwt "github.com/dgrijalva/jwt-go"

type PostData struct {
	ContentType string   `json:"content_type"`
	Message     string   `json:"message"`
	Targets     []string `json:"targets"`
}

type PubsubMessage struct {
	Token       *TokenData
	MessageType string      `json:"type"`
	Data        MessageData `json:"data"`
}

type TokenData struct {
	Token        string
	UserID       string
	ChannelID    string
	Role         string
	OpaqueUserID string
	PubsubPerms  PubsubPerms `json:"pubsub_perms"`
}

type PubsubPerms struct {
	Send   []string `json:"send"`
	Listen []string `json:"listen"`
}
type TokenClaims struct {
	OpaqueUserID string      `json:"opaque_user_id"`
	UserID       string      `json:"user_id"`
	ChannelID    string      `json:"channel_id"`
	Role         string      `json:"role"`
	PubsubPerms  PubsubPerms `json:"pubsub_perms"`
	jwt.StandardClaims
}

type MessageData struct {
	Score int `json:"score"`
}

type TwitchUserData struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
}
type TwitchUserResponse struct {
	Data []TwitchUserData `json:"data"`
}
