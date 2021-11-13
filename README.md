# Golang - Signals (Inspired by Godot)

```go
// Create new signal
s := NewSignal()

// Connect a function which will be executed after each Emit(...)
s.Connect(func1)

// Wait for a data from the signal after it's emitted
s.Wait()

// Connect a function that will be called once
s.ConnectOnce(func1)

// Emit's all the data and send's to all the functions
// Data is interface{} so it's could be any type you want
s.Emit("Some data)
```

# Use Cases - Games for example
```go
// Let's imagine this signal will be callen after
//    each key press on keyboard
// You can connect a key-processor func and check all the pressed key's
onKeyPress := NewSignal()

// This is a key-processor func
// It will NOT miss any call
onKeyPress.Connect(func(keyInterface interface{}) {
    // Let's image this is int with KEY_CODE
    key := keyInterface.(int)

    // Here we check that KEY is UP/DOWN
    // Then we move the player according to a pressed key
    if key == KEY_LEFT {
        player.X -= 1
    } else if key == KEY_RIGHT {
        player.X += 1
    }
})
```