package ron

// An RDT may have a combination of the following features:
//
// * Associative (ab)c = a(bc)
// * Commutative ab = ba
// * Idempotent aa = a
// * Distributed ab=ba, but for concurrent ops only
//
// These features tell us precisely what we can do with an object of a given type.
// e.g. ACID_NONE are single-writer objects;
// ACID_D is the bare minimum necessary for a multiple-writer type.
// ACID_FULL are the most robust: they survive everything but data loss;
// ACID_AID are close to ACID_FULL, except they depend on causal delivery order.
const (
	ACID_NONE = iota // single-writer, strictly "state + linear op log", like MySQL
	ACID_D    = 1    // multiple-writer, partial-order
	ACID_I    = 2    // survives data duplication
	ACID_ID   = ACID_I | ACID_D
	ACID_C    = 4 // arbitrary order (causality violations don't break convergence)
	ACID_CD   = ACID_C | ACID_D
	ACID_CID  = ACID_CD | ACID_I
	ACID_A    = 8 // can form patches
	ACID_AD   = ACID_A | ACID_D
	ACID_AI   = ACID_I | ACID_A
	ACID_AID  = ACID_A | ACID_ID
	ACID_ACD  = ACID_C | ACID_AD
	ACID_FULL = ACID_A | ACID_CID
)
