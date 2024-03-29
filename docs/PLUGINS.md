# Banana agent plugin reference

A banana plugin is an executable file. It'll be spawned by banana agent when needed.

## Installing a plugin

Deposit your plugin's executable into the `/etc/banana/plugins.d` directory. This directory can be changed by setting the `plugins_dir` in `banana.json` or by using the `--plugins-dir` command line flag. To choose the plugin to be used, use the `--plugin` (`-p`) flag with the plugin name as argument, or write its name in the `plugin` configuration key.

## Implementing a plugin

Plugin will be spawned with the main action to run as its first argument. The rest of the arguments will be action-dependant. It is supposed to output the required informations on `stdout`, logs on `stderr`, and custom data on FD 3. It should exit with status code 0 when everything went well, and with any other exit code on error.

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
v1.0.0 (backend v4.2.0)
```

### Backup

Run a backup with the given parameters.

#### Input

```javascript
argv[1] = "backup"
argv[2] = "full" | "incremental"
argv[3] = storage endpoint
argv[4] = bucket name
argv[5] = bucket prefix
argv[6...] = user specified args (plugin_args from schedule configuration || command line arguments after -)
```

#### Output

The plugin should write on standard output a JSON representation of the `BackupMetadata` struct. In additon, it can write on file descriptor 3 plugin-specific datas that will be downloadable from the UI. Those datas should be in `gzip` format.

#### Examples

```
$ ./plugin backup incremental s3://object-storage.r1.nxs.enix.io/bucket/backup-name --include / --exclude /proc
{ "size": 1234 }
$ ./plugin backup full s3://object-storage.r1.nxs.enix.io/bucket/backup-name mongodb://user:pass@localhost:27017
{ "size": 1234 }
```

### Restore

Restore a backup to the given target directory.

#### Input

```javascript
argv[1] = "restore"
argv[2] = timestamp/backup id to restore
argv[3] = storage endpoint
argv[4] = bucket name
argv[5] = bucket prefix
argv[6...] = user specified plugin args
```

#### Output

No output is expected.

#### Examples

```
$ ./plugin restore 20190702T131235Z /etc.bak
$ ./plugin restore 20190702T131235Z mongodb://user:pass@localhost:27017
```
