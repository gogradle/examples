This is a sample project to demostrate how [Gogradle](https://github.com/blindpirate/gogradle) tests golang project and generates html reports.

No need to install `Go` or set `GOPATH` at all. Just install [JRE 8+](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html) and run 

```
./gradlew goTest
```
on POSIX or

```
gradlew goTest
```
on Windows,

then open `./.gogradle/reports/test/index.html` to see what happened!
