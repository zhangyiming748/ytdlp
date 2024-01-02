select * from ytdlps order by created_at desc;
select name,status from ytdlps;

select url from ytdlps where status = '下载失败' order by id desc;

