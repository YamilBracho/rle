RLE encoding
Run Length Encoding is a lossless data compression algorithm. It compresses data by reducing repetitive, and consecutive data called runs. It does so by storing the number of these runs followed by the data.

For the text AAAAAAAAAAAAAHHHEEM (19 characters), RLE will encode it to 13A3H2EM (7 characters).
For the text AAAAHHHEEM, HAHA., it will be encoded as 4A3H2E1M1,1 1H1A1H1A1. (21 characters).
