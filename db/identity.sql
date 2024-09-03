use oracle_x;

create table if not exists user(
    id bigint primary key not null,
    username varchar(128) not null,
    email varchar(128) not null, 
    avatar varchar(128) default '' not null,
    password_hash varchar(128) not null,
    password_salt varchar(128) not null,
    is_deleted tinyint default 0 comment '0:no|1:yes',
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
    deleted_at timestamp,
    index(username, is_deleted),
    index(email, is_deleted)
)engine='innodb' charset='utf8mb4';

create table if not exists telegram(
    id bigint primary key not null comment 'telegram uid|TODO: insertion perfomance',
    uid bigint unique not null,
    username varchar(128) default '' not null,
    first_name varchar(128) default '' not null,
    last_name varchar(128) default '' not null,
    creatd_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
)engine='innodb' charset='utf8mb4';


create table if not exists wallet(
    id int primary key auto_increment not null,
    uid bigint unique not null,
    wallet_type tinyint default 0 not null comment '1:ethereum-compatible wallet|2:solana|3:ton',  
    wallet_address varchar(64) not null,
    chain_id int default 0 not null comment '',
    chain_name varchar(32) default '' not null,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
    is_deleted tinyint default 0 not null,
    unique(wallet_address, wallet_type, is_deleted)
)engine='innodb' charset='utf8mb4';


create table if not exists wallet_login_nonce(
    id int primary key auto_increment not null,
    wallet_type tinyint default 0 not null comment '1:evm|2:solana|3:ton',
    wallet_address varchar(64) default '' not null comment 'ceil(max(eth:42,sol:44,ton:48))=64',
    nonce varchar(512) default '' not null,
    created_at timestamp default current_timestamp not null
)engine='innodb' charset='utf8mb4';
