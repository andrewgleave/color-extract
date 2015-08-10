# color-extract
Extract the dominant color palette from an image.

Uses [quant](https://github.com/soniakeys/quant) to provide both median and mean quantizers.

Default palette size is 16 colors and capped at 256.
Default quatization mode is median.

##Usage

```
color-extract --size=16 --mode=median /path/to/file1.png /path/to/file2.png...
```
