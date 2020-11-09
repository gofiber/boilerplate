# GoFiber Docker Boilerplate

![Release](https://img.shields.io/github/release/gofiber/boilerplate.svg)
[![Discord](https://img.shields.io/badge/discord-join%20channel-7289DA)](https://gofiber.io/discord)
![Test](https://github.com/gofiber/boilerplate/workflows/Test/badge.svg)
![Security](https://github.com/gofiber/boilerplate/workflows/Security/badge.svg)
![Linter](https://github.com/gofiber/boilerplate/workflows/Linter/badge.svg)


## Development

### Start the application 


```bash
go run app.go
```


## Production

```bash
docker build -t gofiber .
docker run -d -p 3000:3000 gofiber
```

Go to `http://localhost:3000`:


![Go Fiber Docker Boilerplate](./go_fiber_boilerplate.gif)
