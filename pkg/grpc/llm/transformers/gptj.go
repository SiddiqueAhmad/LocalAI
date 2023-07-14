package transformers

// This is a wrapper to statisfy the GRPC service interface
// It is meant to be used by the main executable that is the server for the specific backend type (falcon, gpt3, etc)
import (
	"fmt"

	pb "github.com/go-skynet/LocalAI/pkg/grpc/proto"

	transformers "github.com/go-skynet/go-ggml-transformers.cpp"
)

type GPTJ struct {
	gptj *transformers.GPTJ
}

func (llm *GPTJ) Load(opts *pb.ModelOptions) error {
	model, err := transformers.NewGPTJ(opts.Model)
	llm.gptj = model
	return err
}

func (llm *GPTJ) Embeddings(opts *pb.PredictOptions) ([]float32, error) {
	return nil, fmt.Errorf("not implemented")
}

func (llm *GPTJ) Predict(opts *pb.PredictOptions) (string, error) {
	return llm.gptj.Predict(opts.Prompt, buildPredictOptions(opts)...)
}

// fallback to Predict
func (llm *GPTJ) PredictStream(opts *pb.PredictOptions, results chan string) {
	go func() {
		res, err := llm.gptj.Predict(opts.Prompt, buildPredictOptions(opts)...)

		if err != nil {
			fmt.Println("err: ", err)
		}
		results <- res
		close(results)
	}()
}
