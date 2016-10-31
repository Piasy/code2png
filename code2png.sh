#!/bin/bash
highlight -i $1 -o "$1.html" --include-style --line-numbers --font Hack,\'msyh\' && \
sed -i -- 's/charset=ISO-8859-1/charset=utf-8/g' "$1.html" && \
/root/download/wkhtmltox/bin/wkhtmltopdf --quiet --dpi 1200 "$1.html" "$1.ps" && \
gs -q -dBATCH -dNOPAUSE -dSAFER -dNOPROMPT -sDEVICE=png16m -dDEVICEXRESOLUTION=600 -dDEVICEYRESOLUTION=600 -dDEVICEWIDTH=4958 -dDEVICEHEIGHT=7017 -dNOPLATFONTS -dTextAlphaBits=4 -sOutputFile="$1.origin.png" "$1.ps" && \
convert -trim +repage -trim +repage -bordercolor "#f0f0f0" -border 25x25 "$1.origin.png" "$1.png" && \
mv "$1.png" images/ && \
rm $1 "$1.html" "$1.ps" "$1.origin.png"

