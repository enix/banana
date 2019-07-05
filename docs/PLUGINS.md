# Banana agent plugin reference

A banana plugin is an executable file. It'll be spawned by banana agent when needed.

## Installing a plugin

Deposit your plugin's executable into the `/etc/banana/plugins.d` directory. This directory can be changed by setting the `plugins_dir` in `banana.json` or by using the `--plugins-dir` command line flag. To choose the plugin to be used, use the `--plugin` (`-p`) flag with the plugin name as argument, or write its name in the `plugin` configuration key.

## Implementing a plugin

Plugin will be spawned with the main action to run as its first argument. The rest of the arguments will be action-dependant. It is supposed to output the required informations on `stdout` and logs on `stderr`. It should exit with status code 0 when everything went well, and with any other exit code on error.

Here is a list of the different actions that your plugin should be able to handle:

### Version

This should output the version of the plugin.

#### Input

```javascript
argv[1] = "version"
```

#### Output

The version as a plain text string.

#### Example

```
$ ./plugin version
v1.0.0
```

### Backup

Run a backup with the given parameters.

#### Input

```javascript
argv[1] = "backup"
argv[2] = "full" | "incremental"
argv[3] = storage url
argv[4...] = plugin_args from schedule configuration
```

#### Output

No output is expected for now (will be used to output metadata(s)).

#### Examples

```
$ ./plugin backup incremental s3://object-storage.r1.nxs.enix.io/bucket/backup-name --include / --exclude /proc
$ ./plugin backup full s3://object-storage.r1.nxs.enix.io/bucket/backup-name mongodb://user:pass@localhost:27017
```

### Restore

Restore a backup to the given target directory.

#### Input

```javascript
argv[1] = "restore"
argv[2] = timestamp/backup id to restore
argv[3] = target directory
```

#### Output

No output is expected.

#### Examples

```
$ ./plugin restore 20190702T131235Z /etc.bak
$ ./plugin restore 20190702T131235Z mongodb://user:pass@localhost:27017
```
