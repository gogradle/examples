This is a sample project to demostrate how [Gogradle](https://github.com/blindpirate/gogradle) tests golang project and generates html reports.

No need to install `Go` or set `GOPATH` at all. Just install [JRE 8+](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html) and run 

```
./gradlew test
```
on POSIX or

```
gradlew test
```
on Windows and open [./.gogradle/reports/test/index.html](./.gogradle/reports/test/index.html) to see what happened!
