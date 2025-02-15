# gl
Wrapper for OpenGL and WebGL in a single interface.

This project is based on the work from [UnitofTime's Glitch internals](https://github.com/unitoftime/glitch).

## Caveats
This project attempts to provide as close an interface as possible to `go-gl/gl` and `go-gl/glfw` while being compatible with both Desktop and Web runtimes.

- By design, these interfaces impose no thread/goroutine-safety; all functions assume and expect they are being called from the same thread and thus share data that would cause race conditions if called from different threads/goroutines.
   - It's up to higher-level libraries to ensure thread-safe usage of these libraries.
- `normalgopher/gl` and `normalgopher/gl/fw` do their best to mimic the standard `go-gl/gl`/`go-gl/glfw` interfaces, but some calls have been changed to mimic the WebGL version.
   - (Calls like GetInteger* replaced with GetParameter).
### Shader Caveats
If you want to ignore the following caveats, you can prefix `VertexShaderHeader()` and `FragmentShaderHeader()` from the `normalgopher/gl/help` package to your shader source before shader compliation. If you choose to do this, make sure your shader source does not include the `#version` directive.

Otherwise, you'll need to keep the following in mind:

- Shaders should be declare their version as: `#version 300 es`
- Fragment shaders need to declare this near the top (but after the `#version` directive):

```glsl
#ifdef GL_ES
   precision highp float;
#endif
```

# Building/Running on Web
Since this project aims to provide a single library to build both desktop and web applications, choosing the build target is as simple as employing the GOOS and GOARCH.

## Building for Web
Running the following will produce a `main.wasm` file that in conjuction with a `wasm_exec.js` file ([found here](https://github.com/golang/go/blob/master/lib/wasm/wasm_exec.js)) will allow your projects to run in the browser (with a bit of html).

`GOOS=js GOARCH=wasm go build . -o main.wasm` 

Note: you can name your `.wasm` file something other than `main.wasm`.

## Necessary HTML
Following is a minimal HTML file that will, in conjuction with `main.wasm` (your project) and `wasm_exec.js`, will execute your project.

```html
<!DOCTYPE html>

<head>
   <meta charset="utf-8" />
   <script src="wasm_exec.js"></script>
   <script>
      const go = new Go();
      WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
         go.run(result.instance);
      });
   </script>
</head>

<body style="border: none; margin: 0;">
   <canvas id="glfw"></canvas>
</body>

</html>
```

### Three Important Things

1. Your files (`main.wasm` and `wasm_exec.js`) need to be provided via a webserver; otherwise the annoynance that is CORS will refuse to load your files.
2. If you chose to name your wasm something other than `main.wasm`, you'll need to update the `WebAssembly.instantiateStreaming` call to reference your specific filename.
3. Where ever the HTML canvas is, it **must** have an id of `glfw`; otherwise, this library won't know what to hook into.
