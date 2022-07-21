ALTER TABLE cluster_role
	ADD COLUMN replicas INT(11) NOT NULL DEFAULT 0;

ALTER TABLE cluster_role
	ADD COLUMN ready_replicas INT(11) NOT NULL DEFAULT 0;

ALTER TABLE cluster_role
	ADD COLUMN api_version VARCHAR(50) NOT NULL DEFAULT "";

ALTER TABLE cluster_node
	ADD COLUMN host_id VARCHAR(50) NOT NULL DEFAULT "";

ALTER TABLE cluster_node
	ADD COLUMN host_ip VARCHAR(50) NOT NULL DEFAULT "";
