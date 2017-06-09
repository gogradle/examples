This is a sample project to demostrate how [Gogradle](https://github.com/blindpirate/gogradle) handles task name conflit.

Gogradle's task name `build`/`test`/`check` may conflict with other plugins. To solve this, just put following lines in `<PROJECT_ROOT>/gradle/gradle.properties` or `~/.gradle/gradle.properties`:

```
org.gradle.jvmargs=-Dgogradle.alias=true
```

And run tasks with aliases as documented in [Tasks](https://github.com/gogradle/gogradle/blob/master/docs/tasks.md#task-name-conflict).
