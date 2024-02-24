# cmaker
go application to handle cmake builds, flags and add libraries for c++.

## What is working, what is missing
Right now only works the initing of a c++ project: basic cmakefiles are created with user-inserted parameters. Cmake-ing the result compiles without problems.

Everything else is missing.

The stack of the next things to do is:
- *delete*: removes files and folders created by the application, leaving the rest of the project untouched.
- *add-flag*: adds a compilation flag to the project cmake file (such as profile, debug flags, optimization flags, etc).
- *rm-flag*: undos the above command
- *add-lib*: downloads from a repository a library, compiles it and then links it to the project.
- *rm-lib*: un-links the library from the project.

## Compilation
On the root of the project, execute:

```console
make build
```


To clean every file generated by the program and make, run:
```console
make clean
```

## Other info
- Most of the files of the program are stored in a folder called ".cmaker", inside the folder where it is executed.
