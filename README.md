# fix - a utility to fix ripped MP3's names.

## Install prerequisites for ripping CDs to MP3s

In order to rip your CDs, install `abcde`, `cdparanoia` and `lame`, then add the following to your `.bashrc`:

```bash
alias rip="abcde -a cddb,read,encode,tag,clean,move -B -G -p -g -j 4 -o mp3:\"--preset extreme\" -x"
```

I can't remember (nor have time to spend checking) if it also requires packages `id3tools` and `id3v2` to properly tag MP3s.

## Compile `fix`

In order to compile `fix`, you need a Golang compiler; simply clone the repo and run `go build`, then copy `fix` to your path.

## Rip CDs and fix their names

In a directory of your choice, run `rip`; it will calculate the dicid and use it to dowload CD info and artwork, prompting for confirmation.

Once the rip process has completed, run `fix <name of directory>` and it will automatically adjust filenames.



