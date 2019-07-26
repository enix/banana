# Bananagent command-line examples

## Backup

Usage: `bananagent <b|backup> <full|incremental> <name> [-] [plugin args...]`

* Full backup /etc using duplicity.

```bash
bananagent backup full "etc directory" /etc
```

* Incrementally backup /etc using duplicity (will fail if there is no previous full bacukps). The - is facultative beacause there is no flags for the plugin.

```bash
bananagent b incremental "etc directory" - /etc
```

* Pass complex arguments to duplicity. The - is needed to avoid --exclude flags being parsed by bananagent.

```bash
bananagent b full "full host" - / --exclude /proc --exclude /sys --exclude /dev --exclude /tmp
```

* Backup a mysql database. The full/incremental choice will be ignored by the plugin because it makes no sense for a database.

```bash
bananagent b full "my database" -p mysqldump - --all-databases
```

## Restore

Usage: `bananagent <r|restore> <name> <backup id> [-] [plugin args...]`

* Restore a /etc backup. You can get the id only from the filenames in the storage (for now). Yous hould restore to /etc.bak because duplicity will refuse to override the existing /etc.

```bash
bananagent restore "etc directory" 20190702T131235Z /etc.bak
```

* Restore a mysql database.

```bash
bananagent r "my database" -p mysqldump
```
