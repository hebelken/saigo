package corpus

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCorpus(t *testing.T) {
  result := corpus.Analyze("Are you serious? I knew you were.")
  assert.Equal(len(result), 5)
  assert.Equal(result[0].Word, "you")
  assert.Equal(result[0].Count, 2)
}
