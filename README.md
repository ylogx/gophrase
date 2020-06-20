GoPhrase
========

![Go](https://github.com/ylogx/gophrase/workflows/Go/badge.svg)
[![snap gophrase](https://snapcraft.io//gophrase/badge.svg)](https://snapcraft.io/gophrase)
[![snap gophrase](https://snapcraft.io//gophrase/trending.svg?name=0)](https://snapcraft.io/gophrase)

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

##### Using Snap
The tool is available via the [Snap Store][snap-product-page].

```bash
snap install gophrase
```

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

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/gophrase)

[snap-product-page]: https://snapcraft.io/gophrase
[usage-sshot]: https://i.imgur.com/o8WUS0o.png
[xkcd-explained]: https://www.explainxkcd.com/wiki/index.php/936:_Password_Strength
