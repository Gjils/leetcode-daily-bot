db.createUser({
	user: "root",
	pwd: "root",
	roles: [
		{
			role: "readWrite",
			db: "botDB",
		},
	],
});

db = new Mongo().getDB("botDB");

db.createCollection("users", { capped: false });
db.createCollection("groups", { capped: false });
db.createCollection("admins", { capped: false });
