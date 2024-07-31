# vort0_vrp
VRP solution for Vorto coding challenge

### Run Code on Single File
* `git clone https://github.com/bo0m3r31337/vort0_vrp.git`
* `cd vort0_vrp`
* `go install`
* `go mod tidy`
* `go run main.go load.go point.go driver.go util.go insert.go <path to file>`

## Run Binary with python script
* `git clone https://github.com/bo0m3r31337/vort0_vrp.git`
* `cd vort0_vrp`
* `go install`
* `go mod tidy`
* `go build`

### Run MAC
* `python <your_python_script.py> --cmd ./m.app --problemDir <path_directory_of_problems>`

### Run WIN
* `python <your_python_script.py> --cmd m.exe --problemDir <path_directory_of_problems>`

### Run Linux
* `python <your_python_script.py> --cmd ./m --problemDir <path_directory_of_problems>`

## low cost focus
I was able to implement a sort of greedy nearest neighbor approach with permutation swapping routes for a driver also sorting and reducing drivers where possible.  
I would like to get cost-savings -> distance matrix set up

