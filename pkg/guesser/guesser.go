package guesser

import (
	"os"
	"sort"

	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
	"github.com/robherley/guesslang-go/pkg/errors"
)

func init() {
	// Supress TensorFlow logging by default
	// https://github.com/tensorflow/tensorflow/blob/74d4bcde8b51963f8c7401d118e17b987d6c93fd/tensorflow/tsl/platform/default/logging.h#L67-L71
	os.Setenv("TF_CPP_MIN_LOG_LEVEL", "5")
}

type Guesser struct {
	model *tg.Model
}

// New initializes a guesslang model. It will write the TensorFlow SavedModel temporarily to disk, then load it.
func New() (g *Guesser, err error) {
	defer func() {
		// unfortunately, the tfgo library panics instead of returning errors
		if r := recover(); r != nil {
			err = errors.NewFailedLoad(r)
		}
	}()

	modelPath, err := writeModelToTempDir()
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(modelPath)

	model := tg.LoadModel(modelPath, []string{"serve"}, nil)
	return &Guesser{model}, nil
}

// Guess executes the guesslang model on a code snippet, providing sorted confidences for all of the supported languages
// and will also provide whether or not the answer is arbitrarily reliable.
func (g *Guesser) Guess(snippet string) (a *Answer, err error) {
	results, err := g.exec(snippet)
	if err != nil {
		return nil, err
	}

	if len(results) != 2 {
		return nil, errors.NewInvalidResult("expected two outputs")
	}

	var languages []string
	if result, ok := results[0].Value().([][]string); ok && len(result) == 1 {
		languages = result[0]
	} else {
		return nil, errors.NewInvalidResult("invalid result for languages")
	}

	var confidences []float64
	if result, ok := results[1].Value().([][]float32); ok && len(result) == 1 {
		confidences = make([]float64, len(result[0]))
		for i, v := range result[0] {
			confidences[i] = float64(v)
		}
	} else {
		return nil, errors.NewInvalidResult("invalid result for confidences")
	}

	if len(languages) != len(confidences) {
		return nil, errors.NewInvalidResult("mismatch between languages and confidences")
	}

	predictions := make([]Prediction, 0, len(languages))
	for i := range languages {
		predictions = append(predictions, Prediction{
			Confidence: confidences[i],
			Language:   languages[i],
		})
	}

	sort.Slice(predictions, func(i, j int) bool {
		return predictions[i].Confidence > predictions[j].Confidence
	})

	return &Answer{
		Predictions: predictions,
		Reliable:    IsReliable(confidences),
	}, nil
}

func (g *Guesser) exec(snippet string) (results []*tf.Tensor, err error) {
	defer func() {
		// unfortunately, the tfgo library panics instead of returning errors
		if r := recover(); r != nil {
			err = errors.NewFailedExec(r)
		}
	}()

	input, err := tf.NewTensor([1]string{snippet})
	if err != nil {
		return nil, err
	}

	// see docs/model-info.md
	results = g.model.Exec([]tf.Output{
		g.model.Op("head/Tile", 0),
		g.model.Op("head/predictions/probabilities", 0),
	}, map[tf.Output]*tf.Tensor{
		g.model.Op("Placeholder", 0): input,
	})

	return results, nil
}
