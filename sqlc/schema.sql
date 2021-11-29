CREATE TABLE IF NOT EXISTS employee (
   id bigserial PRIMARY KEY,            -- 編號
   name varchar(60) UNIQUE NOT NULL,    -- 名稱
   age integer,                         -- 年齡
   created_on timestamp NOT NULL        -- 建立時間
);