Table of Contents
=================

   * [Swarm Replicated Object Notation 2.0.1](#swarm-replicated-object-notation-201)
      * [Formal model](#formal-model)
      * [Wire format (text)](#wire-format-text)
      * [Wire format (binary)](#wire-format-binary)
         * [Descriptors](#descriptors)
         * [Op terms](#op-terms)
         * [Uncompressed UUIDs](#uncompressed-uuids)
         * [Compressed UUIDs](#compressed-uuids)
         * [Atoms](#atoms)
      * [The math](#the-math)
      * [Acknowledgements](#acknowledgements)
      * [History](#history)
      * [Build status](#build-status)

# Swarm Replicated Object Notation 2.0.1

Swarm Replicated Object Notation is a format for *distributed live data*.  RON's
focus is on continuous data synchronization.  Every RON *object* may naturally
have an unlimited number of *replicas* that synchronize incrementally, mostly in
real-time.  RON data always merges correctly and deterministically.

RON is information-centric: it aims to liberate the data from its location,
storage, application or transport.  There is no "master" replica, no "source of
truth". Every event has an *origin*, but every replica is as good as the other
one.  Every single object, event or data type is uniquely identified and
globally referenceable.  RON metadata makes objects completely independent of
the context.  A program may read RON object versions and/or updates from the
network,  filesystem, database, message bus and/or local cache, in any order,
and merge them correctly.

Consider JSON. It expresses relations by element positioning:

    {
        "foo": {
            "bar": 1
        }
    }

RON may express that state as:

    *lww #1TUAQ+gritzko @`   :bar = 1;
         #(R            @`   :foo > (Q;

Those are two RON *ops*:

 1. some last-write-wins object is created with a field `bar` set to `1` (on
    2017-10-31 10:26:00 UTC, by gritzko),
 2. another object is created with a field `foo` pointing to the first object
    (10:27:00, by gritzko).

Each op is a tuple of four globally-unique UUIDs for its data type, object,
event and location, plus some number of value *atoms*.  You may not see any
UUIDs in the above example, initially.  The notation does a lot to compress that
metadata away.

These are the key features of RON:

* RON's basic unit is an immutable *op*. Every change to the data is an
  *event*; every event produces an op. An op may flow from a replica to a
  replica, from a database to a database, while fully intact and maintaining its
  original identity.
* Each RON op is context-independent. Nothing is implied by the context,
  everything is specified explicitly and unambiguously in the op itself. An op
  has four globally unique UUIDs for its data type, object, event and location.
* An object can be referenced by its UUID (e.g. `> 1TUAQ+gritzko`), thus RON can
  express object graph structures beyond simple nesting.  Overall, RON relates
  pieces of data by their UUIDs.  Thanks to that, RON data can be cached
  locally, updated incrementally and edited while offline.
* An object's state is a *reduction* of its ops. A data type is a reducer
  function: `lww(state,change) = new_state`. Reducers tolerate partial order of
  updates. Hence, all ops are applied immediately, without any linearization by
  a central server.
* There is no sharp border between a state snapshot and a state update. State is
  change and change is state (state-change duality). A transactional unit of
  data storage/transmission is a *frame*. A frame can contain a single op, a
  complete object graph or anything inbetween: object state, stale state, patch,
  otherwise a piece of an object.
* RON model implies no special "source of truth". The event's *origin* is the
  source of truth, not a server in the cloud. Every event/object is marked with
  its origin (e.g. `gritzko` in `1TUAQ+gritzko`).
* A RON frame is not a "message": it has an *origin* but it has no
  "destination". RON speaks in terms of data updates and subscriptions.  Once
  you subscribe to an object, you receive the state and all the future updates,
  till you unsubscribe.
* RON is information-centric. Consider git: once you clone a repo, your copy is
  as good as the original one. Same with RON.
* RON is a hypermedia format, as data pieces can reference each other globally
  (imagine a RON-based real-time World-Wide-Web-of-Data).  Although, both
  replica ids and data routing must work at global scale then (federated, etc).
* RON is not optimized for human consumption. It is a machine-to-machine
  language mostly. "Human" APIs are produced by mappers (see below).
* RON employs compression for its metadata. The RON UUID syntax is specifically
  fine-tuned for easy compression.

Consider the above frame uncompressed:

    *lww #1TUAQ+gritzko @1TUAQ+gritzko :bar = 1;
    *lww #1TUAR+gritzko @1TUAR+gritzko :foo > 1TUAQ+gritzko;


One may say, what metadata solves is [naming things and cache
invalidation][2problems].  What RON solves is compressing that metadata.

RON makes no strong assumptions about consistency guarantees: linearized,
causal-order or gossip environments are all fine (certain restrictions apply,
see below).  Once all the object's ops are propagated to all the object's
replicas, replicas converge to the same state.  RON formal model makes this
process correct.  RON wire format makes this process efficient.


## Formal model

Swarm RON formal model has five key components:

1. An UUID is a globally unique 128-bit identifier. An UUID consists of two
60-bit parts: *value* and *origin*. 4+4 bits are reserved for flags. There are
four UUID types:
    * an event timestamp: logical/hybrid timestamp, e.g. `1TUAQ+gritzko`, value
      is a monotonous counter `1TUAQ`, origin is a a replica id `gritzko`,
      roughly corresponds to RFC4122 v1 UUIDs,
    * a derived timestamp: same as event timestamp, but refers to some derived
      calculation, not the original event (e.g. `1TUAQ-gritzko`),
    * a name, either global or scoped to a replica, e.g. `foo`, `lww`, `bar`
      (global), `MyVariable$gritzko` (scoped),
    * a hash (e.g. `4Js8lam4LB%kj529sMEsl`, both parts are hash sum bits).

2.  An op is an immutable atomic unit of data change. An op is a tuple of four
    or more *atoms*. First four atoms of an op are UUIDs forming the op's key.

    These UUIDs are:

    1. data type UUID, e.g. `lww` a last-write-wins object,
    2. object UUID `1TUAQ+gritzko`,
    3. event UUID `1TUAQ+gritzko` and
    4. location/reference UUID, e.g. `bar`.

    Other atoms (any number, any type) form the op's value. Op atoms types are:

    1. UUID,
    2. integer,
    3. string, or
    4. float.

    Importantly, an op goes under one of four *terms*:

    1. raw ops (a single op, before being processed by a reducer),
    2. reduced ops (an op in a frame, processed by a reducer),
    3. frame headers (first op of a frame, planted by a reducer),
    4. queries (part of connection/subscription state machines).

3. A frame is an ordered collection of ops, a transactional unit of data

    * an object's state is a frame
    * a "patch" (aka "delta", "diff") is also a frame
    * in general, data is seen as a [partially ordered][po] log of frames
      or chunks
    * frame may contain any number of reduced chunks and raw ops in any order;
      a chunk consists of a header or a query header op followed by reduced ops
      belonging to the chunk; raw ops form their own one-op chunk.

4.  A reducer is a RON term for a "data type"; reducers define how object state
    is changed by new ops

    *   a reducer is a pure function: `f(state_frame, change_frame) ->
        new_state_frame`, where frames are either empty frames or single ops or
        products of past reductions by the same reducer,

    *   reducers are:

        1.  associative, e.g. `f( f(state, op1), op2 ) == f( state, patch )`
            where `patch == f(op1,op2)`
        2.  commutative for concurrent ops (can tolerate causally consistent
            partial orders), e.g. `f(f(state,a),b) == f(f(state,b),a)`, assuming `a`
            and `b` originated concurrently at different replicas,
        3.  idempotent, e.g. `f(state, op1) == f(f(state, op1), op1) == f(state,
            f(op1, op1))`, etc.

    *   optionally, reducers may have stronger guarantees, e.g. full commutativity
        (tolerates causality violations),

    *   a frame could be an op, a patch or a complete state. Hence, a baseline
        reducer can "switch gears" from pure op-based CRDT mode to state-based
        CRDT to delta-based, e.g.

        1. `f(state, op)` is op-based
        2. `f(state1, state2)` is state-based
        3. `f(state, patch)` is delta-based

4. a mapper translates a replicated object's state frame into other formats

    * mappers turn RON objects into JSON or XML documents, C++, JavaScript or
      other objects
    * mappers are one-way: RON metadata may be lost in conversion
    * mappers can be pipelined, e.g. one can build a full RON->JSON->HTML
      [MVC][mvc] app using just mappers.


Single ops assume [causally consistent][causal] delivery.  RON implies causal
consistency by default.  Although, nothing prevents it from running in a
linearized [ACIDic][peterb] or gossip environment.  That only relaxes (or
restricts) the choice of reducers.

## Wire format (text)

Design goals for the RON wire format is to be reasonably readable and reasonably
compact.  No less human-readable than regular expressions.  No less compact than
(say) three times plain JSON (and at least three times more compact than JSON
with comparable amounts of metadata).

The syntax outline:

1. atoms follow very predictable conventions:
    * integers: `1`
    * e-notation floats: `3.1415`, `1.0e+6`
    * UTF-8 JSON-escaped strings: `строка\n线\t\u7ebf\n라인`,
      except that `'` (U+0027 APOSTROPHE) must be encoded as `\u0027` or `\'`
    * RON UUIDs `1D4ICC-XU5eRJ`, `1TUAQ+gritzko`
2. UUIDs use a compact custom serialization
    * RON UUIDs are Base64 to save space (compare [RFC4122][rfc4122]
      `123e4567-e89b-12d3-a456-426655440000` and RON `1D4ICC-XU5eRJ`)
    * also, RON timestamp UUIDs may vary in precision, like floats (no need to
      mention nanoseconds everywhere) -- trailing zeroes are skipped
    * UUIDs are lexically/numerically comparable (same order), the Base64
      variant is
      `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~`
3. serialized ops use some punctuation, e.g. `*lww #1D4ICC-XU5eRJ
   @1D4ICC2-XU5eRJ :keyA 'valueA'`
    * `*` starts a data type UUID
    * `#` starts an object UUID
    * `@` starts an op's own event UUID
    * `:` starts a location UUID
    * `=` starts an integer
    * `'` starts and ends a string;
          may occur inside a string if prefixed by backslash — `\'`
    * `^` starts a float (e-notation)
    * `>` starts an UUID
    * `!` ends a frame header op (a reduced chunk has one header op)
    * `?` ends a query header op (a subscription frame has a header)
    * `,` ends a reduced op (optional)
    * `;` ends a raw op
    * `.` ends a frame (required for streaming transports, e.g. TCP)
4. frame format employs cross-columnar compression
    * repeated key UUIDs can be skipped altogether ("same as in the last op");
      in the first op all key UUIDs are mandatory;
    * RON abbreviates similar UUIDs using prefix compression, e.g.
      `1D4ICCE+XU5eRJ` gets compressed to `{E` if preceded by `1D4ICC+XU5eRJ`
      (symbols `([{}])` corespond to 4,5,..9 symbols of shared prefix)
    * by default, a key UUID is compressed against the same UUID in the previous
      op (e.g. event id against the previous event id);
    * backtick \` changes the default UUID to the previous UUID of the same op
      (e.g. event id against same op's object id)
    * the first value UUID is compressed against the object UUID of the op,
      each other is compressed against the previous one.

Consider a simple JSON object:

    {"keyA":"valueA", "keyB":"valueB"}

A RON frame for that object will have three ops: one frame header op and two
key-value ops.  In the tabular form, that frame may look like:

    type object         event           location value
    -----------------------------------------------------
    *lww #1D4ICC+XU5eRJ @1D4ICCE+XU5eRJ :0       !
    *lww #1D4ICC+XU5eRJ @1D4ICCE+XU5eRJ :keyA    'valueA'
    *lww #1D4ICC+XU5eRJ @1D4ICC1+XU5eRJ :keyB    'valueB'

There are lots of repeating bits here.
We may skip repeating UUIDs and prefix-compress close UUIDs.
The compressed frame will be just a bit longer than bare JSON:

    *lww#1D4ICC+XU5eRJ@`{E! :keyA'valueA' @{1:keyB'valueB'

The frame contains *twelve* UUIDs (6 distinct UUIDs, 3 distinct
timestamps) and also the data.
Despite the impressive amount of metadata, it takes less space than *two* [RFC4122
UUIDs][rfc4122].  The point becomes even clearer if we add the
object UUID to JSON using the RFC4122 notation:

    {"_id": "0651a600-2b49-11e6-8000-1696d3000000", "keyA":"valueA",
    "keyB":"valueB"}


We may take this to the extreme if we consider the case of a CRDT-based
collaborative real-time editor.  Then, every letter in the text has its own
UUID.  With RFC4122 UUIDs and JSON, that is simply ridiculous.  With RON, that
is perfectly OK.

Consider "Hello world!" collaboratively written by two users, `bart` and `lisa`
on 27 Nov 2017 around 9am GMT.  A compressed RGA (Replicated Growable Array)
frame would look like:

    *rga#1UQ8p+bart@1UQ8yk+lisa:0!
        @(s+bart'H'@[r'e'@(t'l'@[T'l'@[i'o'
        @(w+lisa' '@(x'w'@(y'o'@[1'r'@{a'l'@[2'd'@[k'!'

The `txt` mapper may convert the RGA frame into text:

    *txt #1UQ8p+bart @1UQ8yk+lisa 'Hello world!'

If nicely indented, the compressed frame is easier to read:

    *rga #1UQ8p+bart @1UQ8yk+lisa :0  !
                     @(s+bart        'H'
                     @[r             'e'
                     @(t             'l'
                     @[T             'l'
                     @[i             'o'
                     @(w+lisa        ' '
                     @(x             'w'
                     @(y             'o'
                     @[1             'r'
                     @{a             'l'
                     @[2             'd'
                     @[k             '!'

If fully uncompressed, the frame takes more space:

    *rga   #1UQ8p+bart   @1UQ8yk+lisa     :0      !
    *rga   #1UQ8p+bart   @1UQ8s+bart      :0     'H'
    *rga   #1UQ8p+bart   @1UQ8sr+bart     :0     'e'
    *rga   #1UQ8p+bart   @1UQ8t+bart      :0     'l'
    *rga   #1UQ8p+bart   @1UQ8tT+bart     :0     'l'
    *rga   #1UQ8p+bart   @1UQ8ti+bart     :0     'o'
    *rga   #1UQ8p+bart   @1UQ8w+lisa      :0     ' '
    *rga   #1UQ8p+bart   @1UQ8x+lisa      :0     'w'
    *rga   #1UQ8p+bart   @1UQ8y+lisa      :0     'o'
    *rga   #1UQ8p+bart   @1UQ8y1+lisa     :0     'r'
    *rga   #1UQ8p+bart   @1UQ8y1a+lisa    :0     'l'
    *rga   #1UQ8p+bart   @1UQ8y2+lisa     :0     'd'
    *rga   #1UQ8p+bart   @1UQ8yk+lisa     :0     '!'


If rendered in JSON, the same document would probably start as

    {
        "_id": "3b127800-d350-11e7-8000-9a5db8000000",
        "_version": "98f38f80-d351-11e7-8000-c2dde5000000",
        ...

...which is already 90% of the size of the entire compressed frame above.
With idiomatic JSON, per-symbol metadata is both difficult and expensive.

So, let's be precise. Let's put UUIDs on everything. RON makes it possible.

## Wire format (binary)

The binary format is more efficient because of higher bit density; it is also
simpler and safer to parse because of explicit field lengths.  Obviously, it is
not human-readable.

Like the text format, the binary one is only optimized for iteration.  Because of
compression, records are inevitably of variable length, so random access is not
possible.  Also, compression depends on iteration, as UUIDs get abbreviated
relative to similar preceding UUIDs.

A binary RON frame starts with magic bytes `RON2` and frame length.  The rest
of the frame is a sequence of *fields*.  Each field starts with a *descriptor*
specifying the type of the field and its length.

Frame length is serialized as a 32-bit big-endian integer.  The maximum length
of a frame is 2^31-1 bytes.  If the length value has its most significant bit
set to 1, then the frame is *chunked*.  A chunked frame is followed by a
continuation frame.  A continuation frame has no magic bytes, just a 4-byte
length field.  The last continuation frame must have the m.s.b. of its length
set to 0.

### Descriptors

A descriptor's first byte spends four most significant (m.s.) bits to describe
the type of the field, other four bits describe its length.

```
   7    6    5    4    3    2    1    0
+----+----+----+----+----+----+----+----+
| major   | minor   |     field         |
|    type |    type |        length     |
+----+----+----+----+----+----+----+----+
  128  64   32   16    8    4    2    1
   80  40   20   10    8    4    2    1
```

Field descriptor major/minor type bits are set as follows:

0. `00` RON op term,
    * `0000` raw op,
    * `0001` reduced op,
    * `0010` header op,
    * `0011` query header op.
1. `01` UUID, uncompressed
    * `0100` type (reducer) id,
    * `0101` object id,
    * `0110` event id,
    * `0111` ref/location id
2. `10` UUID, compressed (zipped)
    * `1000` value UUID, zipped (note: not type id)
    * `1001` object id,
    * `1010` event id,
    * `1011` ref/location id
3. `11` Atom
    * `1100` value UUID, uncompressed (lengths 1..16)
    * `1101` integer (big-endian, [zigzag-coded][zigzag], lengths 1, 2, 4, 8)
    * `1110` string (UTF-8, length 0..2^31-1)
    * `1111` float (IEEE 754-2008, binary 16, 32 or 64, lengths 2, 4, 8 resp)

A descriptor's four least significant bits encode the length of the field in
question.  The length value given by a descriptor does not include the length
of the descriptor itself.

If a field or a frame is 1 to 16 bytes long then it has its length coded
directly in the four l.s. bits of the descriptor. Zero stands for the length of
16 because most field types are limited to that length.  Op terms specify no
length.  With string atoms, zero denotes the presence of an extended length
field which is either 1 or 4 bytes long. The maximum allowed string length is
2Gb (31 bits).  In case the descriptor byte is exactly `1110 0000`, the m.s.
bit of the next byte denotes the length of the extended length field (`0` for
one, `1` for four bytes).  The rest of the next byte (and possibly other three)
is a big-endian integer denoting the byte length of the string.

Consider a time value query frame: `*now?.`

* 4 bytes are magic bytes (RON, `0101 0010  0100 1111  0100 1110  0011 0010`)
* frame length: 4 bytes (length 5, `0000 0000  0000 0000  0000 0000  0000 0101`)
* op term descriptor: 1 byte (`0011 0000`)
* uncompressed UUID descriptor: 1 byte (cited length 3, `0100 0011`)
* `now` RON UUID: 3 bytes (`0000 1100  1011 0011  1110 1100`,
  the "uncompressed" coding still trims a lot of zeroes, see below).

As UUID length is up to 16 bytes, UUID fields never use a separate length
number. UUID descriptors are always 1 byte long. The length of 0 stands for 16.

Length bits `0000` stand for:

* zero length for op terms,
* 16 for integer/float atoms, zipped/unzipped UUIDs,
* for strings, that signals an extended length record (1 or 4 bytes).

An extended length record is used for strings cause those can be up to 2GB
long. An extended length record is either 1 or four bytes. Four-byte record is
a big-endian 32-bit int having its m.s. bit set to 1. Thus, strings of 127
bytes and shorter may use 1 byte long length record.

### Op terms

Op term fields may have cited length of `0000` or be skipped if they match the
previous op's term.  Still, sometimes we want to introduce redundancy,
CRC/checksumming, hashing, etc.  Exactly for this purpose we may use non-empty
terms.  The checksumming method is specified by the field length (TODO).

### Uncompressed UUIDs

Uncompressed UUIDs are not compressed relative to preceding UUIDs (not *zipped*).
Still, zero bytes are skipped to optimize for some often-used cases.
The skip pattern is determined based on the cited field length.

Namely, UUIDs 1..8 bytes long have the *origin* part set to zeros (all 8 bytes)
and the least significant bytes of the value also set to zeroes.
These are often-used "transcendent" name UUIDs (`lww`, `rga`, `db`, `now`, etc).
For example, `lww` is the data type UUID for last-write-wins objects.
In the unabbreviated RON Base64 form, `lww` is `0/lww0000000 00000000000`
(see the [UUID spec](uuid.md) for the details).

UUIDs 9 to 15 bytes long have their l.s. value bytes set to zero.
This case is optimized for arbitrary-precision timestamps.

UUIDs 16 bytes long are full 128-bit RON UUIDs.

### Compressed UUIDs

Zipped UUIDs are serialized as deltas to similar past UUIDs.  That provides
significant savings when UUIDs come from the same source (same origin bytes) or
have close timestamp values.  Repeated UUIDs are simply skipped, same as in the
Base64 notation.

The origin part is either reused in full or rewritten in full. That is decided
by the field length (<9 reuse, >=9 rewrite). Implicitly, origin ids are
considered uncompressible.

There are two zip modes: *short* and *long*.  In the short mode, an UUID is
compressed relative to the same kind of UUID in the previous op (e.g. event id
relative to the previous event id).  In the long mode, an UUID is compressed
relative to a past uncompressed UUID. A decoder must remember 16 last
uncompressed timestamp-based UUIDs (no names, no hashes), to perform
uncompression. For encoders, that is optional.

A zipped UUID starts with a *zip byte* referencing the compression details.

Short zip byte:

```
   7    6    5    4    3    2    1    0
+----+----+----+----+----+----+----+----+
|  0 | zero tail len|                   |
|    | (half-bytes) |  m.s. half-byte   |
+----+----+----+----+----+----+----+----+
  128  64   32   16    8    4    2    1
```

In this mode, the zip byte specifies how many l.s. half-bytes of the value are
zeroes. Based on the field length, we decide how many "middle" half-bytes need
to be changed, relative to the past UUID. M.s. half-bytes stay the same as in
the past UUID.

Long zip byte:

```
   7    6    5    4    3    2    1    0
+----+----+----+----+----+----+----+----+
|  1 |zero tail len | past uncompressed |
|    |  (half-bytes)|   UUID index      |
+----+----+----+----+----+----+----+----+
  128  64   32   16    8    4    2    1
```

In this mode, the zip byte specifies the past uncompressed UUID we use as a
reference. Index 0 points at the recentmost uncompressed UUID, 1 to the
previous one, etc.
Similarly to the short mode, we set a number of l.s. half-bytes to
zeroes, replace middle half-bytes with new values and keep the m.s. half-bytes
the same.

### Atoms

Strings are serialized as UTF-8.

Integers are serialized using the zig-zag coding (the l.s. bit conveys the sign).

Floats are serialized as IEEE 754 floats (4-byte and 8-byte support is required,
other lengths are optional).

## The math

RON is [log-structured][log]: it stores data as a stream of changes first,
everything else second.  Algorithmically, RON is LSMT-friendly (think [BigTable
and friends][lsmt]).  RON is [information-centric][icn]: the data is addressed
independently of its place of storage (think [git][git]).  RON is CRDT-friendly;
[Conflict-free Replicated Data Types][crdt] enable real-time data sync (think
Google Docs).

Swarm RON employs a variety of well-studied computer science models.  The
general flow of RON data synchronization follows the state machine replication
model.  Offline writability, real-time sync and conflict resolution are all
possible thanks to [Commutative Replicated Data Types][crdt] and [partially
ordered][po] op logs.  UUIDs are essentially [Lamport logical
timestamps][lamport], although they borrow a lot from RFC4122 UUIDs.  RON wire
format is a [regular language][regular].  That makes it (formally) simpler than
either JSON or XML.

The core contribution of the RON format is *practicality*.  RON arranges
primitives in a way to make metadata overhead acceptable.  Metadata was a known
hurdle in CRDT-based solutions, as compared to e.g. [OT-family][ot] algorithms.
Small overhead enables such real-time apps as collaborative text editors where
one op is one keystroke.  Hopefully, it will enable some yet-unknown
applications as well.

Use Swarm RON!

## Acknowledgements

* Russell Sullivan
* Yuriy Syrovetskiy

## History

* 2012-2013: project started (initially, as a part of the Yandex Live Letters
  project)
* 2014 Feb: becomes a separate project
* 2014 Oct: version 0.3 is demoed (per-object logs and version vectors, not
  really scalable)
* 2015 Sep: version 0.4 is scrapped, the math is changed to avoid any version
  vector use
* 2016 Feb: version 1.0 stabilizes (no v.vectors, new asymmetric client
  protocol)
* 2016 May: version 1.1 gets peer-to-peer (server-to-server) sync
* 2016 Jun: version 1.2 gets crypto (Merkle, entanglement)
* 2016 Oct: functional generalizations (map/reduce)
* 2016 Dec: cross-columnar compression
* 2017 Jun: Swarm RON 2.0.0
* 2017 Jul: new frame-based Causal Tree / Replicated Growable Array
  implementation
* 2017 Jul: Ragel parser
* 2017 Aug: punctuation tweaks
* 2017 Oct: streaming parser
* 2017 Oct: binary encoding

## Build status

| Package             | Build status         |
|---------------------|----------------------|
| `gritzko/ron`       | [![RON][1]][travis]  | 
| `gritzko/ron/rdt`   | [![CRDTs][1]][travis]| 


[2sided]: http://lexicon.ft.com/Term?term=two_sided-markets
[super]: http://ilpubs.stanford.edu:8090/594/1/2003-33.pdf
[opbased]: http://haslab.uminho.pt/sites/default/files/ashoker/files/opbaseddais14.pdf
[cap]: https://www.infoq.com/articles/cap-twelve-years-later-how-the-rules-have-changed
[swarm]: https://gritzko.gitbooks.io/swarm-the-protocol/content/
[po]: https://en.wikipedia.org/wiki/Partially_ordered_set#Formal_definition
[crdt]: https://en.wikipedia.org/wiki/Conflict-free_replicated_data_type
[icn]: http://www.networkworld.com/article/3060243/internet/demystifying-the-information-centric-network.html
[kafka]: http://kafka.apache.org
[git]: https://git-scm.com
[log]: http://blog.notdot.net/2009/12/Damn-Cool-Algorithms-Log-structured-storage
[re]: https://blogs.msdn.microsoft.com/csliu/2009/11/10/mapreduce-in-functional-programming-parallel-processing-perspectives/
[rfc4122]: https://tools.ietf.org/html/rfc4122
[causal]: https://en.wikipedia.org/wiki/Causal_consistency
[UUID]: https://en.wikipedia.org/wiki/Universally_unique_identifier
[peterb]: https://martin.kleppmann.com/2014/11/isolation-levels.png
[regular]: https://en.wikipedia.org/wiki/Regular_language
[mvc]: https://en.wikipedia.org/wiki/Model–view–controller
[ot]: https://en.wikipedia.org/wiki/Operational_transformation
[lamport]: http://lamport.azurewebsites.net/pubs/time-clocks.pdf
[2problems]: https://martinfowler.com/bliki/TwoHardThings.html
[lsmt]: https://en.wikipedia.org/wiki/Log-structured_merge-tree
[zigzag]: https://developers.google.com/protocol-buffers/docs/encoding#signed-integers

[travis]: https://travis-ci.com/gritzko/ron
[1]: https://travis-matrix-badges.herokuapp.com/repos/gritzko/ron/branches/master/1
[2]: https://travis-matrix-badges.herokuapp.com/repos/gritzko/ron/branches/master/2