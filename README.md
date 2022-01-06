# Bitbucket Contributor Counting

`contributors` counts active contributors in Bitbucket Server.

## Usage
`contributors` takes three parameters, all passed by environment variable:
- `BITBUCKET_URL`: the URL to your Bitbucket Server instance.
- `BITBUCKET_USERNAME`: your username on the Bitbucket Server instance.
- `BITBUCKET_PASSWORD`: your password on the Bitbucket Server instance.

```bash
BITBUCKET_URL=$YOUR_BITBUCKET_URL BITBUCKET_USER=$YOUR_BITBUCKET_USERNAME BITBUCKET_PASSWORD=$YOUR_BITBUCKET_PASSWORD contributors
```

`contributors` is tested on:
- Ubuntu 14.04 / Bitbucket 5.9.1 / git 2.17.0
- Ubuntu 14.04 / Bitbucket 5.9.1 / git 1.9.1

## Example
```bash
$ BITBUCKET_URL=http://172.17.0.1:7990 BITBUCKET_USER=leo BITBUCKET_PASSWORD=Y9iF6kxhNoBs87YNBms9 contributors
Found 112 contributors:
 182 kevin@fossa.io
  91 leo@fossa.io
  92 foo@example.com
  ...
```
