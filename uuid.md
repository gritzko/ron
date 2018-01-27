# RON UUID bit layout - compatibility

The bit layout is semi-compatible with RFC4122.  The RFC reserves m.s.bits of
the 8th byte for the "variant". RON hijacks the 0 variant (NCS backward
compatibility) assuming no original Apollo UUIDs are still in circulation.

    vvvv ....  .... ....  .... ....  .... ....
    .... ....  .... ....  .... ....  .... ....
    00ss ....  .... ....  .... ....  .... ....
    .... ....  .... ....  .... ....  .... ....

`ss` bits stand for the "scheme" of the RON UUID.  `vvvv` stand for "variety"
within the scheme.  For example, scheme: event timestamp `10`, variety: RFC4122
epoch `0010`.

The codes for schemes and their varieties:

00. human readable name
    0000. transcendental/hardcoded name (lww, rga)
       or a scoped name (myvar$gritzko)
    0001. ISBN (1/978$1400075997)
    0011. EAN-13 bar code (3/4006381$333931)
    0101. zip codes (5/2628CD$NL, 5/620078$RU)
    1010. IATA airport code (A/LED)
    1011. ticker name (B/GOOG$NASDAQ)
    1100. ISO 4217 currency code (C/USD, C/GBP)
    1101. short DNS name (D/google$com)
    1110. E.164 intl phone num (E/7999$5631415)
    1111. ISO 3166 country code (F/RU, F/FRA...)
01. number
    0000. decimal index (4%, 4%8)
    0001. SHA-1, first 120 bits
    0010. SHA-2
    0011. SHA-3
    1010. random number (A/k3R9w_2F8w%Le~6dDScsw)
    1100. crypto id, public key fingerprint
10. event (Lamport timestamp, value and origin)
    0000. Base64 calendar (MMDHmS...)
       origin is an app-scoped handle
    0001. Logical (1/00001, 10000000002...)
    0010. Epoch (RFC4122 epoch, 100ns since 1582)
    1000-1010. as 0000-0010, origin is a hash of an UUID
11. derived event (timestamp, same as event)

Event/derived RON UUIDs are very much like RFC4122 v1 time-based UUIDs.  The
difference is subtle but very important: RON event UUIDs are logical clocks.
They are *monotonous* and causally consistent. Practically speaking, RON UUIDs
assume an internet connection is/was available to somewhat synchronize the
clock.  A new RON event UUID is always greater than any past UUID.  A program
must reject any data with UUIDs "from the future".

8th byte's m.s. bit meanings are:

00. RON UUID,
01. RON atom (int, float, string),
10. RC4122 UUID,
11. Microsoft something.

## Base64 rendering

Any UUID, including an RFC4122 UUIDs, could be Base64-serialized as 22 chars

    A0123456789 8abcdefghij

where the leading char is `0` to `F` in each half.  The half-separating space is
optional.  The Base64 serialization reproduces the bit layout, verbatim. 

For scheme `00` or variety `0000`, the leading char might be skipped.
(The separating whitespace becomes non-optional then.)

In case of RON UUIDs, the leading char denoting the scheme might be replaced by
one of four special characters (`$` for `0`, `%` for `1`, `+` for `10`, `-` for
`11`).

Tailing zeroes might be skipped.  Then, non-zero leading char should be
separated by a slash `/`.  These are different renderings of the same UUID:

    ALED0000000 00000000000
    ALED000000000000000000
    ALED0000000 0000000000
    ALED0000000$0000000000
    ALED0000000
    A/LED000 0
    A/LED$0
    A/LED

In presence of context (previous UUIDs), UUIDs might be prefix-compressed.
