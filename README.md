# color-extract
Extracts the dominant color palettes from a set of images and outputs them as a list of hex colors. 

Supports GIF, PNG and JPEG.

Uses [quant](https://github.com/soniakeys/quant) to provide both median and mean quantizers.

Default palette size is 16 and capped at 256 colors.
Default quatization mode is median.

##Usage

```
color-extract --size=16 --mode=median /path/to/file1.png /path/to/file2.png...
```
