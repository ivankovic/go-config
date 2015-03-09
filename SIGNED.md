##### Signed by https://keybase.io/ivankovic
```
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2.0.22 (GNU/Linux)

iQIcBAABAgAGBQJU/jOaAAoJEHL+w/ZMh983IX4P/31a1brWAJMYNvHqffSvPReO
K4U+ykw64BaBCJG/KgG3YwzLtVzrTMFwakxdE2Z959a3SxfyvOBqtBGQ7vK/VDRy
nmib3nQSWI4NHOReuvF8CE2yCRVzmEwm9aJoWdVxnNYMNLq8VDLJYRAbq1Skwxc4
RBQ+e1DrGWKM2FsKb4rmCb+cR5mjtpl3K8BRksCAM5Xk8/41CbNuN/D13N0ysa7v
2eHZw/U8PFKlfGsdDs8PniiqrvG0vl8bZYHYXRSp5tMb4kHX3Kf/wBbJnnpOFnl9
En7VU1imB4LrFqeNCqBfrkk4eHMfYttwD8P2N+1FVldZwdg2rKqI3Aepvi7l6RGs
A8XKidouhMErfzdPETBtw8X27ZkupAM3YGdb6IViZhn3YkxBgv2FO4WyCa6W6Zju
VKzzwCkZk0b94awuihRyCvBkKxBYRXfwZPCscDzc+yuSDkys0x8bqUkzvFKrLp3I
Y/8F1y7Z0uIfM62isnqWqtKuW1e9n6fwmXLm+RTnpRTp4p9CQ0yzyiiK5hmKv2Os
uANtm/JvBpdxw1LsxUF9BJBJWVR6GiUpCBQ79Hg8l4t6P+coqY4DcL6avBwQ8cDv
FBeC835+qtzfkrdTQmJ0uYcs0sL9UYp1+Eu6POrgPcNNkv+Gjott3Zw6gWoQHzaJ
mmUmH5PQmo+Mlz9ERy0f
=SG4g
-----END PGP SIGNATURE-----

```

<!-- END SIGNATURES -->

### Begin signed statement 

#### Expect

```
size   exec  file                  contents                                                        
             ./                                                                                    
266            .gitignore          cb568f716e3315bcebfc75bbc274b56577c2734ec3da6729ffac15919f416240
11325          LICENSE             cb5e8e7e5f4a3988e1063c142c60dc2df75605f4c46515e776e3aca6df976e14
186            README.md           42697970ae6595d230b2d37f78b151f35bc85795ff986152fe95370c5294b1da
               config/                                                                             
2064             config.go         11b47f9fa5c6e911f8017fae12c4d1f6e216f4e6ea5c10d6a6825f90832fa29c
2629             config_test.go    d6a3fb6bd2855db7124ffff93600e63ef7d81cf05ac0c12926764e33b29daa02
163              test_config.conf  816bcff7ef10a73f4dd283302f6bdc3fda04eb37f7a42f4338e809e526c19928
```

#### Ignore

```
/SIGNED.md
```

#### Presets

```
git      # ignore .git and anything as described by .gitignore files
dropbox  # ignore .dropbox-cache and other Dropbox-related files    
kb       # ignore anything as described by .kbignore files          
```

<!-- summarize version = 0.0.9 -->

### End signed statement

<hr>

#### Notes

With keybase you can sign any directory's contents, whether it's a git repo,
source code distribution, or a personal documents folder. It aims to replace the drudgery of:

  1. comparing a zipped file to a detached statement
  2. downloading a public key
  3. confirming it is in fact the author's by reviewing public statements they've made, using it

All in one simple command:

```bash
keybase dir verify
```

There are lots of options, including assertions for automating your checks.

For more info, check out https://keybase.io/docs/command_line/code_signing