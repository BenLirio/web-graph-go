package producer

type Producer struct {
	out chan<- string
}

func NewProducer(out chan<- string) *Producer {
	producer := Producer{out: out}
	return &producer
}

func (producer *Producer) generate() {
	producer.out <- "one"
	producer.out <- "two"
	producer.out <- "three"
	close(producer.out)
}
