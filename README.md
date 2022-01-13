# Infant Lang Playground Sandbox API

## Dependencies
- [Mux](https://github.com/gorilla/mux) for routing
- [Infant Lang Interpreter](https://github.com/infant-lang/interpreter) The Infant Lang Interpreter

## Environment Variables
- `PORT`: The port to run the sandbox on. Change the value in the .env file or do manual port forwarding.

## Info
- It uses a modified version of the [Infant Lang Interpreter](https://github.com/infant-lang/interpreter) for the playground.
- Uses Mux for routing.
- `interpreter.go` is the main file for the interpreter.

## API
- `GET /api/1.0.1/play/` - Returns `READY` to indicate that the sandbox is ready to receive the POST request. (This is used because of the cold boot times of the Dyno instances)

- `POST /api/1.0.1/play/` - The POST request should contain the code to be executed.
    - Body:
        - `data`: The stringified JSON object containing the code to be executed.
    - Response:
        - JSON Object containing the result of the code execution.

The JSON object should contain the following fields:
```ts
{
    "code": string[] // array of lines of code
}
```


Enjoy 🎉