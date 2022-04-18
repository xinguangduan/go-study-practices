-- auto-generated definition
create table english_words
(
    id         int          not null,
    word_name  varchar(50)  not null
        primary key,
    sound_mark varchar(50)  null,
    paraphrase varchar(200) null,
    frequency  int          null,
    memo       varchar(200) null
);