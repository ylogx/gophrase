GoPhrase
========

[![Build Status](https://travis-ci.com/ylogx/gophrase.svg?branch=master)](https://travis-ci.com/ylogx/gophrase)

A random passphrase generator.

### Why?

A password such as "Tr0ub4dor&3" is bad because it is easy for password cracking software and hard for humans to remember, leading to insecure practices like writing the password down on a post-it attached to the monitor.

On the other hand, a password such as "correcthorsebatterystaple" is hard for computers to guess due to having more entropy but quite easy for humans to remember.

<img src="https://imgs.xkcd.com/comics/password_strength.png" alt="xkcd password strength" width="600"/>

See [xkcd explained][xkcd-explained] for more details

### Usage

<!-- ![usage sshot][usage-sshot] -->
<img src="https://i.imgur.com/o8WUS0o.png" alt="usage sshot" width="400"/>

### Installation

##### Pre-built Binaries
Download binaries from release page and copy in `/usr/local/bin`

```bash
wget -O gophrase https://github.com/ylogx/gophrase/releases/download/v0.0.1/gophrase-v0.0.1.darwin.amd64
#wget -O gophrase https://github.com/ylogx/gophrase/releases/download/v0.0.1/gophrase-v0.0.1.linux.amd64
install gophrase /usr/local/bin
```

##### From source
```bash
git pull https://github.com/ylogx/gophrase
make install

# Ensure ~/.local/bin is in $PATH (i.e `[ -d $HOME/.local/bin ] && export PATH=$PATH:$HOME/.local/bin`) or use following command:
make build && install -pv gophrase /usr/local/bin
```

[usage-sshot]: https://i.imgur.com/o8WUS0o.png
[xkcd-explained]: https://www.explainxkcd.com/wiki/index.php/936:_Password_Strength
