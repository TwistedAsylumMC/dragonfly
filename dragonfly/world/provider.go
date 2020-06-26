package world

import (
	"github.com/df-mc/dragonfly/dragonfly/world/chunk"
	"github.com/df-mc/dragonfly/dragonfly/world/difficulty"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	"io"
)

// Provider represents a value that may provide world data to a World value. It usually does the reading and
// writing of the world data so that the World may use it.
type Provider interface {
	io.Closer
	// WorldName returns the name of the world that the provider provides for. When setting the provider of a
	// World, the World will replace its current name with this one.
	WorldName() string
	// SetWorldName sets the name of the world to a new name.
	SetWorldName(name string)
	// WorldSpawn returns the spawn position of the world. Although players may spawn at different positions,
	// every new player spawns at this position.
	WorldSpawn() BlockPos
	// SetWorldSpawn sets the spawn of a world to a new position.
	SetWorldSpawn(pos BlockPos)
	// LoadChunk attempts to load a chunk from the chunk position passed. If successful, a non-nil chunk is
	// returned and exists is true and err nil. If no chunk was saved at the chunk position passed, the chunk
	// returned is nil, and so is the error. If the chunk did exist, but if the data was invalid, nil is
	// returned for the chunk and true, with a non-nil error.
	// If exists ends up false, the chunk at the position is instead newly generated by the world.
	LoadChunk(position ChunkPos) (c *chunk.Chunk, exists bool, err error)
	// SaveChunk saves a chunk at a specific position in the provider. If writing was not successful, an error
	// is returned.
	SaveChunk(position ChunkPos, c *chunk.Chunk) error
	// LoadEntities loads all entities stored at a particular chunk position. If the entities cannot be read,
	// LoadEntities returns a non-nil error.
	LoadEntities(position ChunkPos) ([]Entity, error)
	// SaveEntities saves a list of entities in a chunk position. If writing is not successful, an error is
	// returned.
	SaveEntities(position ChunkPos, entities []Entity) error
	// LoadBlockNBT loads the block NBT, also known as block entities, at a specific chunk position. If the
	// NBT cannot be read, LoadBlockNBT returns a non-nil error.
	LoadBlockNBT(position ChunkPos) ([]map[string]interface{}, error)
	// SaveBlockNBT saves block NBT, or block entities, to a specific chunk position. If the NBT cannot be
	// stored, SaveBlockNBT returns a non-nil error.
	SaveBlockNBT(position ChunkPos, data map[[3]int]map[string]interface{}) error
	// LoadTime loads the time of the world.
	LoadTime() int64
	// SaveTime saves the time of the world.
	SaveTime(time int64)
	// SaveTimeCycle saves the state of the time cycle: Either stopped or started. If true is passed, the time
	// is running. If false, the time is stopped.
	SaveTimeCycle(running bool)
	// LoadTimeCycle loads the state of the time cycle: If time is running, true is returned. If the time
	// cycle is stopped, false is returned.
	LoadTimeCycle() bool
	// LoadDefaultGameMode loads the default game mode of the world.
	LoadDefaultGameMode() gamemode.GameMode
	// SaveDefaultGameMode sets the default game mode of the world.
	SaveDefaultGameMode(mode gamemode.GameMode)
	// LoadDifficulty loads the difficulty of a world.
	LoadDifficulty() difficulty.Difficulty
	// SaveDifficulty saves the difficulty of a world.
	SaveDifficulty(d difficulty.Difficulty)
}

// NoIOProvider implements a Provider while not performing any disk I/O. It generates values on the run and
// dynamically, instead of reading and writing data, and returns otherwise empty values.
type NoIOProvider struct{}

// LoadDifficulty ...
func (NoIOProvider) LoadDifficulty() difficulty.Difficulty { return difficulty.Normal{} }

// SaveDifficulty ...
func (NoIOProvider) SaveDifficulty(difficulty.Difficulty) {}

// LoadDefaultGameMode ...
func (NoIOProvider) LoadDefaultGameMode() gamemode.GameMode { return gamemode.Adventure{} }

// SaveDefaultGameMode ...
func (NoIOProvider) SaveDefaultGameMode(gamemode.GameMode) {}

// SetWorldSpawn ...
func (NoIOProvider) SetWorldSpawn(BlockPos) {}

// SaveTimeCycle ...
func (NoIOProvider) SaveTimeCycle(bool) {}

// LoadTimeCycle ...
func (NoIOProvider) LoadTimeCycle() bool {
	return true
}

// LoadTime ...
func (NoIOProvider) LoadTime() int64 {
	return 0
}

// SaveTime ...
func (NoIOProvider) SaveTime(int64) {}

// LoadEntities ...
func (NoIOProvider) LoadEntities(ChunkPos) ([]Entity, error) {
	return nil, nil
}

// SaveEntities ...
func (NoIOProvider) SaveEntities(ChunkPos, []Entity) error {
	return nil
}

// LoadBlockNBT ...
func (NoIOProvider) LoadBlockNBT(ChunkPos) ([]map[string]interface{}, error) {
	return nil, nil
}

// SaveBlockNBT ...
func (NoIOProvider) SaveBlockNBT(ChunkPos, map[[3]int]map[string]interface{}) error {
	return nil
}

// SaveChunk ...
func (NoIOProvider) SaveChunk(ChunkPos, *chunk.Chunk) error {
	return nil
}

// LoadChunk ...
func (NoIOProvider) LoadChunk(ChunkPos) (*chunk.Chunk, bool, error) {
	return nil, false, nil
}

// WorldName ...
func (NoIOProvider) WorldName() string {
	return ""
}

// SetWorldName ...
func (NoIOProvider) SetWorldName(string) {}

// WorldSpawn ...
func (NoIOProvider) WorldSpawn() BlockPos {
	return BlockPos{0, 30, 0}
}

// Close ...
func (NoIOProvider) Close() error {
	return nil
}
