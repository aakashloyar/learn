package franzgopackage
import (
	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
}

func NewProducer(client *kgo.Client) (*Producer, error){
	return &Producer{
		client : client,
	}, nil 
}
func main() {
	
}