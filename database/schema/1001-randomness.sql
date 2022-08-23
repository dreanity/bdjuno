CREATE TABLE randomness_chain_info
(
	"hash"         TEXT NOT NULL PRIMARY KEY,
	"public_key"   TEXT NOT NULL,
	"period"       BIGINT NOT NULL,
	"genesis_time" BIGINT NOT NULL
)

CREATE TABLE randomness
(
	"round" 			 BIGINT NOT NULL PRIMARY KEY,
	"round_time" 		 BIGINT NOT NULL,
	"randomness" 		 TEXT,
	"signature" 		 TEXT,
	"previous_signature" TEXT
)

