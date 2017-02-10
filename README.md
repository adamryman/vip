# vip; vim pipe; or how I learned I want to learn more sed

```
go get github.com/adamryman/vip
godoc github.com/adamryman/vip | vip -c "normal 4dd" -c "%s/vip/CODING/g" | less
```

My `vim` foo has gotten good enough where I wanted to use it in a unix
pipeline. Take some input text, run some vim commands / macros on it, and then
pipe the altered text to `cut`, `head`, `pandoc`, etc.

You can pipe stdin into vim, but not pipe its output to stdout.
[Or rather, I could not find a way when I started making this](http://blog.robertelder.org/use-vim-inside-a-unix-pipe-like-sed-or-awk/)

First I tried doing it all in bash, and you can see the result below and a bash
script attempt in `vip.sh`.

```
FOOBAR=$(mktemp -d) rm ${FOOBAR}/tmp && \
	godoc github.com/adamryman/kit/dbconn | \
	vim -i "NONE" -u "NONE" \
	-c "normal 4ddjjVGd" \
	-c "saveas ${FOOBAR}/tmp" -c "wq!" \
	- ; cat ${FOOBAR}/tmp
```

As you can see, it looks scary.

So I made `vip`. Basically it does what the bash script does, but in golang.

```
cat main.go | vip -c "normal dd:s/vip/CODING/g<CR>" | less
```

Though I have learned `sed` is literally built for this. And it is streaming
I'll try that out soon... `man sed`.

### Issues

Need to redirect stderr of vim to a terminal. It hangs like
crazy if stdout is not a terminal

### LICENSE

This software is in the public domain. See the LICENSE.md.
