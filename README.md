# go-kluge

This repository is an adaptation from [christiankastner/klugen-portrait](https://github.com/christiankastner/klugen-portrait), which is a implementation of [Espen Kluge Alternatives](https://www.espen.xyz/alternatives-new-portrait-series-2019)

# Installing

Install the command with go:
```bash
go install github.com/WendelHime/go-kluge/cmd/kluge
```

# Usage

```bash
kluge -h
Usage of kluge: kluge -filepath images/ada_in_blank.png -threshold 0.35 -minDist 50 -output images/ada_output.png
  -filepath string
        the filepath for a segmented image (default "/home/wotan/Pictures/ada_in_blank.png")
  -minDist float
        minimum distance of points (default 70)
  -output string
        output filepath (default "./output.png")
  -threshold float
        a threshold for random creation of points, the value must be in the range 0, 100. (default 0.25)
```

# Examples


Command:
```bash
kluge -filepath images/ada_in_blank.png -threshold 0.35 -minDist 50 -output images/ada_output.png
```

Original image:
<img src="https://github.com/WendelHime/go-kluge/raw/main/images/ada_in_blank.png" width="320" height="400">

Output:
<img src="https://github.com/WendelHime/go-kluge/raw/main/images/ada_output.png" width="320" height="400">
