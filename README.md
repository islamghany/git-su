# git-su

A simple git wrapper to make it easier to work with multiple git accounts.

## Installation

- Clone the repository
- Depending on your OS, you may need to make the file executable. You can do this by running `chmod +x git-su`
- Move the file to a directory in your PATH. For example, `/usr/local/bin`

### Linux-x64

```bash
git clone
cd git-su
chmod +x git-su
sudo cp bin/linux-x64/* /usr/local/bin/
```

## Usage

### Current account info

```bash
git-su which
```

### Add a new account

```bash
git-su add -id=<id> -name=<name> -email=<email>
```

```bash
git-su add -id=work -name="John Doe" -email="work@work.com"
```

### Switch to a different account

```bash
git-su id
```

```bash
git-su work
```

### List all accounts

```bash
git-su list
```

```bash
git-su ls
```

### Remove an account

```bash
git-su remove -id=<id>
```

```bash
git-su rm -id=work
```

### Help

```bash
git-su --help
```

```bash
git-su -h
```
