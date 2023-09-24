CREATE TABLE rooms(
    id VARCHAR(255) PRIMARY KEY,
	roundDuration INT,
	roomDuration INT,
	roomStartDate INT,
	roundStartDate INT
);
CREATE TABLE users_rooms(
    id_room VARCHAR(255) REFERENCES rooms(id) ON DELETE CASCADE NOT NULL,
    id_user INT,
    id_wallet VARCHAR(255), 
    score INT
);