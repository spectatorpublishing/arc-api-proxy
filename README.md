# GoRestApi
	
### Development:

	go get github.com/codegangsta/gin
	gin -p 5000

This above will run the application at port 5000

### Build and Deploy:

If you wish the build the application in the traditional "go way",

	go build

To run the binary file, run:

	export PORT=8080 && ./arc-api-proxy

For the app to work correctly, please ensure you fill out the ``.env`` with the correct environment variable. It will use the following by default if no environment variable given.

	MONGO_URL=mongodb://localhost:27017

You can change the port to the one your prefer.

### Seed User:
In Mongo console, run
```
use database_base_name
db.users.insert({"access_token": "test"})
```


TODOS:

Add ability to configure more Mongodb settings.

### Usage:

Currently, only one endpoint functions and you can make a sample request to see if everything is running

	POST   http://localhost:5000/events?name=chicago

Sample Response:

	{
	  "id": "56ae2b503d10891dced23e05",
	  "event_date": "0001-01-01T00:00:00Z",
	  "city": "",
	  "name": "chicago",
	  "country": "",
	  "weather_url": "",
	  "updated_at": "2016-01-31T10:42:08.44106907-05:00",
	  "created_at": "2016-01-31T10:42:08.44106907-05:00"
	}

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

### Dealing with Mongodb

Enter the console by typing ``mongo`` in terminal

The following command allows you to rename field/column

	db.events.update({},{ $rename: { 'current_field_name': 'new_name'}}, { multi: true })
