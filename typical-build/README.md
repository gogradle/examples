This is a sample project to demostrate how to build a typical golang project with [Gogradle](https://github.com/blindpirate/gogradle). 

No need to install `Go` or set `GOPATH` at all. Just install [JRE 8+](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html) and run 

```
./gradlew goBuild 
```
on POSIX or

```
gradlew goBuild 
```
on Windows.

Then open following files in project root directory

- `StringReverse`
- `.gogradle/reports/test/index.html`
- `.gogradle/reports/coverage/index.html`

to see what happened!
