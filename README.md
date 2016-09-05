# Arc API Proxy
	
### Development:

	go get github.com/codegangsta/gin
	gin -p 8080

This above will run the application at port 5000

### Build and Deploy:

If you wish the build the application in the traditional "go way",

	go build

To run the binary file, run:

	export PORT=8080 && ./arc-api-proxy

For the app to work correctly, please ensure you fill out the ``.env`` with the correct environment variable. It will use the following by default if no environment variable given.
```
	ARC_API_URL="https://api.spectator.arcpublishing.com"
	ARC_API_USERNAME="spectator"
	ARC_API_PASSWORD="password"
```

### Build Docker Image/Deployment

#### Build Image
```
docker build -t columbiaspectator/arc-api-proxy .
```

#### Run Image
```
docker run -d -p 80:8080 -e PORT=8080 -e ARC_API_URL="https://api.spectator.arcpublishing.com" -e ARC_API_USERNAME="spectator" -e ARC_API_PASSWORD="password" columbiaspectator/arc-api-proxy

```

### Dependency Vendoring

We are currently using ``godep`` to manage dependencies. All dependencies are tracked in ``Godeps.json`` and copied into the ``Godeps`` directory. If the project directory is changed, say from ``tim`` to ``timothy``, run the following. However, if any of the source files are referencing same-project package(s), you need to change those import paths first.

For example, I have a package in the same project named ``connection`` and my old project name is ``tim``, I will have to reference ``connection`` pacakge in my ``main.go`` as ``tim/connection``. Now the directory name is ``timothy`` and I will have to change it to ``timothy/connection`` and then run the following,

	godep save -r ./...
		
``godep`` is not smart enough yet to distinguish whether the package belongs to the same project or is an external dependency.

Go Debugging
---

If you have used Rails, you will miss using binding.pry. However, Go does have similar. Go has a package called "Godebug"

Install Godebug by running:

	go get github.com/mailgun/godebug

Insert a breakpoint anywhere in a source file you want to debug:
	
	_ = "breakpoint"

Replace ``<pkgs>`` with the packages we will be debugging if it is not the ``main`` package

	godebug build -instrument <pkgs>

For example:
	
	godebug build -instrument arc-api-proxy/connection

godebug will generate a binary named ``yourprojectname.debug``, run that binary with the necessary arguments or environment variables and use it as you would binding.pry

For example,

	PORT=8080 ./arc-api-proxy.debug
