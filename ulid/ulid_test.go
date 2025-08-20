package ulid

// crand "crypto/rand"
import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func TestULIDGeneration_Basic(t *testing.T) {
	// simple use
	id := ulid.Make()
	fmt.Println(id)
	assert.Equal(t, 26, len(id.String()))
}

func TestULIDGeneration_Custom_Entropy(t *testing.T) {
	// customized entropy
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	assert.Nil(t, err)
	assert.Equal(t, 26, len(id.String()))
	fmt.Println(id)
}

func TestULIDGeneration_Secure_Entropy(t *testing.T) {
	// // secure entropy
	// entropy = crand.Reader(rand.NewSource(time.Now().UnixNano()))
	// ms = ulid.Timestamp(time.Now())
	// fmt.Println(ulid.New(ms, entropy))
}
