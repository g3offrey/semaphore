# semaphore

semaphore is a Go package, implementing a simple counting semaphore

## Installation

```bash
go get github.com/g3offrey/semaphore
```

## Usage

```go
import "github.com/g3offrey/semaphore"

func TestSemaphore() {
    sem := semaphore.Make(capacity)

    sem.Acquire()
    // statements ...
    sem.Release()
}

```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)