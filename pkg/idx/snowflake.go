package idx

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

// NOTICE: Clock Rollback Problem

// SnowFlakeID snowflake id
type SnowFlakeID struct {
	*snowflake.Node
}

var snowFlakeIDGenerator *SnowFlakeID

func init() {
	rand.New(rand.NewSource(time.Now().Unix()))
	node, err := snowflake.NewNode(int64(rand.Intn(1000)) + 1)
	if err != nil {
		panic(err.Error())
	}
	snowFlakeIDGenerator = &SnowFlakeID{node}
}

func ID() snowflake.ID {
	id := snowFlakeIDGenerator.Generate()
	return id
}

func IDStr12() string {
	id := snowFlakeIDGenerator.Generate()
	return id.Base58()
}

func IDStr() string {
	id := snowFlakeIDGenerator.Generate()
	return id.Base32()
}
