# Go Utils

- Build docker image
```bash
docker build -t utils/go-utils .
```

- Run test
```bash
docker run --rm utils/go-utils
```

- Work in docker container
```bash
docker run -it --rm -v $(pwd):/go-utils utils/go-utils sh
```
