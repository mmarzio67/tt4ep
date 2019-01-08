drop table trackt cascade;
drop table users;
drop table project;
drop table task;
drop table actions;



create table users (
  id        serial primary key,
  name      varchar(32),
	surname   varchar(32),
	func  varchar(32)
);


create table project (
  id        serial primary key,
  name      varchar(32),
	description  varchar(64)
);


create table task (
  id        serial primary key,
  name      varchar(32),
	description varchar(64)
);


create table actions (
  id        serial primary key,
  name      varchar(16)
);

create table trackt (
  id         serial primary key,
  usid       integer references users(id),
  action_id  integer references actions(id),
	project_id integer references project(id),
	task_id    integer references task(id),
	createdAt  timestamp
);