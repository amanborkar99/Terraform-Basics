# CMD Arguments and Environment Variables

* An environment variable is a way to supply dynamic configuration information to programs at runtime. Environment variables are often used to make the same program work in different environments like Local, QA, or Production
    * **os.Setenv(Key, Value)**
    * **os.Getenev(key)**
    * **os.Unsetenv(key)**

* Command-line arguments are a way to supply additional information to a program when it is started. The easiest way to supply command line arguments is by specifying a space separated list of values after the command while running it