# Replicated Object Notation 2.1

The 2.1 version will introduce crypto-related features to make RON suitable for distributed revision control and decentralized web. 

Dynamic structured data (i.e. a database) in a decentralized environment is known to be a difficult problem. 

RON 2.1 will maintain backward compatibility with RON 2.0. The new version will introduce four key changes. 

## Merkle structure

RON ops will form a Merkle DAG. The hashes will not depend on the serialization format (text, binary, etc) or reductions. Ops are logically immutable, maintaining their hash for their entire lifecycle. 

In general, Merkle hashing requires stricter, more disciplined approach, hence the introduction of the nominal format and RON-open.
RON 2.1 introduces content- addressing (hashes) along with UUIDs. That also necessitates offset-uuid-hash conversions, see below.

## Nominal format

RON may use various serialization formats,  depending on convenience and performance needs. But, all those formats allow for some bitwise liberties (e.g. field order in JSON, arbitrary whitespace in any textual format, varint lengths in binary protocols, etc). Historically,  that was a headache for every single format (e.g. how to sign an XML), and we use several formats. 

To enable uniform bitwise hashing,  we "legalise" the existing iterator-internal binary format as the *nominal* RON format (all atoms 16 bytes wide, little endian, using the 60+4 bit layout).
All hashes are defined relative to the nominal format.

Also, the nominal RON serves as a universal intermediary format.
That will ease experiments with new serialization formats (e.g. JSON, CBOR) as we only have to implement a parser (x to nominal), and a builder (nominal to x). 

## Transparent offset-uuid-hash conversions

Any RON entity is referenced by its UUID. RON 2.1 will also be able to reference by an offset (in a trusted local context) and by a hash (untrusted distributed context). 
RON can not use just one of the three, for a  number of reasons (e.g. a typical op is smaller than its hash).
All three kinds of references (offsets, UUIDs and hashes) must be transparently convertible back and forth, as they refer to the same entities.

## Open RON, closed RON

As of 2.0, a RON op certainly has some redundancy. An op specifier consists of four UUIDs: type, object,  event and location ids, which may be derived from each other in many cases. 

While a detailed four-component specifier makes it easy to process an op, it also makes it more difficult to handle a RON database as a Merkle structure. 
In general, that redundancy may lead to contradictions and vulnerabilities. 

The solution is to structure each object as a causal tree to make object and type ids 100% derivable. In theory, that creates two versions of the protocol: *Open RON* (two-UUID specifier, event and ref) and *Closed RON* (four-UUID specifier, for "thin" clients only). In practice, full replicas may ignore type/object ids, the protocol stays the same.

All these changes lead to corresponding changes in RDTs. Most notably,  `lww` gets deprecated. The notation stays the same. 

-------

RON 2.1 development is supported by Protocol Labs (and JetBrains).
