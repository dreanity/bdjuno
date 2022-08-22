CREATE TABLE randomness_chain_info
(
	"hash"         TEXT NOT NULL PRIMARY KEY,
	"public_key"   TEXT NOT NULL,
	"period"       BIGINT NOT NULL,
	"genesis_time" BIGINT NOT NULL
)

CREATE TABLE unproven_randomness
(
	"round" BIGINT NOT NULL PRIMARY KEY
)

CREATE TABLE proven_randomness
(
	"round" 			 BIGINT NOT NULL PRIMARY KEY,
	"randomness" 		 TEXT NOT NULL,
	"signature" 		 TEXT NOT NULL,
	"previous_signature" TEXT NOT NULL
)

