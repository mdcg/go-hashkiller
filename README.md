# HASHKILLER

[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mdcg/hashkiller/blob/master/LICENSE)

_A simple password cracker written in Go_

Inspired by "[John The Ripper](https://tools.kali.org/password-attacks/john)", this is my implementation of a password cracker! It currently has analysis for MD5, SHA1 and SHA256 encryptions. To use, just build the package and use something like:

```
./hashkiller -hash=417f4ef874925faebdcc8c79b32246ca -type=md5 -wordlist=awesome_wordlist.txt
```

Where
- In "hash" you will need to enter the hash you want to break;
- In "type", it will be the encryption type (currently we have MD5, SHA1 and SHA256);
- In "wordlist" you need to inform the path from your directory to the wordlist you will use (you can easily generate one with "[Crunch](https://tools.kali.org/password-attacks/crunch)").


### Contributing

Feel free to do whatever you want with this project. :-)