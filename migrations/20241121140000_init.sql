-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.song
(
    song_id uuid NOT NULL DEFAULT gen_random_uuid(),
    "group" character varying COLLATE pg_catalog."default" NOT NULL,
    name character varying COLLATE pg_catalog."default" NOT NULL,
    release_date date,
    text character varying[] COLLATE pg_catalog."default" NOT NULL DEFAULT '{}'::character varying[],
    link character varying COLLATE pg_catalog."default",
    CONSTRAINT song_pkey PRIMARY KEY (song_id)
);
INSERT INTO public.song VALUES
('87b42eb1-5ba8-4662-96dc-0c4aa52e453c','Кино','Звезда по имени Солнце','1989-01-01',E'{"Белый снег, серый лед,\nНа растрескавшейся земле.\nОдеялом лоскутным на ней -\nГород в дорожной петле.","А над городом плывут облака,\nЗакрывая небесный свет.\nА над городом - желтый дым,\nГороду две тысячи лет,\nПрожитых под светом Звезды\nПо имени Солнце...","И две тысячи лет - война,\nВойна без особых причин.\nВойна - дело молодых,\nЛекарство против морщин.","Красная, красная кровь\nЧерез час уже просто земля,\nЧерез два на ней цветы и трава,\nЧерез три она снова жива\nИ согрета лучами Звезды\nПо имени Солнце...","И мы знаем, что так было всегда,\nЧто Судьбою больше любим,\nКто живет по законам другим\nИ кому умирать молодым.","Он не помнит слово да и слово нет,\nОн не помнит ни чинов, ни имен.\nИ способен дотянуться до звезд,\nНе считая, что это сон,\nИ упасть, опаленным Зведой по имени Солнце"}','https://youtu.be/jQV5VXfKDYc'),
('f3e4d076-9ec0-4b90-a64f-8d08f6992e32','Кино','Кукушка','1990-01-01',E'{"Песен еще ненаписанных, сколько?\nСкажи, кукушка, пропой.\nВ городе мне жить или на выселках,\nКамнем лежать или гореть звездой?\nЗвездой.","Солнце мое – взгляни на меня,\nМоя ладонь превратилась в кулак,\nИ если есть порох – дай огня.\nВот так...","Кто пойдет по следу одинокому?\nСильные да смелые\nГоловы сложили в поле в бою.\nМало кто остался в светлой памяти,\nВ трезвом уме да с твердой рукой в строю,\nВ строю.","Солнце мое – взгляни на меня,\nМоя ладонь превратилась в кулак,\nИ если есть порох – дай огня.\nВот так...","Где же ты теперь, воля вольная?\nС кем же ты сейчас\nЛасковый рассвет встречаешь? Ответь.\nХорошо с тобой, да плохо без тебя,\nГолову да плечи терпеливые под плеть,\nПод плеть.","Солнце мое – взгляни на меня,\nМоя ладонь превратилась в кулак,\nИ если есть порох – дай огня.\nВот так..."}','https://youtu.be/yNp9SBW4xTA'),
('3a7c2d18-29fb-4a7e-82d8-2800c510ec10','Nautilus Pompilius','Прогулки по воде','1993-01-01',E'{"С пpичала pыбачил апостол Андpей\nА Спаситель ходил по воде\nИ Андpей доставал из воды пескаpей\nА Спаситель погибших людей\nИ Андpей закpичал: я покину пpичал\nЕсли ты мне откpоешь секpет\nИ Спаситель ответил: Спокойно Андpей\nНикакого секpета здесь нет\nВидишь там на гоpе возвышается кpест\nПод ним десяток солдат повиси-ка на нем\nА когда надоест возвpащайся назад\nГулять по воде\nГулять по воде\nГулять по воде со мной\n","Hо учитель на касках блистают pога\nЧеpный воpон кpужит над кpестом\nОбъясни мне сейчас пожалей дуpака\nА pаспятье оставь на потом\nОнемел Спаситель и топнул в сеpдцах\nПо водной глади ногой\nТы и веpно дуpак и Андpей в слезах\nПобpел с пескаpями домой\nВидишь там на гоpе возвышается кpест\nПод ним десяток солдат повиси-ка на нем\nА когда надоест возвpащайся назад\nГулять по воде\nГулять по воде\nГулять по воде со мной\nВидишь там на гоpе возвышается кpест\nПод ним десяток солдат повиси-ка на нем\nА когда надоест возвpащайся назад\nГулять по воде\nГулять по воде\nГулять по воде со мной"}','https://youtu.be/vQWUcWn-OJc'),
('96375e8f-c402-4548-84c8-a3c2e1991a16','Кино','Группа крови','1988-01-01',E'{"Тёплое место\nНо улицы ждут oтпечатков наших ног\nЗвёздная пыль — на сапогах.\nМягкое кресло, клетчатый плед,\nНе нажатый вовремя курок.\nСолнечный день - в ослепительных снах.","Группа крови — на рукаве,\nМой порядковый номер — на рукаве,\nПожелай мне удачи в бою, пожелай мне:\nНе остаться в этой траве,\nНе остаться в этой траве.\nПожелай мне удачи, пожелай мне удачи!","И есть чем платить, но я не хочу\nПобеды любой ценой.\nЯ никому не хочу ставить ногу на грудь.\nЯ хотел бы остаться с тобой,\nПросто остаться с тобой,\nНо высокая в небе звезда зовёт меня в путь.","Группа крови — на рукаве,\nМой порядковый номер — на рукаве,\nПожелай мне удачи в бою, пожелай мне:\nНе остаться в этой траве,\nНе остаться в этой траве.\nПожелай мне удачи, пожелай мне удачи"}','https://youtu.be/k9p4B6s0j4U'),
('3459ebd5-936d-402c-9a78-19f706cbbb17','Кино','Перемен','1989-01-01',E'{"Вместо тепла — зелень стекла,\nВместо огня — дым,\nИз сетки календаря выхвачен день.\nКрасное солнце сгорает дотла,\nДень догорает с ним,\nНа пылающий город падает тень.","Перемен! — требуют наши сердца.\nПеремен! — требуют наши глаза.\nВ нашем смехе и в наших слезах,\nИ в пульсации вен —\nПеремен!\nМы ждём перемен!","Электрический свет продолжает наш день,\nИ коробка от спичек пуста,\nНо на кухне синим цветком горит газ.\nСигареты в руках, чай на столе —\nЭта схема проста,\nИ больше нет ничего, всё находится в нас.","Перемен! — требуют наши сердца.\nПеремен! — требуют наши глаза.\nВ нашем смехе и в наших слезах,\nИ в пульсации вен —\nПеремен!\nМы ждём перемен!","Мы не можем похвастаться мудростью глаз\nИ умелыми жестами рук,\nНам не нужно всё это, чтобы друг друга понять.\nСигареты в руках, чай на столе —\nТак замыкается круг,\nИ вдруг нам становится страшно что-то менять.","Перемен! — требуют наши сердца.\nПеремен! — требуют наши глаза.\nВ нашем смехе и в наших слезах,\nИ в пульсации вен —\nПеремен!\nМы ждём перемен!"}','https://youtu.be/jyorQevSPI0'),
('c748cc8f-4843-4c09-aa19-dd652675043d','Кино','В наших глазах','1988-01-01',E'{"Постой, не уходи!\nМы ждали лета - пришла зима.\nМы заходили в дома,\nНо в домах шел снег.\nМы ждали завтрашний день,\nКаждый день ждали завтрашний день.\nМы прячем глаза за шторами век.","В наших глазах крики Вперед!\nВ наших глазах окрики Стой!\nВ наших глазах рождение дня\nИ смерть огня.\nВ наших глазах звездная ночь,\nВ наших глазах потерянный рай,\nВ наших глазах закрытая дверь.\nЧто тебе нужно? Выбирай!","Мы хотели пить, не было воды.\nМы хотели света, не было звезды.\nМы выходили под дождь\nИ пили воду из луж.\nМы хотели песен, не было слов.\nМы хотели спать, не было снов.\nМы носили траур, оркестр играл туш...","В наших глазах крики Вперед!\nВ наших глазах окрики Стой!\nВ наших глазах рождение дня\nИ смерть огня.\nВ наших глазах звездная ночь,\nВ наших глазах потерянный рай,\nВ наших глазах закрытая дверь.\nЧто тебе нужно? Выбирай!"}','https://youtu.be/jyorQevSPI0'),
('63318518-828f-49d4-b93b-a7e9212222af','Кино','Электричка','1982-01-01',E'{"Я вчера слишком поздно лёг, сегодня рано встал,\nЯ вчера слишком поздно лёг, я почти не спал.\nМне, наверно, с утра нужно было пойти к врачу,\nА теперь электричка везёт меня туда, куда я не хочу.","Электричка везёт меня туда, куда я не хочу","В тамбуре холодно, и в то же время как-то тепло,\nВ тамбуре накурено, и в то же время как-то свежо.\nПочему я молчу, почему не кричу? Молчу.","Электричка везёт меня туда, куда я не хочу"}','https://youtu.be/4lhxm6_aqAA'),
('a5bb1db4-bf4d-432b-9512-ea897c4c074a','Кино','Мама, Мы Все Тяжело Больны','1988-01-01',E'{"Зерна упали в землю, зерна просят дождя.\nИм нужен дождь.\nРазрежь мою грудь, посмотри мне внутрь,\nТы увидишь, там все горит огнем.\nЧерез день будет поздно, через час будет поздно,\nЧерез миг будет уже не встать.\nЕсли к дверям не подходят ключи, вышиби двери плечом.","Мама, мы все тяжело больны...\nМама, я знаю, мы все сошли с ума...","Сталь между пальцев, сжатый кулак.\nУдар выше кисти, терзающий плоть,\nНо вместо крови в жилах застыл яд, медленный яд.\nРазрушенный мир, разбитые лбы, разломанный надвое хлеб.\nИ вот кто-то плачет, а кто-то молчит,\nА кто-то так рад, кто-то так рад..","Мама, мы все тяжело больны...\nМама, я знаю, мы все сошли с ума...","Ты должен быть сильным, ты должен уметь сказать:\nРуки прочь, прочь от меня!\nТы должен быть сильным, иначе зачем тебе быть.\nЧто будет стоить тысячи слов,\nКогда важна будет крепость руки?\nИ вот ты стоишь на берегу и думаешь: Плыть или не плыть?","Мама, мы все тяжело больны...\nМама, я знаю, мы все сошли с ума..."}','https://youtu.be/pLE2ngmVqZY'),
('6b150eac-3729-49ac-9aeb-48d9c9e983de','Кино','Пачка сигарет','1989-01-01',E'{"Я сижу и смотрю в чужое небо из чужого окна\nИ не вижу ни одной знакомой звезды.\nЯ ходил по всем дорогам и туда, и сюда,\nОбернулся — и не смог разглядеть следы.","Но если есть в кармане пачка сигарет,\nЗначит все не так уж плохо на сегодняшний день.\nИ билет на самолет с серебристым крылом,\nЧто, взлетая, оставляет земле лишь тень.","И никто не хотел быть виноватым без вина,\nИ никто не хотел руками жар загребать,\nА без музыки и на миру смерть не красна,\nА без музыки не хочется пропадать.","Но если есть в кармане пачка сигарет,\nЗначит все не так уж плохо на сегодняшний день.\nИ билет на самолет с серебристым крылом,\nЧто, взлетая, оставляет земле лишь тень."}','https://youtu.be/ciyB5a189_o'),
('e6670633-7396-424c-98db-fe2a0b63ac5a','Пикник','У шамана три руки','2005-01-01',E'{"У шамана три руки\nИ крыло из-за плеча\nОт дыхания его\nРазгорается свеча\nИ порою сам себя\nСам себя не узнает\nА распахнута душа\nНадрывается, поет","У шамана три руки\nМир вокруг, как темный зал\nНа ладонях золотых\nНарисованы глаза\nВидит розовый рассвет\nПрежде солнца самого\nА казалось, будто спит\nИ не знает ничего","У шамана три руки.\nСад в рубиновых лучах.\nОт дыхания его\nРазгорается... Разгорается...\nРазгорается...","Разгорается свеча...\nРазгорается свеча.","У шамана три...\nУ шамана три...\nУ шамана три..."}','https://youtu.be/Xhsl6j6xOEM'),
('1a8b1880-e3aa-4c60-823e-25e632642276','Пикник','Там, на самом, на краю Земли','1991-01-01',E'{"Там, на самом на краю Земли\nВ небывалой голубой дали\nВнемля звукам небывалых слов,\nСладко-сладко замирает кровь.","Там ветра летят, касаясь звeзд\nТам деревья не боятся гроз\nОкеаном бредят корабли\nТам, на самом, на краю Земли.","Что ж ты, сердце, рвeшься из груди?\nПогоди немного, погоди\nЧистый голос в небесах поeт\nСветлый полдень над Землeй встаeт."}','https://youtu.be/YIdEbLa6L6s'),
('f8a11d33-cfc7-4dab-8e93-d3d35397b086','Пикник','Из мышеловки','2007-01-01',E'{"Ещё грязь не смыта с кожи,\nТолько страха больше нет.\nПотому затих прохожий,\nУдивленно смотрит вслед.","А движения неловки,\nБудто бы из мышеловки,\nБудто бы из мышеловки,\nТолько вырвалась она.","А движения неловки,\nБудто бы из мышеловки,\nБудто бы из мышеловки,\nТолько вырвалась она.","А в груди попеременно,\nБыл то пепел, то алмаз.\nТо лукаво то надменно,\nНочь прищуривала глаз.","А движения неловки,\nБудто бы из мышеловки,\nБудто бы из мышеловки,\nТолько вырвалась она.","За три слова до паденья,\nЗа второе за рожденье,\nЗа второе за рожденье,\nПей до полночи одна.","И уже не искалечит,\nСмех лощенных дураков.\nВедь запрыгнул ей на плечи,\nМокрый ангел с облаков.","А движения...","А движения неловки,\nБудто бы из мышеловки,\nБудто бы из мышеловки,\nТолько вырвалась она."}','https://youtu.be/bS5ZXt-PVlY'),
('6b1b627f-1609-488c-9c38-c87a3b70d33d','Nautilus Pompilius','Крылья','1995-01-01',E'{"Ты снимаешь вечернее платье,\nСтоя лицом к стене.\nИ я вижу свежие шрамы\nНа гладкой, как бархат, спине.","Мне хочется плакать от боли\nИли забыться во сне,\nГде твои крылья,\nКоторые так нравились мне?","Где твои крылья,\nКоторые нравились мне?\nГде твои крылья,\nКоторые нравились мне?","Раньше у нас было время,\nТеперь у нас есть дела —\nДоказывать, что сильный жрет слабых,\nДоказывать, что сажа бела.","Мы все потеряли что-то\nНа этой безумной войне.\nКстати, где твои крылья,\nКоторые нравились мне?","Где твои крылья,\nКоторые нравились мне?\nГде твои крылья,\nКоторые нравились мне?","Я не спрашиваю, сколько у тебя денег,\nНе спрашиваю сколько мужей.\nЯ вижу — ты боишься открытых окон\nИ верхних этажей.","И если завтра начнется пожар\nИ все здание будет в огне,\nМы погибнем без этих крыльев,\nКоторые нравились мне."}','https://youtu.be/ouKj-7bc-ZQ'),
('65347c0d-2a7f-4bf8-b892-ae7f012c417a','Nautilus Pompilius','Музыка','1994-01-01',E'{"Звуки пойманные слухом\nВы сроднитесь с плотным духом\nВы вплетаете свободно\nВ пульс творенья что угодно\nСердца нашего биенья\nВы живые отраженья","Звук обступает и слева и справа\nСмутным таинственным войском вдали\nПеной ложится у кромки земли\nИ до утра до скончания ночи\nЗвук обступает и слева и справа\nБесится бегает дико хохочетс\nСмутным таинственным войском вдали","Стрелами света кружит над тобой\nСхватка оркестра с волшебной трубой\nчтобы и завтра безвестная сила\nВсе повторяла и все возносила Все повторяла и все возносила\nСловно и радость твоя и тоска\nСтрелами света кружит над тобой","Море бурное оркестра\nВот тогда душа моя\nСнова музыка поманит\nЗвуки пойманные слухом\nВы сроднитесь с плотным духом\nВы летаете свободно","За крылатым вашим дымом\nЯ сыщу владенья ваши"}','https://youtu.be/0aTzMbrAW20');
--add_song
CREATE OR REPLACE FUNCTION public.add_song(
	"group" character varying,
	name character varying,
	release_date date,
	text character varying[],
	link character varying)
    RETURNS SETOF song 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
