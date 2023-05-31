# Prerequisites:
- Golang should be installed in the environment where the application will be built and run.
- Basic understanding of Makefile and command-line interface.

## Project Lifecycle:
- `build`: compile the source code of the application to a binary named `awesome-api` using the command `go build`.
- `run`: run the application in the background by executing the binary `awesome-api` and write logs into a file named `awesome-api.log` using the command `./awesome-api >./awesome-api.log 2>&1 &`.
- `stop`: stop the application using the command `kill XXXXX` where XXXXX is the Process ID of the application.
- `clean`: stop the application and delete the binary `awesome-api` and the log file `awesome-api.log`.
- `test`: test the application to ensure that it behaves as expected.
- `lint`: run the application through lint to check for errors
- `unit-tests`: run tests to check for validation and HelloHandler coverage

## Requirements:
- A Makefile should be present and valid.
- The binary `awesome-api` must NOT exist at the beginning, in the source code.
- The goals `build`, `run`, `stop`, `clean`, and `test` should be implemented and mapped to the life-cycle stages of the same name.
- The `help` goal must be implemented and print a list of all the goals with a sentence each.