package boMemeCoin

import "meme_coin_api/service/internal/model"

type CreateArgs struct {
	Name        string
	Description string
}

type CreateReply struct {
	ID int64 `json:"id,string"`
}

type GetArgs struct {
	ID int64
}

type GetReply struct {
	*model.MemeCoin
}

type UpdateArgs struct {
	ID          int64
	Description string
}

type IncreasePopularityScoreArgs struct {
	ID int64
}

type DeleteArgs struct {
	ID int64
}