BEGIN
	RETURN QUERY
	INSERT INTO public.song
	VALUES (
	DEFAULT,
	add_song.group,
	add_song.name,
	add_song.release_date,
	COALESCE(add_song.text,'{}'::character varying[]),
	add_song.link
	)
	RETURNING *;
END
$BODY$;
--get_info
CREATE OR REPLACE FUNCTION public.get_info(
	"group" character varying,
	name character varying)
    RETURNS SETOF song 
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
BEGIN
	RETURN QUERY
	SELECT *
	FROM public.song
	WHERE 
		song.group=get_info.group AND
		song.name=get_info.name
	LIMIT 1;
END;
$BODY$;
--get_song
CREATE OR REPLACE FUNCTION public.get_song(
	song_id uuid)
    RETURNS SETOF song 
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
BEGIN
	RETURN QUERY
	SELECT * 
	FROM public.song
	WHERE song.song_id=get_song.song_id;
END
$BODY$;
--get_song_text
CREATE OR REPLACE FUNCTION public.get_song_text(
	song_id uuid,
	"offset" integer,
	"limit" integer)
    RETURNS character varying[]
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
AS $BODY$
BEGIN
	RETURN (
		SELECT ( 
			SELECT "text" FROM public.song WHERE song.song_id=get_song_text.song_id
		)[get_song_text.offset+1:get_song_text.limit+get_song_text.offset] AS verse
	); 
