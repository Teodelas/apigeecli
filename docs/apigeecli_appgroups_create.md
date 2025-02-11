## apigeecli appgroups create

Create an AppGroup

### Synopsis

Create an AppGroup

```
apigeecli appgroups create [flags]
```

### Options

```
      --attrs stringToString   Custom attributes (default [])
  -i, --channelid string       channel identifier identifies the owner maintaining this grouping
  -u, --channelurl string      A reference to the associated storefront/marketplace
  -d, --display-name string    app group name displayed in the UI
  -h, --help                   help for create
  -n, --name string            Name of the AppGroup
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --disable-check    Disable check for newer versions
      --no-output        Disable printing all statements to stdout
  -o, --org string       Apigee organization name
      --print-output     Control printing of info log statements (default true)
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli appgroups](apigeecli_appgroups.md)	 - Manage Apigee Application Groups

###### Auto generated by spf13/cobra on 27-Jul-2023