END
$BODY$;
--get_songs
CREATE OR REPLACE FUNCTION public.get_songs(
	"offset" integer,
	"limit" integer,
	"group" character varying,
	name character varying,
	release_date date,
	text character varying,
	link character varying)
    RETURNS SETOF song 
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
BEGIN
	RETURN QUERY
	SELECT *
	FROM public.song
	WHERE 
		(get_songs.group IS null OR song.group ILIKE '%'||get_songs.group||'%') AND
		(get_songs.name IS null OR song.name ILIKE '%'||get_songs.name||'%') AND
		(get_songs.release_date IS null OR song.release_date = get_songs.release_date::date) AND
		(get_songs.text IS null OR array_to_string(song.text,'\n\n') ILIKE '%'||get_songs.text||'%') AND
		(get_songs.link IS null OR song.link ILIKE '%'||get_songs.link||'%')
	OFFSET get_songs.offset LIMIT get_songs.limit; 
END
$BODY$;
--remove_song
CREATE OR REPLACE FUNCTION public.remove_song(
	song_id uuid)
    RETURNS SETOF uuid 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
BEGIN
	RETURN QUERY
	DELETE FROM public.song
	WHERE song.song_id=remove_song.song_id
	RETURNING song.song_id ;
END
$BODY$;
--update_song
CREATE OR REPLACE FUNCTION public.update_song(
	song_id uuid,
	"group" character varying,
	name character varying,
	release_date date,
	text character varying[],
	link character varying)
    RETURNS SETOF uuid 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
BEGIN
	RETURN QUERY
	UPDATE public.song
		SET
			"group"=update_song.group,
			"name"=update_song.name,
			release_date=update_song.release_date,
			"text"= update_song.text,
			"link"=update_song.link
	WHERE song.song_id=update_song.song_id
	RETURNING song.song_id ;
END
$BODY$;
-- +goose StatementEnd
-- +goose Down
DROP DATABASE song_library;